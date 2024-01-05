package file

import (
	"context"
	"errors"
	"lc/netdisk/common/constant"
	"lc/netdisk/model"

	"lc/netdisk/internal/svc"
	"lc/netdisk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListFolderMovableFolderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListFolderMovableFolderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListFolderMovableFolderLogic {
	return &ListFolderMovableFolderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListFolderMovableFolderLogic) ListFolderMovableFolder(req *types.ParentFolderIdReq) ([]*types.ListFolderStruct, error) {
	var (
		userId  = l.ctx.Value(constant.UserIdKey).(int64)
		engine  = l.svcCtx.Xorm
		folders []*model.Folder
	)

	if err := engine.Cols("id", "name").Where("user_id = ?", userId).
		And("id != ?", req.ParentFolderId).Find(&folders); err != nil {
		return nil, errors.New("出错了" + err.Error())
	}

	var resp []*types.ListFolderStruct
	for _, folder := range folders {
		lfs := &types.ListFolderStruct{}
		lfs.Id = folder.Id
		lfs.Name = folder.Name
		resp = append(resp, lfs)
	}
	return resp, nil
}
