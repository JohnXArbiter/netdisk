package file

import (
	"context"
	"errors"
	"lc/netdisk/common/constant"
	"lc/netdisk/internal/svc"
	"lc/netdisk/internal/types"
	"lc/netdisk/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateFolderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateFolderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateFolderLogic {
	return &CreateFolderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateFolderLogic) CreateFolder(req *types.CreateFolderReq) error {
	var (
		userId   = l.ctx.Value(constant.UserIdKey).(int64)
		engine   = l.svcCtx.Xorm
		parentId = req.ParentFolderId
	)

	if exist, err := engine.ID(parentId).And("user_id = ?", userId).
		Exist(&model.Folder{}); err != nil || !exist {
		return errors.New("发生错误！")
	}

	newFolder := &model.Folder{
		ParentId: parentId,
		Name:     req.Name,
		UserId:   userId,
		DelFlag:  constant.StatusFolderUndeleted,
	}
	if _, err := engine.Insert(newFolder); err != nil {
		return errors.New("创建失败了！")
	}

	return nil
}
