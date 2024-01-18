package file

import (
	"context"
	"errors"
	"lc/netdisk/common/constant"
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
	)

	if err := engine.Cols("id").In("id", req).And("user_id = ?", userId).
		And("del_flag = ?", constant.StatusFileDeleted).Find(&files); err != nil {
		return err
	}

	length := len(req.Ids)
	if length != len(files) {
		return errors.New("发生错误！")
	}

	if affected, err := engine.In("id", req.Ids).
		Delete(&model.File{}); err != nil || affected != int64(length) {
		return errors.New("发生错误！" + err.Error())
	}

	// TODO: MQ
	return nil
}
