package logic

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"lc/netdisk/common/redis"
	"lc/netdisk/common/utils"
	"lc/netdisk/internal/logic/mqs"
	"lc/netdisk/model"
	"strconv"
	"strings"
	"time"

	"lc/netdisk/internal/svc"
	"lc/netdisk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminLoginLogic {
	return &AdminLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminLoginLogic) AdminLogin(req *types.LoginReq) (*types.LoginResp, error) {
	var (
		username = strings.TrimSpace(req.Username)
		password = strings.TrimSpace(req.Password)
		engine   = l.svcCtx.Xorm
		rdb      = l.svcCtx.Redis
		err      error
	)

	defer mqs.LogSend(l.ctx, err, "AdminLogin", username)

	admin := &model.Admin{Username: username}
	if _, err2 := engine.Cols("id", "username", "password",
		"name").Get(admin); err2 != nil {
		err = errors.New("帐号或密码错误，" + err2.Error())
		return nil, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(admin.Password),
		[]byte(password)); err != nil {
		return nil, errors.New("帐号或密码错误！")
	}

	token, err := utils.GenToken(admin.Id, admin.Name)
	if err != nil {
		return nil, errors.New("出错啦，请重试！")
	}

	key := redis.UserLogin + strconv.FormatInt(admin.Id, 10)
	if err = rdb.Set(l.ctx, key, token, 7*24*time.Hour).Err(); err != nil {
		logx.Errorf("AdminLogin，保存用户token失败，userid：%v %v\n", admin.Id, err)
		l.svcCtx.Redis.Set(l.ctx, key, token, 7*24*time.Hour) // 重试
	}

	var userInfo types.UserInfo
	userInfo.UserId = admin.Id
	userInfo.Name = admin.Name
	userInfo.Status = admin.Status

	resp := &types.LoginResp{}
	resp.UserInfo = userInfo
	resp.Token = token
	return resp, nil
}
