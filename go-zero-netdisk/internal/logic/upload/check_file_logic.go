package upload

import (
	"context"
	"github.com/yitter/idgenerator-go/idgen"
	"lc/netdisk/common"
	"lc/netdisk/common/xorm"
	"lc/netdisk/internal/svc"
	"lc/netdisk/internal/types"
	"lc/netdisk/model"
	"strconv"
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
		userId = l.ctx.Value(common.UserIdKey).(int64)
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
		data, err := engine.DoTransaction(l.createFsAndNetdiskRecord(req))
		if err != nil {
			return nil, err
		}
		return data.(*types.CheckFileResp), nil
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
			FileId: fileNetdisk.Id,
			Status: 1,
		}
	}

	return resp, nil
}

// 创建实际存储和用户存储记录
func (l *CheckFileLogic) createFsAndNetdiskRecord(req *types.CheckFileReq) xorm.TxFn {
	return func(session *xorm.Session) (interface{}, error) {
		var (
			userId = l.ctx.Value(common.UserIdKey).(int64)
			status int8
			fsId   int64
			err    error
		)

		if req.Size > common.NeedShardingSize {
			status = -2
		}

		name := req.Name + "|" + strconv.FormatInt(time.Now().Unix(), 10) + req.Ext
		objectName := l.svcCtx.Minio.GenObjectName(req.Hash, name)
		fileFs := &model.FileFs{
			Bucket:     l.svcCtx.Minio.BucketName,
			Ext:        req.Ext,
			ObjectName: objectName,
			Hash:       req.Hash,
			Name:       name,
			Size:       req.Size,
			Url:        "",
			Status:     status,
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
			Name:     req.Name + req.Ext,
		}
		if _, err = session.Insert(netdisk); err != nil {
			return nil, err
		}

		resp := &types.CheckFileResp{
			FileId: netdiskId,
			Status: status,
		}
		return resp, nil
	}
}
