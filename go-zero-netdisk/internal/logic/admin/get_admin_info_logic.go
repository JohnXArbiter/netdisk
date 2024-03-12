package admin

import (
	"context"
	"errors"
	"lc/netdisk/common/constant"
	"lc/netdisk/model"

	"github.com/zeromicro/go-zero/core/logx"
	"lc/netdisk/internal/svc"
)

type GetAdminInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAdminInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAdminInfoLogic {
	return &GetAdminInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAdminInfoLogic) GetAdminInfo() (interface{}, error) {
	var (
		adminId = l.ctx.Value(constant.UserIdKey).(int64)
		engine  = l.svcCtx.Xorm
	)

	var admin model.Admin
	if has, err := engine.ID(adminId).
		Get(&admin); err != nil {
		logx.Errorf("GetAdminInfo，获取失败，ERR: [%v]", err)
		return nil, err
	} else if !has {
		return nil, errors.New("信息获取失败")
	}

	return admin, nil
}
