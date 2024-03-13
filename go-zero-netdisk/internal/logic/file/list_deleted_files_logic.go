package file

import (
	"context"
	"lc/netdisk/common/constant"
	"lc/netdisk/model"
	"log"

	"lc/netdisk/internal/svc"
	"lc/netdisk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListDeletedFilesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListDeletedFilesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListDeletedFilesLogic {
	return &ListDeletedFilesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListDeletedFilesLogic) ListDeletedFiles() (resp []*types.DeletedFilesResp, err error) {
	var (
		userId  = l.ctx.Value(constant.UserIdKey).(int64)
		engine  = l.svcCtx.Xorm
		folders []*model.Folder
		files   []*model.File
	)

	if err = engine.Where("user_id = ?", userId).
		And("del_flag = ?", constant.StatusFileDeleted).
		Asc("del_time").Find(&files); err != nil {
		return nil, err
	}

	log.Println(len(files), userId)

	var folderIds []int64
	m := make(map[int64]string)
	for _, file := range files {
		if file.FolderId != 0 {
			if _, ok := m[file.FolderId]; !ok {
				folderIds = append(folderIds, file.FolderId)
			} else {
				m[file.FolderId] = ""
			}
		}
	}

	if err = engine.Select("id, name").In("id", folderIds).
		//And("del_flag = ?", constant.StatusFolderUndeleted).
		Find(&folders); err != nil {
		return nil, err
	}

	m = map[int64]string{}
	for _, folder := range folders {
		m[folder.Id] = folder.Name
	}

	for _, file := range files {
		resp = append(resp, &types.DeletedFilesResp{
			Id:         file.Id,
			Name:       file.Name,
			Url:        file.Url,
			Status:     file.Status,
			Size:       file.Size,
			FolderId:   file.FolderId,
			FolderName: m[file.FolderId],
			DelTime:    file.DelTime,
		})
	}

	return
}
