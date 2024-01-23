package upload

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"lc/netdisk/common/constant"
	"lc/netdisk/common/redis"
	"lc/netdisk/common/xorm"
	"lc/netdisk/internal/svc"
	"lc/netdisk/internal/types"
	"lc/netdisk/model"
	"strconv"
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
		userId = l.ctx.Value(constant.UserIdKey).(int64)
		engine = l.svcCtx.Xorm
		rdb    = l.svcCtx.Redis
	)

	fileIdStr := strconv.FormatInt(req.FileId, 10)
	key := redis.UploadCheckKey + fileIdStr
	fileInfo, err := rdb.HGetAll(l.ctx, key).Result()
	if err != nil {
		return nil, err
	}

	if fileInfo["userId"] != strconv.FormatInt(userId, 10) {
		return nil, errors.New("信息有误1")
	}
	if fileInfo["name"] != fileParam.FileHeader.Filename {
		return nil, errors.New("信息有误2")
	}

	folderId := fileInfo["folderId"]
	if folderId != "0" {
		if has, err := engine.Where("id = ?", folderId).
			Exist(&model.FileFs{}); err != nil {
			return nil, err
		} else if !has {
			return nil, errors.New("信息有误3")
		}
	}

	_, err = engine.DoTransaction(l.store2dbAndFs(fileInfo, fileParam))
	if err != nil {
		return nil, err
	}

	go rdb.Del(l.ctx, key)
	return nil, nil
}

func (l *UploadLogic) store2dbAndFs(fileInfo map[string]string, fileParam *types.FileParam) xorm.TxFn {
	return func(session *xorm.Session) (interface{}, error) {
		var (
			minioSvc = l.svcCtx.Minio.NewService()
			header   = fileParam.FileHeader
			fileData = fileParam.File
			err      error
		)

		filename, objectName := l.svcCtx.Minio.GenObjectName(fileInfo["hash"], header.Filename)

		file := &model.File{}
		file.Name = filename
		file.Url = ""
		file.Status = constant.StatusFileUploaded
		if _, err = session.Insert(file); err != nil {
			return nil, err
		}

		fileFs := &model.FileFs{}
		fileFs.Name = filename
		fileFs.ObjectName = objectName
		fileFs.Url = ""
		fileFs.Status = constant.StatusFsUploaded
		if _, err = session.Insert(file); err != nil {
			return nil, err
		}

		if err = minioSvc.UploadFile(l.ctx, objectName, fileData); err != nil {
			return nil, err
		}
		return nil, nil
	}
}
