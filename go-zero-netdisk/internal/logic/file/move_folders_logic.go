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
		if _, err = engine.ID(parentFolderId).
			And("user_id = ?", userId).
			Get(folder); err != nil {
			err = errors.New("发生错误！")
			return err
		}
		if folder.Id <= 0 {
			err = errors.New("该目录不存在")
			return err
		}
	}

	folder = &model.Folder{ParentId: parentFolderId}
	if _, err = engine.Cols("parent_id").
		In("id", req.FolderIds).
		Update(folder); err != nil {
		err = errors.New("移动文件夹发送错误，" + err.Error())
		return err
	}

	return nil
}
