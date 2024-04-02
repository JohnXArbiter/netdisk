package file

import (
	"context"
	"fmt"
	redis2 "github.com/redis/go-redis/v9"
	"lc/netdisk/common/constant"
	"lc/netdisk/common/redis"
	"lc/netdisk/model"
	"time"

	"lc/netdisk/internal/svc"
	"lc/netdisk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListFileByTypeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListFileByTypeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListFileByTypeLogic {
	return &ListFileByTypeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListFileByTypeLogic) ListFileByType(req *types.FileTypeReq) ([]*types.FileResp, error) {
	var (
		userId   = l.ctx.Value(constant.UserIdKey).(int64)
		engine   = l.svcCtx.Xorm
		rdb      = l.svcCtx.Redis
		minioSvc = l.svcCtx.Minio.NewService()
		key      = fmt.Sprintf(redis.FileTypeDownloadUrlKey, userId, req.FileType)
		files    []*model.File
		resp     []*types.FileResp
	)

	if err := engine.Desc("created").
		Select("id, name, size, object_name, type, status, created, updated").
		Where("type = ?", req.FileType).
		And("user_id = ?", userId).
		And("del_flag = ?", constant.StatusFileUndeleted).
		Find(&files); err != nil {
		return nil, err
	}

	zs, redisErr := rdb.ZRevRangeWithScores(l.ctx, key, 0, -1).Result()
	if redisErr != nil && redisErr != redis2.Nil {
		logx.Errorf("通过类型获取文件列表，redis获取set失败，err: %v", redisErr)
	}

	var urls []redis2.Z
	for i, file := range files {
		var url string
		if len(zs) == len(files) && redisErr == nil {
			url = zs[i].Member.(string)
		} else {
			url2, err := minioSvc.GenUrl(file.ObjectName, file.Name, true)
			if err != nil {
				logx.Errorf("通过类型获取文件列表，[%d]获取url失败，err: %v", file.ObjectName, redisErr)
			} else {
				url = url2
				urls = append(urls, redis2.Z{Member: url, Score: float64(file.Created.Unix())})
				if err = rdb.ZAdd(l.ctx, key, urls...).Err(); err != nil {
					logx.Errorf("通过类型获取文件列表，redis缓存url失败，err: %v", err)
				}
				if err = rdb.Expire(l.ctx, key, 7*24*time.Hour).Err(); err != nil {
					logx.Errorf("通过类型获取文件列表，设置缓存expire失败，ERR: [%v]", err)
				}
			}
		}

		resp = append(resp, &types.FileResp{
			Id:      file.Id,
			Name:    file.Name,
			Url:     url,
			Type:    file.Type,
			Size:    file.Size,
			Status:  file.Status,
			Updated: file.Updated.Format(constant.TimeFormat1),
		})
	}

	if err := rdb.Expire(l.ctx, key, redis.DownloadTypeExpire).Err(); err != nil {
		logx.Errorf("ListFileByType，设置过期时间失败，ERR: [%v]", err)
		return nil, err
	}

	return resp, nil
}
