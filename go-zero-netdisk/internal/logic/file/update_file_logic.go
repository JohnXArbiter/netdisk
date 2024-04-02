package file

import (
	"context"
	"lc/netdisk/common/constant"
	"lc/netdisk/common/variable"
	"lc/netdisk/internal/logic/mqs"
	"lc/netdisk/model"
	"strings"

	"lc/netdisk/internal/svc"
	"lc/netdisk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateFileLogic {
	return &UpdateFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateFileLogic) UpdateFile(req *types.UpdateNameReq) error {
	var (
		userId = l.ctx.Value(constant.UserIdKey).(int64)
		engine = l.svcCtx.Xorm
		err    error
	)

	defer mqs.LogSend(l.ctx, err, "UpdateFile", req.Id, req.Name)

	ext := req.Name[strings.LastIndex(req.Name, "."):]
	fType := variable.GetTypeByBruteForce(ext)
	file := &model.File{Name: req.Name, Type: fType}
	if _, err = engine.ID(req.Id).
		And("user_id = ?", userId).
		Update(file); err != nil {
		return err
	}

	return nil
}
