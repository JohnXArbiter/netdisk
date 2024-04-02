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

type MoveFilesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMoveFilesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MoveFilesLogic {
	return &MoveFilesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MoveFilesLogic) MoveFiles(req *types.MoveFilesReq) error {
	var (
		userId   = l.ctx.Value(constant.UserIdKey).(int64)
		engine   = l.svcCtx.Xorm
		folderId = req.FolderId
		err      error
	)

	defer mqs.LogSend(l.ctx, err, "MoveFiles", req.FileIds, req.FolderId)

	if folderId != 0 {
		has, err2 := engine.ID(folderId).And("user_id = ?", userId).Get(&model.Folder{})
		if err2 != nil {
			err = errors.New("发生错误！" + err2.Error())
			return err
		} else if !has {
			err = errors.New("该目录不存在")
			return err
		}
	}

	file := &model.File{FolderId: folderId}
	if affected, err2 := engine.In("id", req.FileIds).
		Update(file); err != nil {
		err = errors.New("移动出错！" + err2.Error())
		return err
	} else if affected != int64(len(req.FileIds)) {
		err = errors.New("移动出错！")
		return err
	}
	return nil
}
