package upload

import (
	"context"
	"github.com/yitter/idgenerator-go/idgen"
	"lc/netdisk/common/constant"
	"lc/netdisk/common/xorm"
	"lc/netdisk/internal/svc"
	"lc/netdisk/internal/types"
	"lc/netdisk/model"
	"math"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
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
	var (
		userId = l.ctx.Value(constant.UserIdKey).(int64)
		ext    = req.Ext
		engine = l.svcCtx.Xorm
		has    bool
		err    error
		fileFs model.FileFs
		resp   *types.CheckFileResp
	)

	if has, err = engine.Where("hash = ?", req.Hash).
		And("ext = ?", ext).Get(&fileFs); err != nil {
		return nil, err
	}

	// 文件不存在时
	if !has {
		res, err := engine.DoTransaction(l.createFsAndFileRecord(req))
		return res.(*types.CheckFileResp), err
	}

	// 文件存在时
	if has {
		// 判断该用户是否上传过
		var file model.File
		if has, err = engine.Where("user_id = ?", userId).
			And("fs_id = ?", fileFs.Id).Get(&file); err != nil {
			return nil, err
		}

		// 用户上传过，提示并返回
		if !has {
			// 用户未上传，信息落库
			if fileFs.Size > constant.ShardingSizeFloor {
				file.IsBig = constant.BigFileFlag
			}
			file.Id = idgen.NextId()
			file.UserId = userId
			file.FsId = fileFs.Id
			file.Name = req.Name + ext
			file.FolderId = req.FolderId
			file.Status = constant.StatusFileUploaded
			file.Url = fileFs.Url
			file.DoneAt = time.Now().Local()
			file.DelFlag = constant.StatusFileUndeleted
			if _, err = engine.Insert(&file); err != nil {
				return nil, err
			}
		}
		resp = &types.CheckFileResp{
			FileId: file.Id,
			Status: 1,
		}
	}

	return resp, nil
}

// 创建实际存储和用户存储记录
func (l *CheckFileLogic) createFsAndFileRecord(req *types.CheckFileReq) xorm.TxFn {
	return func(session *xorm.Session) (interface{}, error) {
		var (
			userId = l.ctx.Value(constant.UserIdKey).(int64)
			status = constant.StatusFsFileUnuploaded
			isBig  = constant.SmallFileFlag
			fsId   int64
			err    error
		)

		if req.Size > constant.ShardingSizeFloor {
			status = constant.StatusFsBigFileUnuploaded
			isBig = constant.BigFileFlag
		}

		chunkNum := math.Ceil(float64(req.Size / constant.ShardingSizeFloor))
		fileFs := &model.FileFs{
			Bucket:   l.svcCtx.Minio.BucketName,
			Ext:      req.Ext,
			Hash:     req.Hash,
			Size:     req.Size,
			ChunkNum: int64(chunkNum),
			Status:   status,
		}
		if fsId, err = session.Insert(fileFs); err != nil {
			return nil, err
		}

		fileId := idgen.NextId()
		file := &model.File{
			Model: model.Model{
				Id: fileId,
			},
			UserId:   userId,
			FsId:     fsId,
			FolderId: req.FolderId,
			IsBig:    isBig,
		}
		if _, err = session.Insert(file); err != nil {
			return nil, err
		}

		resp := &types.CheckFileResp{
			FileId: fileId,
			Status: status,
		}
		return resp, nil
	}
}
