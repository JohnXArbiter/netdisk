package user

import (
	"context"
	"errors"
	redis2 "github.com/redis/go-redis/v9"
	"lc/netdisk/common/constant"
	"lc/netdisk/common/redis"
	"lc/netdisk/model"
	"strconv"

	"lc/netdisk/internal/svc"
	"lc/netdisk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDetailLogic {
	return &GetDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDetailLogic) GetDetail(req *types.IdPathReq) (interface{}, error) {
	var (
		loginUserId = l.ctx.Value(constant.UserIdKey).(int64)
		userIdStr   = strconv.FormatInt(loginUserId, 10)
		engine      = l.svcCtx.Xorm
		rdb         = l.svcCtx.Redis
		minioSvc    = l.svcCtx.Minio.NewService()
		user        model.User
	)

	targetUserId := req.Id
	if req.Id == 0 {
		targetUserId = loginUserId
	}

	key := redis.UserInfoKey + userIdStr
	m, err := rdb.HGetAll(l.ctx, key).Result()
	if err != nil && err != redis2.Nil {
		logx.Errorf("获取用户info，redis获取失败，ERR: [%v]", err)
	} else if id, ok := m["id"]; err == redis2.Nil || !ok || id == "" {
		cols := "id, name, username, avatar, email, signature, status, used, capacity"
		if has, err := engine.Select(cols).ID(targetUserId).Get(&user); err != nil {
			logx.Errorf("更新用户info，数据库info获取失败，ERR: [%v]", err)
			return nil, err
		} else if !has {
			return nil, errors.New("用户信息有误")
		}
		url, _ := minioSvc.GenUrl(user.Avatar, false)
		m2 := map[string]interface{}{
			"id":        user.Id,
			"name":      user.Name,
			"username":  user.Username,
			"avatar":    url,
			"email":     user.Email,
			"signature": user.Signature,
			"capacity":  user.Capacity,
			"status":    user.Status,
			"used":      user.Used,
		}
		if err = rdb.HSet(l.ctx, key, m2).Err(); err != nil {
			logx.Errorf("更新用户info，info->redis失败，ERR: [%v]", err)
		}
		return m2, nil
	}
	return m, nil
}
