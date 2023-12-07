package user

import (
	"context"
	"errors"
	"lc/netdisk/model"

	"lc/netdisk/internal/svc"
	"lc/netdisk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDetailLogic {
	return &GetDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetDetail TODO
func (l *GetDetailLogic) GetDetail(req *types.GetUserDetailReq) (interface{}, error) {
	var (
		engine = l.svcCtx.Xorm
		user   model.User
	)

	has, err := engine.ID(req.UserId).Get(&user)
	if err != nil {
		return nil, err
	}

	if !has {
		return nil, errors.New("未找到该用户信息")
	}

	return user, nil
}
