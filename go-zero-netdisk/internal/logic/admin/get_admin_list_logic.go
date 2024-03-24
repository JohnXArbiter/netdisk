package admin

import (
	"context"
	"lc/netdisk/internal/types"
	"lc/netdisk/model"

	"github.com/zeromicro/go-zero/core/logx"
	"lc/netdisk/internal/svc"
)

type GetAdminListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAdminListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAdminListLogic {
	return &GetAdminListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAdminListLogic) GetAdminList(req *types.PageReq) (interface{}, error) {
	var (
		engine = l.svcCtx.Xorm
		admins []*model.Admin
	)

	offset := int((req.Page - 1) * req.Size)
	if err := engine.Limit(int(req.Size), offset).Find(&admins); err != nil {
		logx.Errorf("获取用户列表，查询users失败，ERR: [%v]", err)
		return nil, err
	}

	return admins, nil
}
