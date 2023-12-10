package file

import (
	"context"
	"errors"
	"lc/netdisk/common/constant"
	"lc/netdisk/common/xorm"
	"lc/netdisk/model"

	"lc/netdisk/internal/svc"
	"lc/netdisk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MoveLogic {
	return &MoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MoveLogic) Move(req *types.MoveReq) error {
	var (
		userId = l.ctx.Value(constant.UserIdKey).(int64)
		engine = l.svcCtx.Xorm
		folder model.Folder
	)

	if has, err := engine.ID(req.ParentFolderId).
		And("user_id = ?", userId).Get(&folder); err != nil {
		return errors.New("发生错误！")
	} else if !has {
		return errors.New("该目录不存在")
	}

	if folder.ParentId == 0 {
		return errors.New("该目录不可移动")
	}

	res, err := engine.DoTransaction(l.updateFolderId(req, userId))
	if err != nil {
		return errors.New("移动过程发生错误！")
	}
	if res == 0 {
		return errors.New("文件夹信息有误！")
	}

	return nil
}

func (l *MoveLogic) updateFolderId(param *types.MoveReq, userId int64) xorm.TxFn {
	return func(session *xorm.Session) (interface{}, error) {
		var folderId = param.ParentFolderId
		for _, id := range param.FileIds {
			if affected, err := session.ID(id).And("user_id = ?", userId).
				Update(&model.File{FolderId: folderId}); err != nil {
				return nil, err
			} else if affected != 1 {
				return 0, nil
			}
		}

		for _, id := range param.FolderIds {
			if affected, err := session.ID(id).And("user_id = ?", userId).
				Update(&model.Folder{ParentId: folderId}); err != nil {
				return nil, err
			} else if affected != 1 {
				return 0, nil
			}
		}

		return 1, nil
	}
}
