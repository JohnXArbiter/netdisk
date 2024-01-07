package file

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"lc/netdisk/internal/svc"
)

type DeleteFoldersTrulyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteFoldersTrulyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteFoldersTrulyLogic {
	return &DeleteFoldersTrulyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteFoldersTrulyLogic) DeleteFoldersTruly() error {
	// todo: add your logic here and delete this line

	return nil
}
