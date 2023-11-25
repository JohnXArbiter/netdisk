package logic

import (
	"context"

	"netdisk/internal/svc"
	"netdisk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type NetdiskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewNetdiskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NetdiskLogic {
	return &NetdiskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *NetdiskLogic) Netdisk(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
