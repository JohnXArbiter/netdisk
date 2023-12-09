package download

import (
	"context"
	"lc/netdisk/common/constant"
	"lc/netdisk/internal/svc"
	"lc/netdisk/internal/types"
	"lc/netdisk/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type DownloadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDownloadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DownloadLogic {
	return &DownloadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DownloadLogic) Download(req *types.DownloadReq) (string, error) {
	var (
		userId       = l.ctx.Value(constant.UserIdKey).(int64)
		engine       = l.svcCtx.Xorm
		minioService = l.svcCtx.Minio.NewService()
		fileNetdisk  model.File
		fileFs       model.FileFs
	)

	//key := redis.DownloadGetFsKey + strconv.FormatInt(req.FileNetdiskId, 10)
	//fsMap, err := l.svcCtx.Redis.HGetAll(l.ctx, key).Result()
	//if err == nil && fsMap != nil {
	//	if _, ok := fsMap["id"]; ok {
	//		return minioService.DownloadFile(l.ctx, fsMap["objectName"])
	//	}
	//}

	if has, err := engine.ID(req.FileId).
		And("user_id = ?", userId).
		Get(&fileNetdisk); err != nil || !has {
		return "", err
	}

	if has, err := engine.ID(fileNetdisk.FsId).
		Get(&fileFs); err != nil || !has {
		return "", err
	}

	filename, err := minioService.DownloadFile(l.ctx, fileFs.ObjectName)
	if err != nil {
		return "", err
	}

	return filename, nil
}
