package file

import (
	"context"
	"lc/netdisk/common/constant"
	"lc/netdisk/model"

	"lc/netdisk/internal/svc"
	"lc/netdisk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DownloadFolderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDownloadFolderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DownloadFolderLogic {
	return &DownloadFolderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DownloadFolderLogic) DownloadFolder(req *types.DFReq) (resp []string, err error) {
	var (
		minioSvc = l.svcCtx.MinioSvc
		files    []*model.File
		engine   = l.svcCtx.Xorm
	)

	if err = engine.In("folder_id", req.FolderIds).
		Find(&files); err != nil {
		return nil, err
	}

	for _, file := range files {
		if file.Status != constant.StatusFileUploaded {
			continue
		}
		url, err2 := minioSvc.GenUrl(file.ObjectName, file.Name, true)
		if err2 != nil {
			logx.Errorf("通过文件夹id获取文件列表，[%d]获取url失败，ERR: [%v]", file.Id, err2)
			continue
		}
		resp = append(resp, url)
	}

	return
}
