package logic

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"lc/netdisk/common/redis"
	"lc/netdisk/common/utils"
	"lc/netdisk/model"
	"strconv"
	"strings"
	"time"

	"lc/netdisk/internal/svc"
	"lc/netdisk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (*types.LoginResp, error) {
	var (
		username    = strings.TrimSpace(req.Username)
		password    = strings.TrimSpace(req.Password)
		engine      = l.svcCtx.Xorm
		redisClient = l.svcCtx.Redis
	)

	user := &model.User{Username: username}
	if has, err := engine.Cols("id", "username", "password",
		"name").Get(user); err != nil || !has {
		return nil, errors.New("帐号或密码错误！")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password),
		[]byte(password)); err != nil {
		return nil, errors.New("帐号或密码错误！")
	}

	token, err := utils.GenToken(user)
	if err != nil {
		return nil, errors.New("出错啦，请重试！")
	}

	key := redis.LoggedUser + strconv.FormatInt(user.Id, 10)
	if err = redisClient.Set(l.ctx, key, token, 7*24*time.Hour).Err(); err != nil {
		logx.Errorf("[REDIS ERROR] Login 保存用户token失败，userid：%v %v\n", user.Id, err)
		l.svcCtx.Redis.Set(l.ctx, key, token, 7*24*time.Hour) // 重试
	}

	var userInfo types.UserInfo
	userInfo.UserId = user.Id
	userInfo.Name = user.Name
	userInfo.Avatar = user.Avatar
	userInfo.Email = user.Email
	userInfo.Signature = user.Signature
	userInfo.Status = user.Status

	resp := &types.LoginResp{}
	resp.UserInfo = userInfo
	resp.Token = token
	return resp, nil
}
