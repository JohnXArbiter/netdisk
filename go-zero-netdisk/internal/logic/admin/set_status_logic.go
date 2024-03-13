package admin

import (
	"context"
	"lc/netdisk/model"

	"lc/netdisk/internal/svc"
	"lc/netdisk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetStatusLogic {
	return &SetStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetStatusLogic) SetStatus(req *types.SetStatusReq) error {
	var (
		engine = l.svcCtx.Xorm
	)

	bean := &model.User{Status: req.Status}
	if _, err := engine.ID(req.Id).
		Cols("status").
		Update(bean); err != nil {
		logx.Errorf("SetStatus，更新失败，ERR: [%v]", err)
		return err
	}

	return nil
}
