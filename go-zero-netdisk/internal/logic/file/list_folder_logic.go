package file

import (
	"context"

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

func (l *ListFolderLogic) ListFolder(req *types.ListFileFolderReq) error {
	// todo: add your logic here and delete this line

	return nil
}
