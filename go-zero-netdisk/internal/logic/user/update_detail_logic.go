package user

import (
	"context"
	"errors"
	redis2 "github.com/redis/go-redis/v9"
	"lc/netdisk/common/constant"
	"lc/netdisk/common/redis"
	"lc/netdisk/common/xorm"
	"lc/netdisk/internal/logic/mqs"
	"lc/netdisk/model"
	"strconv"

	"lc/netdisk/internal/svc"
	"lc/netdisk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDetailLogic {
	return &UpdateDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateDetailLogic) UpdateDetail(req *types.UpdateUserDetailReq) error {
	var (
		loginUserId = l.ctx.Value(constant.UserIdKey).(int64)
		userIdStr   = strconv.FormatInt(loginUserId, 10)
		engine      = l.svcCtx.Xorm
		rdb         = l.svcCtx.Redis
		user        model.User
		err         error
	)

	defer mqs.LogSend(l.ctx, err, "UpdateDetail", user)

	key := redis.UserInfoKey + userIdStr
	m, err := rdb.HGetAll(l.ctx, key).Result()
	if err != nil && err != redis2.Nil {
		logx.Errorf("更新用户info，redis获取失败，ERR: [%v]", err)
	}

	if id, ok := m["id"]; err == redis2.Nil || !ok || id == "" {
		if has, err2 := engine.ID(loginUserId).Get(&user); err != nil {
			err = errors.New("更新用户info，数据库info获取失败，ERR: " + err2.Error())
			return err
		} else if !has {
			return errors.New("用户信息有误")
		}
	} else {
		user.Id = loginUserId
		user.Name = req.Name
		user.Username = m["username"]
		user.Avatar = m["avatar"]
		user.Email = req.Email
		user.Signature = req.Signature
	}

	if req.Email != user.Email {
		code, err2 := rdb.Get(l.ctx, redis.EmailValidCode+req.Email).Result()
		if err2 != nil && err2 != redis2.Nil {
			err = errors.New("更新用户info，redis获取邮箱验证码失败，ERR: " + err2.Error())
			return err
		}
		if code != req.Code {
			return errors.New("验证码错误！请重新获取😭")
		}
	}

	user.Email = req.Email
	_, _ = engine.DoTransaction(func(session *xorm.Session) (interface{}, error) {
		if affected, err2 := engine.Cols("name", "email", "signature").
			ID(loginUserId).Update(user); err != nil {
			err = errors.New("更新用户info，更新数据库失败，ERR: " + err2.Error())
			return nil, err
		} else if affected != 1 {
			return nil, errors.New("出错了，请稍后")
		}

		if err = rdb.Del(l.ctx, key).Err(); err != nil {
			err = errors.New("更新用户info，后删redis记录失败，ERR: " + err.Error())
			return nil, err
		}
		return nil, nil
	})
	err = errors.New("测试")
	return nil
}
