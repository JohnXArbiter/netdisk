package admin

import (
	"context"
	"lc/netdisk/common/redis"
	"lc/netdisk/model"

	"github.com/zeromicro/go-zero/core/logx"
	"lc/netdisk/internal/svc"
)

type (
	StatisticLogic struct {
		logx.Logger
		ctx    context.Context
		svcCtx *svc.ServiceContext
	}

	TypeCnt struct {
		Type int8  `json:"type"`
		Cnt  int64 `json:"cnt"`
	}
)

func NewStatisticLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StatisticLogic {
	return &StatisticLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StatisticLogic) Statistic() (map[string]interface{}, error) {
	var (
		engine = l.svcCtx.Xorm
		rdb    = l.svcCtx.Redis
		user   = &model.User{}
		share  = &model.Share{}
		file   = &model.File{}
		folder = &model.Folder{}
		fs     = &model.FileFs{}
		resp   map[string]interface{}
	)

	userCnt, err := engine.Count(user)
	if err != nil {
		return nil, err
	}

	used, err := engine.SumInt(user, "used")
	if err != nil {
		return nil, err
	}

	capacity, err := rdb.Get(l.ctx, redis.NetdiskCapacity).Result()
	if err != nil {
		return nil, err
	}

	shareSums, err := engine.SumsInt(share, "download_num", "click_num")
	if err != nil {
		return nil, err
	}

	var typeCnt []*TypeCnt
	if err = engine.Select("type, count(*) as cnt").
		GroupBy("type").
		Table(file).
		Find(&typeCnt); err != nil {
		return nil, err
	}

	fileCnt, err := engine.Count(file)
	if err != nil {
		return nil, err
	}

	folderCnt, err := engine.Count(folder)
	if err != nil {
		return nil, err
	}

	fsCnt, err := engine.Count(fs)
	if err != nil {
		return nil, err
	}

	resp = map[string]interface{}{
		"cnt": map[string]interface{}{
			"user":     userCnt,
			"download": shareSums[0],
			"click":    shareSums[1],
			"file": map[string]interface{}{
				"file":   fileCnt,
				"type":   typeCnt,
				"folder": folderCnt,
				"fs":     fsCnt,
			},
		},
		"cap": map[string]interface{}{
			"used":     used,
			"capacity": capacity,
		},
	}
	return resp, nil
}
