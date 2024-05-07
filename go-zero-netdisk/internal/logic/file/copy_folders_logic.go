package file

import (
	"context"
	"errors"
	"github.com/yitter/idgenerator-go/idgen"
	"lc/netdisk/common/constant"
	"lc/netdisk/internal/logic/mqs"
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
		err            error
	)

	defer mqs.LogSend(l.ctx, err, "CopyFolders", req.ParentFolderId)

	if parentFolderId != 0 {
		folder := &model.Folder{}
		if _, err = engine.ID(parentFolderId).
			And("user_id = ?", userId).
			Get(folder); err != nil {
			err = errors.New("发生错误！." + err.Error())
			return err
		}
		if folder.Id == 0 {
			return errors.New("该目录不存在")
		}
	}

	if err = engine.In("id", req.FolderIds).
		Find(&folders); err != nil {
		err = errors.New("发生错误！.." + err.Error())
		return err
	}

	for _, folder := range folders {
		folder.Id = idgen.NextId()
		folder.Name = folder.Name + "_" + time.Now().Format(constant.TimeFormat2) + "复制"
		folder.ParentId = req.ParentFolderId
	}

	if _, err = engine.Insert(&folders); err != nil {
		return errors.New("发生错误！..." + err.Error())
	}
	return nil
}
