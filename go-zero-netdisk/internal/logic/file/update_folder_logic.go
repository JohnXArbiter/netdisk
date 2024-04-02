package file

import (
	"context"
	"errors"
	"lc/netdisk/common/constant"
	"lc/netdisk/internal/logic/mqs"
	"lc/netdisk/model"

	"lc/netdisk/internal/svc"
	"lc/netdisk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateFolderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateFolderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateFolderLogic {
	return &UpdateFolderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateFolderLogic) UpdateFolder(req *types.UpdateNameReq) error {
	var (
		userId = l.ctx.Value(constant.UserIdKey).(int64)
		engine = l.svcCtx.Xorm
		err    error
	)

	defer mqs.LogSend(l.ctx, err, "UpdateFolder", req.Id, req.Name)

	if affected, err := engine.ID(req.Id).And("user_id = ?", userId).
		Update(&model.Folder{Name: req.Name}); err != nil {
		return err
	} else if affected != 1 {
		return errors.New("文件信息有误！")
	}

	return nil
}
