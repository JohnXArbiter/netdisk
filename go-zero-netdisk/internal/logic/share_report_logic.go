package logic

import (
	"context"
	"errors"
	"lc/netdisk/common/constant"
	"lc/netdisk/internal/logic/mqs"
	"lc/netdisk/model"

	"lc/netdisk/internal/svc"
	"lc/netdisk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShareReportLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareReportLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareReportLogic {
	return &ShareReportLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareReportLogic) ShareReport(req *types.ReportReq) error {
	var (
		engine = l.svcCtx.Xorm
		err    error
	)

	defer mqs.LogSend(l.ctx, err, "ShareReport", req.ShareId)

	bean := &model.Share{
		Status: constant.StatusShareIllegal,
		Reason: req.Reason,
	}
	if _, err = engine.ID(req.ShareId).
		Update(bean); err != nil {
		err = errors.New("ShareReport，举报分享: %v 失败，ERR: " + err.Error())
		logx.Error(err)
		return err
	}

	return nil
}
