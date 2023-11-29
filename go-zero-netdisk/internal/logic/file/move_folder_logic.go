package file

import (
	"context"

	"lc/netdisk/internal/svc"
	"lc/netdisk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MoveFolderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMoveFolderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MoveFolderLogic {
	return &MoveFolderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MoveFolderLogic) MoveFolder(req *types.MoveFolderReq) error {
	// todo: add your logic here and delete this line

	return nil
}
