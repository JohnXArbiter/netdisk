package file

import (
	"context"
	"lc/netdisk/common/constant"
	"lc/netdisk/common/xorm"
	"lc/netdisk/model"

	"lc/netdisk/internal/svc"
	"lc/netdisk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShareCancelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareCancelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareCancelLogic {
	return &ShareCancelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareCancelLogic) ShareCancel(req *types.IdStrsReq) error {
	var (
		userId = l.ctx.Value(constant.UserIdKey).(int64)
		engine = l.svcCtx.Xorm
	)

	_, err := engine.DoTransaction(func(session *xorm.Session) (interface{}, error) {
		if _, err := session.In("id", req.Ids).
			And("user_id = ?", userId).
			Delete(&model.Share{}); err != nil {
			logx.Errorf("ShareCancel，删除share记录失败，ERR: [%v]", err)
			return nil, err
		}

		if _, err := session.In("share_id", req.Ids).
			Delete(&model.ShareFile{}); err != nil {
			logx.Errorf("ShareCancel，删除shareFile记录失败，ERR: [%v]", err)
			return nil, err
		}
		return nil, nil
	})

	return err
}
