package admin

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"lc/netdisk/common/constant"
	"lc/netdisk/common/variable"
	"lc/netdisk/internal/logic/mqs"
	"lc/netdisk/model"
	"strings"

	"lc/netdisk/internal/svc"
	"lc/netdisk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddAdminLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddAdminLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddAdminLogic {
	return &AddAdminLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddAdminLogic) AddAdmin(req *types.AddAdminReq) error {
	var (
		engine = l.svcCtx.Xorm
		admin  *model.Admin
		err    error
	)

	defer mqs.LogSend(l.ctx, err, "AddAdmin", req.Name, req.Username)

	admin, err = l.validate(req)
	if err != nil {
		return err
	}

	admin.Name = req.Name
	admin.Status = constant.StatusAdminNormal
	if _, err = engine.Insert(admin); err != nil {
		err = errors.New("" + err.Error())
		return err
	}

	return nil
}

// 校验账号密码
func (l *AddAdminLogic) validate(req *types.AddAdminReq) (*model.Admin, error) {
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

	admin := &model.Admin{
		Username: username,
		Password: string(password),
	}
	return admin, nil
}
