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

type DeleteFoldersTrulyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteFoldersTrulyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteFoldersTrulyLogic {
	return &DeleteFoldersTrulyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteFoldersTrulyLogic) DeleteFoldersTruly(req *types.IdsReq) error {
	var (
		userId  = l.ctx.Value(constant.UserIdKey).(int64)
		engine  = l.svcCtx.Xorm
		folders []*model.Folder
	)

	if err := engine.Cols("id").In("id", req.Ids).And("user_id = ?", userId).
		And("del_flag = ?", constant.StatusFileDeleted).Find(&folders); err != nil {
		return err
	}

	length := len(req.Ids)
	if length != len(folders) {
		return errors.New("发生错误！")
	}

	if affected, err := engine.In("id", req.Ids).
		Delete(&model.Folder{}); err != nil || affected != int64(length) {
		return errors.New("发生错误！" + err.Error())
	}

	// TODO: MQ
	return nil
}
