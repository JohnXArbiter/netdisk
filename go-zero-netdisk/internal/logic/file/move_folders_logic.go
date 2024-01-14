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
	)

	if parentFolderId != 0 {
		has, err := engine.ID(parentFolderId).And("user_id = ?", userId).Get(folder)
		if err != nil {
			return errors.New("发生错误！")
		} else if !has {
			return errors.New("该目录不存在")
		}
	}

	folder = &model.Folder{ParentId: parentFolderId}
	affected, err := engine.In("id", req.FolderIds).Update(folder)
	if err != nil || affected != int64(len(req.FolderIds)) {
		return errors.New("" + err.Error())
	}

	return nil
}
