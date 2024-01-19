package user

import (
	"context"
	"errors"
	"lc/netdisk/common/constant"
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

func (l *GetDetailLogic) GetDetail(req *types.IdPathReq) (interface{}, error) {
	var (
		loginUserId = l.ctx.Value(constant.UserIdKey).(int64)
		engine      = l.svcCtx.Xorm
		user        model.User
	)

	targetUserId := req.Id
	if req.Id == 0 {
		targetUserId = loginUserId
	}

	cols := "id, name, avatar, email, signature, status"
	has, err := engine.Select(cols).ID(targetUserId).Get(&user)
	if err != nil {
		return nil, err
	}

	if !has {
		return nil, errors.New("未找到该用户信息")
	}
	return user, nil
}
