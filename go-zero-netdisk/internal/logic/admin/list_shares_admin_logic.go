package admin

import (
	"context"
	"lc/netdisk/model"

	"lc/netdisk/internal/svc"
	"lc/netdisk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListSharesAdminLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListSharesAdminLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListSharesAdminLogic {
	return &ListSharesAdminLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListSharesAdminLogic) ListSharesAdmin(req *types.PageReq) ([]*model.Share, error) {
	var (
		engine = l.svcCtx.Xorm
		shares []*model.Share
	)

	offset := int((req.Page - 1) * req.Size)
	if err := engine.Limit(int(req.Size), offset).Find(&shares); err != nil {
		logx.Errorf("获取分享列表，查询shares失败，ERR: [%v]", err)
		return nil, err
	}

	return shares, nil
}
