package file

import (
	"context"

	"lc/netdisk/internal/svc"
	"lc/netdisk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RecoverFoldersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRecoverFoldersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RecoverFoldersLogic {
	return &RecoverFoldersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RecoverFoldersLogic) RecoverFolders(req *types.RecoverReq) error {
	// todo: add your logic here and delete this line

	return nil
}
