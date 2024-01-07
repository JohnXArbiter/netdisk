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

type CopyFilesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCopyFilesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CopyFilesLogic {
	return &CopyFilesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CopyFilesLogic) CopyFiles(req *types.CopyFilesReq) error {
	var (
		userId         = l.ctx.Value(constant.UserIdKey).(int64)
		engine         = l.svcCtx.Xorm
		parentFolderId = req.ParentFolderId
		files          []*model.File
	)

	has, err := engine.ID(parentFolderId).And("user_id = ?", userId).Get(&model.Folder{})
	if err != nil {
		return errors.New("发生错误！" + err.Error())
	} else if !has {
		return errors.New("该目录不存在")
	}

	if err = engine.In("id", req.FileIds).Find(&files); err != nil {
		return errors.New("发生错误！" + err.Error())
	}

	now := time.Now()
	for _, file := range files {
		file.Id = idgen.NextId()
		file.Name = file.Name + "_复制"
		file.Created = now
		file.Updated = now
	}

	affected, err := engine.Insert(files)
	if err != nil || affected != int64(len(files)) {
		return errors.New("发生错误！" + err.Error())
	}
	return nil
}
