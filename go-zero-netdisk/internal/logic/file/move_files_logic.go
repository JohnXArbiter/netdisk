package file

import (
	"context"
	"errors"
	"fmt"
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
		userId   = l.ctx.Value(constant.UserIdKey).(int64)
		engine   = l.svcCtx.Xorm
		folderId = req.FolderId
	)

	if folderId != 0 {
		has, err := engine.ID(folderId).And("user_id = ?", userId).Get(&model.Folder{})
		if err != nil {
			return errors.New("发生错误！" + err.Error())
		} else if !has {
			return errors.New("该目录不存在")
		}
	}

	file := &model.File{FolderId: folderId}
	affected, err := engine.In("id", req.FileIds).Update(file)
	fmt.Println(affected, err)
	if err != nil {
		return errors.New("移动出错！" + err.Error())
	} else if affected != int64(len(req.FileIds)) {
		return errors.New("移动出错！")
	}
	return nil
}
