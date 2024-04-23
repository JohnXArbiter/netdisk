package logic

import (
	"context"
	"errors"
	"github.com/yitter/idgenerator-go/idgen"
	"golang.org/x/crypto/bcrypt"
	"lc/netdisk/common/constant"
	"lc/netdisk/common/redis"
	"lc/netdisk/common/variable"
	"lc/netdisk/internal/svc"
	"lc/netdisk/internal/types"
	"lc/netdisk/model"
	"math/rand"
	"strconv"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) error {
	var (
		engine = l.svcCtx.Xorm
		rdb    = l.svcCtx.Redis
	)

	key := redis.EmailValidCode + req.Email
	code, err := rdb.Get(l.ctx, key).Result()
	if err != nil {
		return errors.New("发送错误，" + err.Error())
	}

	if req.Code != code {
		return errors.New("验证码错误，请重新获取")
	}

	userInfo, err := l.validate(req)
	if err != nil {
		return err
	}

	userInfo.Email = req.Email
	userInfo.Id = idgen.NextId()
	userInfo.Name = "user_" + strconv.FormatInt(int64(rand.Int31()), 10)
	userInfo.Capacity = constant.DefaultCapacity
	//userInfo.Avatar = l.svcCtx.BgUrl + "/avatar" + strconv.Itoa(rand.Intn(3)) + ".jpg"
	if _, err = engine.Insert(userInfo); err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return errors.New("账号已经被抢走啦🫠")
		}
		return errors.New("注册失败，请重试")
	}

	return nil
}

// 校验账号密码
func (l *RegisterLogic) validate(req *types.RegisterReq) (*model.User, error) {
	var (
		username       = strings.TrimSpace(req.Username)
		password       = []byte(strings.TrimSpace(req.Password))
		passwordRepeat = []byte(strings.TrimSpace(req.PasswordRepeat))
	)

	if string(passwordRepeat) != string(password) {
		return nil, errors.New("两次密码不相等")
	}

	if !variable.Upattern.MatchString(username) {
		return nil, errors.New("账号格式错误，只能包含数字和字母，5-20位")
	}

	if !variable.Ppattern.MatchString(string(password)) {
		return nil, errors.New("密码格式错误，只能包含数字、字母和英文符号，6-20位")
	}

	password, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	userInfo := &model.User{
		Username: username,
		Password: string(password),
	}
	return userInfo, nil
}
