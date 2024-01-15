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

func (l *ListFolderLogic) ListFolder(req *types.ParentFolderIdReq) ([]*types.ListFolderStruct, error) {
	var (
		userId  = l.ctx.Value(constant.UserIdKey).(int64)
		engine  = l.svcCtx.Xorm
		folders []*model.Folder
		resp    []*types.ListFolderStruct
	)

	if err := engine.Where("parent_id = ?", req.ParentFolderId).
		And("user_id = ?", userId).And("del_flag = ?",
		constant.StatusFolderUndeleted).Find(&folders); err != nil {
		return nil, err
	}

	for _, folder := range folders {
		resp = append(resp, &types.ListFolderStruct{
			Id:      folder.Id,
			Name:    folder.Name,
			Updated: folder.Updated.Format(constant.TimeFormat1),
		})
	}

	return resp, nil
}
