package file

import (
	"context"

	"lc/netdisk/internal/svc"
	"lc/netdisk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListFileByTypeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListFileByTypeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListFileByTypeLogic {
	return &ListFileByTypeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListFileByTypeLogic) ListFileByType(req *types.ParentFolderIdReq) (resp []*types.ListFileStruct, err error) {
	// todo: add your logic here and delete this line

	return
}
