package file

import (
	"context"
	"lc/netdisk/common/constant"
	"lc/netdisk/model"

	"lc/netdisk/internal/svc"
	"lc/netdisk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListDeletedItemsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListDeletedItemsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListDeletedItemsLogic {
	return &ListDeletedItemsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListDeletedItemsLogic) ListDeletedItems() (resp *types.ListDeletedItemsResp, err error) {
	var (
		userId  = l.ctx.Value(constant.UserIdKey).(int64)
		engine  = l.svcCtx.Xorm
		folders []*model.Folder
		files   []*model.File
	)

	if err = engine.Cols("id", "name", "del_time").Where("user_id = ?", userId).
		And("del_flag = ?", constant.StatusFolderDeleted).Find(&folders); err != nil {
		return nil, err
	}

	if err = engine.Cols("id", "name", "url", "del_time").Where("user_id = ?", userId).
		And("del_flag = ?", constant.StatusFileDeleted).Find(&files); err != nil {
		return nil, err
	}

	for _, folder := range folders {
		resp.Folders = append(resp.Folders, &types.ListDeletedFolderStruct{
			Id:      folder.Id,
			Name:    folder.Name,
			DelTime: folder.DelTime,
		})
	}

	for _, file := range files {
		resp.Files = append(resp.Files, &types.ListDeletedFileStruct{
			Id:      file.Id,
			Name:    file.Name,
			Url:     file.Url,
			DelTime: file.DelTime,
		})
	}

	return
}
