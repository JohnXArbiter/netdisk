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

type DeleteFilesTrulyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteFilesTrulyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteFilesTrulyLogic {
	return &DeleteFilesTrulyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteFilesTrulyLogic) DeleteFilesTruly(req *types.IdsReq) error {
	var (
		userId = l.ctx.Value(constant.UserIdKey).(int64)
		engine = l.svcCtx.Xorm
		files  []*model.File
		err    error
	)

	defer mqs.LogSend(l.ctx, err, "DeleteFilesTruly", req.Ids)

	if err = engine.In("id", req.Ids).
		And("user_id = ?", userId).
		And("del_flag = ?", constant.StatusFileDeleted).
		Find(&files); err != nil {
		return err
	}

	length := len(req.Ids)
	if length != len(files) {
		return errors.New("发生错误！")
	}

	if affected, err2 := engine.In("id", req.Ids).
		Delete(&model.File{}); err2 != nil {
		err = errors.New("发生错误，" + err2.Error())
		return err
	} else if affected != int64(length) {
		err = errors.New("发生错误！")
		return err
	}

	// TODO: MQ
	return nil
}
