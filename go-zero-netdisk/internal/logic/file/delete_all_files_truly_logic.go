package file

import (
	"context"
	"errors"
	"lc/netdisk/common/constant"
	"lc/netdisk/internal/logic/mqs"
	"lc/netdisk/model"

	"github.com/zeromicro/go-zero/core/logx"
	"lc/netdisk/internal/svc"
)

type DeleteAllFilesTrulyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteAllFilesTrulyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteAllFilesTrulyLogic {
	return &DeleteAllFilesTrulyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteAllFilesTrulyLogic) DeleteAllFilesTruly() error {
	var (
		userId = l.ctx.Value(constant.UserIdKey).(int64)
		engine = l.svcCtx.Xorm
		files  []*model.File
		err    error
	)

	defer mqs.LogSend(l.ctx, err, "DeleteAllFilesTruly")

	if err = engine.Where("user_id = ?", userId).
		And("del_flag = ?", constant.StatusFileDeleted).
		Find(&files); err != nil {
		return errors.New("出错啦！，" + err.Error())
	}

	if affected, err := engine.Where("user_id = ?", userId).
		And("del_flag = ?", constant.StatusFileDeleted).
		Delete(&model.File{}); err != nil {
		return errors.New("删除过程出错啦，" + err.Error())
	} else if affected != int64(len(files)) {
		return errors.New("删除过程出错啦！")
	}

	// TODO: MQ
	go l.fs(files)

	return nil
}

func (l *DeleteAllFilesTrulyLogic) fs(files []*model.File) {

}
