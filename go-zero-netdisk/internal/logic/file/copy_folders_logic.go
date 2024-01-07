package file

import (
	"context"
	"errors"
	"github.com/yitter/idgenerator-go/idgen"
	"lc/netdisk/common/constant"
	"lc/netdisk/model"
	"time"

	"lc/netdisk/internal/svc"
	"lc/netdisk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CopyFoldersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCopyFoldersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CopyFoldersLogic {
	return &CopyFoldersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CopyFoldersLogic) CopyFolders(req *types.CopyFoldersReq) error {
	var (
		userId         = l.ctx.Value(constant.UserIdKey).(int64)
		engine         = l.svcCtx.Xorm
		parentFolderId = req.ParentFolderId
		folders        []*model.Folder
	)

	has, err := engine.ID(parentFolderId).And("user_id = ?", userId).Get(&model.Folder{})
	if err != nil {
		return errors.New("发生错误！" + err.Error())
	} else if !has {
		return errors.New("该目录不存在")
	}

	if err = engine.In("id", req.FolderIds).Find(&folders); err != nil {
		return errors.New("发生错误！" + err.Error())
	}

	now := time.Now()
	for _, folder := range folders {
		folder.Id = idgen.NextId()
		folder.Name = folder.Name + "_复制"
		folder.Created = now
		folder.Updated = now
	}

	affected, err := engine.Insert(folders)
	if err != nil || affected != int64(len(folders)) {
		return errors.New("发生错误！" + err.Error())
	}
	return nil
}
