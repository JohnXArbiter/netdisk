package admin

import (
	"context"
	"errors"
	"lc/netdisk/model"

	"lc/netdisk/internal/svc"
	"lc/netdisk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUrlLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUrlLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUrlLogic {
	return &GetUrlLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUrlLogic) GetUrl(req *types.IdPathReq) (string, error) {
	var (
		engine    = l.svcCtx.Xorm
		minioSvc  = l.svcCtx.Minio.NewService()
		shareFile model.ShareFile
		file      model.File
	)

	var objectName string
	if has, err := engine.Table(&model.Share{}).Alias("a").Select("c.object_name").
		Join("LEFT", []string{shareFile.TableName(), "b"}, "a.id = b.share_id").
		Join("LEFT", []string{file.TableName(), "c"}, "c.id = b.file_id").
		Where("a.id = ?", req.Id).Get(&objectName); err != nil {
		logx.Errorf("GetUrl，获取objectName失败，ERR: [%v]", err)
		return "", err
	} else if !has {
		return "", errors.New("没有找到文件链接")
	}

	url, err := minioSvc.GenUrl(objectName, true)
	if err != nil {
		return "", err
	}

	return url, nil
}
