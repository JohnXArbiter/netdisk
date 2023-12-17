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

	userInfo := &model.User{Username: username}
	if has, err := engine.Cols("id", "username", "password",
		"name").Get(userInfo); err != nil || !has {
		return nil, errors.New("帐号或密码错误！")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userInfo.Password),
		[]byte(password)); err != nil {
		return nil, errors.New("帐号或密码错误！")
	}

	token, err := utils.GenToken(userInfo)
	if err != nil {
		return nil, errors.New("出错啦，请重试！")
	}

	key := redis.LoggedUser + strconv.FormatInt(userInfo.Id, 10)
	if err = redisClient.Set(l.ctx, key, token, 7*24*time.Hour).Err(); err != nil {
		logx.Errorf("[REDIS ERROR] Login 保存用户token失败，userid：%v %v\n", userInfo.Id, err)
		l.svcCtx.Redis.Set(l.ctx, key, token, 7*24*time.Hour) // 重试
	}

	resp := &types.LoginResp{}
	resp.UserId = userInfo.Id
	resp.Token = token
	return resp, nil
}
