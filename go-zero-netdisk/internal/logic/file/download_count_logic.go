package file

import (
	"context"
	"lc/netdisk/model"

	"lc/netdisk/internal/svc"
	"lc/netdisk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DownloadCountLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDownloadCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DownloadCountLogic {
	return &DownloadCountLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DownloadCountLogic) DownloadCount(req *types.IdPathReq) error {
	var (
		engine = l.svcCtx.Xorm
	)

	if _, err := engine.SetExpr("download_num", "download_num + 1").
		ID(req.Id).Update(&model.Share{}); err != nil {
		logx.Errorf("获取分享文件列表，生成文件url出错：ERR: [%v]", err)
		return err
	}
	return nil
}
