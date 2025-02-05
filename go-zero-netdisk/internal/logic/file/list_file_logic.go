package file

import (
	"context"
	"fmt"
	redis2 "github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/logx"
	"lc/netdisk/common/constant"
	"lc/netdisk/common/redis"
	"lc/netdisk/internal/svc"
	"lc/netdisk/internal/types"
	"lc/netdisk/model"
)

type ListFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListFileLogic {
	return &ListFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListFileLogic) ListFile(req *types.ParentFolderIdReq) ([]*types.FileResp, error) {
	var (
		userId   = l.ctx.Value(constant.UserIdKey).(int64)
		engine   = l.svcCtx.Xorm
		rdb      = l.svcCtx.Redis
		minioSvc = l.svcCtx.Minio.NewService()
		key      = fmt.Sprintf(redis.FileFolderDownloadUrlKey, userId, req.ParentFolderId)
		files    []*model.File
		resp     []*types.FileResp
	)

	if err := engine.Desc("created").
		Select("id, name, size, object_name, type, status, created, updated").
		Where("folder_id = ?", req.ParentFolderId).
		And("user_id = ?", userId).
		And("del_flag = ?", constant.StatusFileUndeleted).
		Find(&files); err != nil {
		return nil, err
	}

	zs, redisErr := rdb.ZRevRangeWithScores(l.ctx, key, 0, -1).Result()
	if redisErr != nil && redisErr != redis2.Nil {
		logx.Errorf("通过文件夹id获取文件列表，redis获取zset失败，ERR: [%v]", redisErr)
	}

	var urls []redis2.Z
	for i, file := range files {
		var url string
		if file.Status == constant.StatusFileUploaded {
			if len(zs) == len(files) && redisErr == nil {
				url = zs[i].Member.(string)
			} else {
				url2, err := minioSvc.GenUrl(file.ObjectName, file.Name, true)
				if err != nil {
					logx.Errorf("通过文件夹id获取文件列表，[%d]获取url失败，ERR: [%v]", file.Id, err)
					continue
				} else {
					url = url2
					urls = append(urls, redis2.Z{Member: url, Score: float64(file.Created.Unix())})
				}
				if i == len(files)-1 {
					if err = rdb.ZAdd(l.ctx, key, urls...).Err(); err != nil {
						logx.Errorf("通过文件夹id获取文件列表，redis缓存url失败，ERR: [%v]", err)
					}
					if err = rdb.Expire(l.ctx, key, redis.DownloadExpire).Err(); err != nil {
						logx.Errorf("ListFileByType，设置过期时间失败，ERR: [%v]", err)
						return nil, err
					}
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
	return resp, nil
}
