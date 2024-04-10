package admin

import (
	"context"
	"lc/netdisk/internal/logic/mqs"
	"lc/netdisk/model"

	"lc/netdisk/internal/svc"
	"lc/netdisk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetUserStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetUserStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetUserStatusLogic {
	return &SetUserStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetUserStatusLogic) SetUserStatus(req *types.SetStatusReq) error {
	var (
		engine = l.svcCtx.Xorm
		err    error
	)

	defer mqs.LogSend(l.ctx, err, "SetUserStatus", req.Id, req.Status)

	bean := &model.User{Status: req.Status}
	if _, err = engine.ID(req.Id).
		Cols("status").
		Update(bean); err != nil {
		logx.Errorf("SetStatus，更新失败，ERR: [%v]", err)
		return err
	}

	return nil
}
