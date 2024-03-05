package file

import (
	"context"

	"lc/netdisk/internal/svc"
	"lc/netdisk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListShareFilesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListShareFilesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListShareFilesLogic {
	return &ListShareFilesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListShareFilesLogic) ListShareFiles() (resp []*types.ShareFileStruct, err error) {
	// todo: add your logic here and delete this line

	return
}
