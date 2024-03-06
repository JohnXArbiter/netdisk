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

type ListShareFilesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListShareFilesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListShareFilesLogic {
	return &ListShareFilesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListShareFilesLogic) ListShareFiles() (resp []*types.ListShareStruct, err error) {
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
		if share.Expired < time.Now().Unix() {
			expiredShares = append(expiredShares, share.Id)
			status = constant.ShareExpired
		}
		resp = append(resp, &types.ListShareStruct{
			Id:          share.Id,
			Name:        share.Name,
			Created:     share.Created,
			Expired:     share.Expired,
			Status:      status,
			DownloadNum: share.DownloadNum,
			ClickNum:    share.ClickNum,
		})
	}

	return
}
