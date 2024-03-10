package file

import (
	"context"
	"lc/netdisk/common/constant"
	"lc/netdisk/model"
	"time"

	"lc/netdisk/internal/svc"
	"lc/netdisk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListSharesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListSharesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListSharesLogic {
	return &ListSharesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListSharesLogic) ListShares() (resp []*types.ListShareStruct, err error) {
	var (
		userId = l.ctx.Value(constant.UserIdKey).(int64)
		engine = l.svcCtx.Xorm
		shares []*model.Share
	)

	if err = engine.Where("user_id = ?", userId).Find(&shares); err != nil {
		logx.Errorf("获取分享列表失败，ERR: [%v]", err)
		return nil, err
	}

	var expiredShares []string
	for _, share := range shares {
		status := share.Status
		if share.Status == constant.StatusShareNotExpired &&
			time.Now().Unix() > share.Expired {
			expiredShares = append(expiredShares, share.Id)
			status = constant.StatusShareExpired
		}
		resp = append(resp, &types.ListShareStruct{
			Id:          share.Id,
			Pwd:         share.Pwd,
			Name:        share.Name,
			Created:     share.Created.Format(constant.TimeFormat1),
			Expired:     share.Expired,
			Status:      status,
			DownloadNum: share.DownloadNum,
			ClickNum:    share.ClickNum,
			Type:        share.Type,
		})
	}

	if len(expiredShares) > 0 {
		if _, err = engine.In("id", expiredShares).Update(&model.Share{
			Status: constant.StatusShareExpired}); err != nil {
			logx.Errorf("获取分享列表，更新 [%v] 过期状态失败，ERR: [%v]",
				expiredShares, err)
		}
	}

	return
}
