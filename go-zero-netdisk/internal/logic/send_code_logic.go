package logic

import (
	"context"
	"errors"
	"lc/netdisk/common/redis"
	"lc/netdisk/internal/svc"
	"lc/netdisk/internal/types"
	"lc/netdisk/model"
	"math/rand"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendCodeLogic {
	return &SendCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendCodeLogic) SendCode(req *types.SendCodeReq) error {
	var (
		engine = l.svcCtx.Xorm
		rdb    = l.svcCtx.Redis
	)

	key := redis.EmailValidCode + req.Email
	//if res := rdb.Get(l.ctx, key).Val(); res != "" {
	//	return errors.New("验证码已发送，请勿重复发送")
	//}

	if cnt, err := engine.
		Where("email = ?", req.Email).
		Count(&model.User{}); err != nil {
		logx.Errorf("发送验证码，检查邮箱是否重复出错，ERR: [%v]", err)
		return err
	} else if cnt > 0 {
		return errors.New("该邮箱已被抢注啦😨")
	}

	logx.Infof(req.Email)
	randNum := rand.Intn(100000) + 1000000
	if err := l.svcCtx.Email.SendCode(req.Email, strconv.Itoa(randNum)); err != nil {
		logx.Errorf("发送验证码，发送邮件出错，ERR: [%v]", err)
		return err
	}

	rdb.Set(l.ctx, key, randNum, redis.RegisterCodeExpire)

	return nil
}
