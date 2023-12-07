package upload

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"lc/netdisk/common/constant"
	"lc/netdisk/common/xorm"
	"lc/netdisk/internal/svc"
	"lc/netdisk/internal/types"
	"lc/netdisk/model"
)

type UploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadLogic {
	return &UploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadLogic) Upload(req *types.UploadReq, fileParam *types.FileParam) (interface{}, error) {
	var (
		userId      = l.ctx.Value(constant.UserIdKey).(int64)
		engine      = l.svcCtx.Xorm
		fileNetdisk model.FileNetdisk
		fileFs      model.FileFs
		has         bool
		err         error
	)

	if has, err = engine.ID(req.FileNetdiskId).And("user_id = ?", userId).
		Get(&fileNetdisk); err != nil {
		return nil, err
	} else if !has {
		return nil, errors.New("文件上传发生错误！")
	}

	if fileNetdisk.Status == constant.StatusNetdiskUploaded {
		return nil, nil
	}

	if has, err = engine.ID(fileNetdisk.FsId).Get(&fileFs); err != nil {
		return nil, err
	} else if !has {
		return nil, errors.New("文件上传发生错误！")
	}

	if _, err = engine.DoTransaction(l.uploadAndUpdateFsAndNetdiskRecord(
		&fileFs, &fileNetdisk, fileParam)); err != nil {
		return nil, err
	}

	return nil, nil
}

func (l *UploadLogic) uploadAndUpdateFsAndNetdiskRecord(fileFs *model.FileFs,
	fileNetdisk *model.FileNetdisk, fileParam *types.FileParam) xorm.TxFn {
	return func(session *xorm.Session) (interface{}, error) {
		var (
			minioService = l.svcCtx.Minio.NewService()
			header       = fileParam.FileHeader
			file         = fileParam.File
			err          error
		)

		filename, objectName := l.svcCtx.Minio.GenObjectName(fileFs.Hash, header.Filename)

		fileNetdisk.Name = filename
		fileNetdisk.Url = ""
		fileNetdisk.Status = constant.StatusNetdiskUploaded
		if _, err = session.Insert(fileNetdisk); err != nil {
			return nil, err
		}

		fileFs.Name = filename
		fileFs.ObjectName = objectName
		fileFs.Url = ""
		fileFs.Status = constant.StatusFsUploaded
		if _, err = session.Insert(fileNetdisk); err != nil {
			return nil, err
		}

		if err = minioService.UploadFile(l.ctx, objectName, file); err != nil {
			return nil, err
		}
		return nil, nil
	}
}
