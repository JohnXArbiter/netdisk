package file

import (
	"context"
	"lc/netdisk/common/constant"
	"lc/netdisk/model"

	"lc/netdisk/internal/svc"
	"lc/netdisk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListFolderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListFolderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListFolderLogic {
	return &ListFolderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListFolderLogic) ListFolder(req *types.ParentFolderIdReq) (*types.ListFileFolderResp, error) {
	var (
		userId  = l.ctx.Value(constant.UserIdKey).(int64)
		engine  = l.svcCtx.Xorm
		folders []*model.Folder
		files   []*model.File
		resp    types.ListFileFolderResp
	)

	if err := engine.Where("parent_id = ?", req.ParentFolderId).
		And("user_id = ?", userId).And("del_flag = ?",
		constant.StatusFileUndeleted).Find(&folders); err != nil {
		return nil, err
	}

	if err := engine.Where("folder_id = ?", req.ParentFolderId).
		And("user_id = ?", userId).And("del_flag = ?",
		constant.StatusFolderUndeleted).Find(&files); err != nil {
		return nil, err
	}

	for _, folder := range folders {
		resp.Folders = append(resp.Folders, &types.ListFolderStruct{
			Id:      folder.Id,
			Name:    folder.Name,
			Updated: folder.Updated.Format(constant.TimeFormat),
		})
	}

	for _, file := range files {
		resp.Files = append(resp.Files, &types.ListFileStruct{
			Id:      file.Id,
			Name:    file.Name,
			Url:     file.Url,
			Size:    file.Size,
			Status:  file.Status,
			Updated: file.Updated.Format(constant.TimeFormat),
		})
	}

	return &resp, nil
}
