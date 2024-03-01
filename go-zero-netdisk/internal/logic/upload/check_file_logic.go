package upload

import (
	"context"
	"errors"
	"github.com/yitter/idgenerator-go/idgen"
	"github.com/zeromicro/go-zero/core/logx"
	"lc/netdisk/common/constant"
	"lc/netdisk/common/redis"
	"lc/netdisk/common/variable"
	"lc/netdisk/internal/svc"
	"lc/netdisk/internal/types"
	"lc/netdisk/model"
	"math"
	"strconv"
	"time"
)

type CheckFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckFileLogic {
	return &CheckFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckFileLogic) CheckFile(req *types.CheckFileReq) (*types.CheckFileResp, error) {
	var fileFs model.FileFs
	has, err := l.svcCtx.Xorm.Where("hash = ?", req.Hash).Get(&fileFs)
	if err != nil {
		return nil, err
	} else if !has { // 文件不存在时
		return l.doWhenNotExist(req)
	}

	// 文件存在
	return l.doWhenExist(req, &fileFs)
}

func (l *CheckFileLogic) doWhenNotExist(req *types.CheckFileReq) (*types.CheckFileResp, error) {
	var (
		userId = l.ctx.Value(constant.UserIdKey).(int64)
		rdb    = l.svcCtx.Redis
		resp   types.CheckFileResp
	)

	fileId := idgen.NextId()
	fileIdStr := strconv.FormatInt(fileId, 10)
	key := redis.UploadCheckKey + fileIdStr
	resp.Status = constant.StatusFileUnuploaded
	resp.FileId = fileId
	fileInfo := map[string]interface{}{
		"fileId":   fileId,
		"folderId": req.FolderId,
		"hash":     req.Hash,
		"ext":      req.Ext,
		"name":     req.Name,
		"size":     req.Size,
		"userId":   userId,
	}

	// 大文件情况
	if req.Size > int64(constant.ShardingFloor) {
		resp.ConfirmShard = constant.ConfirmShard
		key = redis.UploadCheckBigFileKey + fileIdStr
		fileInfo["chunkNum"] = math.Ceil(float64(req.Size) / constant.ShardingSize)
		fileInfo["chunkSum"] = 0
		if _, err := rdb.HSet(l.ctx, key, fileInfo).Result(); err != nil {
			return nil, err
		}
	}
	if _, err := rdb.HSet(l.ctx, key, fileInfo).Result(); err != nil {
		return nil, err
	}
	go rdb.Expire(l.ctx, key, redis.UploadCheckExpire)
	return &resp, nil
}

func (l *CheckFileLogic) doWhenExist(req *types.CheckFileReq, fileFs *model.FileFs) (*types.CheckFileResp, error) {
	var (
		userId = l.ctx.Value(constant.UserIdKey).(int64)
		engine = l.svcCtx.Xorm
		file   model.File
		resp   types.CheckFileResp
	)

	// 先判断该用户在该目录下有无该文件
	has, err := engine.Where("fs_id = ?", fileFs.Id).
		And("folder_id = ?", req.FolderId).
		And("user_id = ?", userId).Get(&file)
	if err != nil {
		return nil, err
	} else if has {
		return nil, errors.New("当前文件夹已存在该文件😈")
	}

	// 该文件夹无该文件，信息落库
	isBigFlag := constant.SmallFileFlag
	if fileFs.Size > int64(constant.ShardingFloor) {
		isBigFlag = constant.BigFileFlag
	}
	file.UserId = userId
	file.FsId = fileFs.Id
	file.Name = req.Name
	file.FolderId = req.FolderId
	file.Type = variable.GetTypeByBruteForce(req.Ext)
	file.Status = constant.StatusFileUploaded
	file.Url = fileFs.Url
	file.ObjectName = fileFs.ObjectName
	file.IsBig = isBigFlag
	file.DoneAt = time.Now().Local()
	file.DelFlag = constant.StatusFileUndeleted
	if fileFs.Size > int64(constant.ShardingFloor) {
		file.IsBig = constant.BigFileFlag
	}
	if _, err = engine.Insert(&file); err != nil {
		return nil, err
	}

	resp.Status = constant.StatusFileUploaded
	return &resp, nil
}
