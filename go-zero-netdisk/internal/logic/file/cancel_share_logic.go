package file

import (
	"context"

	"lc/netdisk/internal/svc"
	"lc/netdisk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CancelShareLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCancelShareLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CancelShareLogic {
	return &CancelShareLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CancelShareLogic) CancelShare(req *types.IdsReq) error {
	// todo: add your logic here and delete this line

	return nil
}
