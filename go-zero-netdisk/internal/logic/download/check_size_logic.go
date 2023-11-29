package download

import (
	"context"

	"lc/netdisk/internal/svc"
	"lc/netdisk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckSizeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckSizeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckSizeLogic {
	return &CheckSizeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckSizeLogic) CheckSize(req *types.CheckSizeReq) error {
	// todo: add your logic here and delete this line

	return nil
}
