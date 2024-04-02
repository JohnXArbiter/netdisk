package admin

import (
	"context"
	"lc/netdisk/internal/logic/mqs"
	"lc/netdisk/model"

	"lc/netdisk/internal/svc"
	"lc/netdisk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetFileStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetFileStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetFileStatusLogic {
	return &SetFileStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetFileStatusLogic) SetFileStatus(req *types.SetFileStatusReq) error {
	var (
		engine = l.svcCtx.Xorm
		err    error
	)

	defer mqs.LogSend(l.ctx, err, "SetFileStatus", req.Ids, req.Status)

	bean := &model.File{Status: req.Status}
	if _, err = engine.In("id", req.Ids).
		Cols("status").
		Update(bean); err != nil {
		logx.Errorf("SetFileStatus，更新失败，ERR: [%v]", err)
		return err
	}

	return nil
}
