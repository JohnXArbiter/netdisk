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
		userId         = l.ctx.Value(constant.UserIdKey).(int64)
		engine         = l.svcCtx.Xorm
		parentFolderId = req.ParentFolderId
	)

	has, err := engine.ID(parentFolderId).And("user_id = ?", userId).Get(&model.Folder{})
	if err != nil {
		return errors.New("发生错误！" + err.Error())
	} else if !has {
		return errors.New("该目录不存在")
	}

	file := &model.File{FolderId: parentFolderId}
	affected, err := engine.In("id", req.FileIds).Update(file)
	if err != nil || affected != int64(len(req.FileIds)) {
		return errors.New("移动出错！" + err.Error())
	}
	return nil
}
