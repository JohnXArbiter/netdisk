package admin

import (
	"context"
	"errors"
	"fmt"
	"lc/netdisk/internal/logic/mqs"
	"lc/netdisk/model"

	"lc/netdisk/internal/svc"
	"lc/netdisk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetAdminStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetAdminStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetAdminStatusLogic {
	return &SetAdminStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetAdminStatusLogic) SetAdminStatus(req *types.SetStatusReq) error {
	var (
		engine = l.svcCtx.Xorm
		err    error
	)

	defer mqs.LogSend(l.ctx, err, "SetAdminStatus", req.Id, req.Status)

	bean := &model.Admin{Status: req.Status}
	if _, err = engine.ID(req.Id).
		Cols("status").
		Update(bean); err != nil {
		err = errors.New(fmt.Sprintf("SetAdminStatus，更新失败，ERR: [%v]", err.Error()))
		logx.Error(err)
		return err
	}
	return nil
}
