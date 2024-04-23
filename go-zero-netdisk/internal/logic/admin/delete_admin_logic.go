package admin

import (
	"context"
	"errors"
	"lc/netdisk/internal/logic/mqs"
	"lc/netdisk/internal/svc"
	"lc/netdisk/internal/types"
	"lc/netdisk/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteAdminLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteAdminLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteAdminLogic {
	return &DeleteAdminLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteAdminLogic) DeleteAdmin(req *types.IdStrReq) error {
	var (
		engine = l.svcCtx.Xorm
		err    error
	)

	defer mqs.LogSend(l.ctx, err, "DeleteAdmin", req.Id)

	if _, err = engine.Where("id = ?", req.Id).
		Delete(&model.Admin{}); err != nil {
		err = errors.New("DeleteAdmin，删除失败：" + err.Error())
		return err
	}

	return nil
}
