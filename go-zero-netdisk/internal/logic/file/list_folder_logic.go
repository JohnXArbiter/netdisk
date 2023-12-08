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

func (l *ListFolderLogic) ListFolder(req *types.ListFileFolderReq) (*types.ListFileFolderResp, error) {
	var (
		userId       = l.ctx.Value(constant.UserIdKey).(int64)
		engine       = l.svcCtx.Xorm
		folders      []*model.Folder
		fileNetdisks []*model.FileNetdisk
		resp         types.ListFileFolderResp
	)

	if err := engine.Where("parent_id = ?", req.ParentFolderId).
		And("user_id = ?", userId).Find(&folders); err != nil {
		return nil, err
	}

	if err := engine.Where("folder_id = ?", req.ParentFolderId).
		And("user_id = ?", userId).Find(&fileNetdisks); err != nil {
		return nil, err
	}

	for _, folder := range folders {
		resp.Folders = append(resp.Folders, &types.ListFoldersStruct{
			Id:   folder.Id,
			Name: folder.Name,
		})
	}

	for _, fileNetdisk := range fileNetdisks {
		resp.FileNetdisks = append(resp.FileNetdisks, &types.ListFileStruct{
			Id:     fileNetdisk.Id,
			Name:   fileNetdisk.Name,
			Url:    fileNetdisk.Url,
			Status: fileNetdisk.Status,
		})
	}

	return &resp, nil
}
