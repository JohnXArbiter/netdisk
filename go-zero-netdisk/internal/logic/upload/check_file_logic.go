package upload

import (
	"context"
	"github.com/yitter/idgenerator-go/idgen"
	"lc/netdisk/common/constant"
	"lc/netdisk/common/xorm"
	"lc/netdisk/internal/svc"
	"lc/netdisk/internal/types"
	"lc/netdisk/model"
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
		res, err := engine.DoTransaction(l.createFsAndNetdiskRecord(req))
		return res.(*types.CheckFileResp), err
	}

	// 文件存在时
	if has {
		// 判断该用户是否上传过
		var fileNetdisk model.FileNetdisk
		if has, err = engine.Where("user_id = ?", userId).
			And("fs_id = ?", fileFs.Id).Get(&fileNetdisk); err != nil {
			return nil, err
		}

		// 用户上传过，提示并返回
		if !has {
			// 用户未上传，信息落库
			fileNetdisk.Id = idgen.NextId()
			fileNetdisk.UserId = userId
			fileNetdisk.FsId = fileFs.Id
			fileNetdisk.Name = req.Name + ext
			fileNetdisk.FolderId = req.FolderId
			fileNetdisk.Status = 1
			fileNetdisk.Url = fileFs.Url
			fileNetdisk.DoneAt = time.Now().Local()
			if _, err = engine.Insert(&fileNetdisk); err != nil {
				return nil, err
			}
		}
		resp = &types.CheckFileResp{
			FileNetdiskId: fileNetdisk.Id,
			Status:        1,
		}
	}

	return resp, nil
}

// 创建实际存储和用户存储记录
func (l *CheckFileLogic) createFsAndNetdiskRecord(req *types.CheckFileReq) xorm.TxFn {
	return func(session *xorm.Session) (interface{}, error) {
		var (
			userId = l.ctx.Value(constant.UserIdKey).(int64)
			status = constant.StatusFsFileUnuploaded
			fsId   int64
			err    error
		)

		if req.Size > constant.NeedShardingSize {
			status = constant.StatusFsBigFileUnuploaded
		}

		fileFs := &model.FileFs{
			Bucket: l.svcCtx.Minio.BucketName,
			Ext:    req.Ext,
			Hash:   req.Hash,
			Size:   req.Size,
			Status: status,
		}
		if fsId, err = session.Insert(fileFs); err != nil {
			return nil, err
		}

		netdiskId := idgen.NextId()
		netdisk := &model.FileNetdisk{
			Model: model.Model{
				Id: netdiskId,
			},
			UserId:   userId,
			FsId:     fsId,
			FolderId: req.FolderId,
		}
		if _, err = session.Insert(netdisk); err != nil {
			return nil, err
		}

		resp := &types.CheckFileResp{
			FileNetdiskId: netdiskId,
			Status:        status,
		}
		return resp, nil
	}
}
