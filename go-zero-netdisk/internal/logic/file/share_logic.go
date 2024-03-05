package file

import (
	"context"

	"lc/netdisk/internal/svc"
	"lc/netdisk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShareLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareLogic {
	return &ShareLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareLogic) Share(req *types.ShareReq) error {
	// todo: add your logic here and delete this line

	return nil
}
