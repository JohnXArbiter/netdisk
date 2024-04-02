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

type MoveFoldersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMoveFoldersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MoveFoldersLogic {
	return &MoveFoldersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MoveFoldersLogic) MoveFolders(req *types.MoveFoldersReq) error {
	var (
		userId         = l.ctx.Value(constant.UserIdKey).(int64)
		engine         = l.svcCtx.Xorm
		parentFolderId = req.ParentFolderId
		folder         = &model.Folder{}
		err            error
	)

	defer mqs.LogSend(l.ctx, err, "MoveFolders", req.FolderIds, req.ParentFolderId)

	if parentFolderId != 0 {
		if has, err2 := engine.ID(parentFolderId).
			And("user_id = ?", userId).
			Get(folder); err2 != nil {
			err = errors.New("发生错误！")
			return err
		} else if !has {
			err = errors.New("该目录不存在")
			return err
		}
	}

	folder = &model.Folder{ParentId: parentFolderId}
	if _, err2 := engine.In("id", req.FolderIds).
		Update(folder); err2 != nil {
		err = errors.New("移动文件夹发送错误，" + err2.Error())
		return err
	}

	return nil
}
