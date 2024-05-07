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
		userId   = l.ctx.Value(constant.UserIdKey).(int64)
		engine   = l.svcCtx.Xorm
		folderId = req.FolderId
		files    []*model.File
		err      error
	)

	defer mqs.LogSend(l.ctx, err, "CopyFiles", req.FileIds)

	if folderId != 0 {
		folder := &model.Folder{}
		if _, err = engine.ID(folderId).
			And("user_id = ?", userId).
			Get(folder); err != nil {
			err = errors.New("发生错误！." + err.Error())
			return err
		}
		if folder.Id == 0 {
			return errors.New("该目录不存在")
		}
	}

	if err = engine.In("id", req.FileIds).
		Find(&files); err != nil {
		err = errors.New("发生错误！.." + err.Error())
		return err
	}

	for _, file := range files {
		file.Id = idgen.NextId()
		file.Name = file.Name + "_" + time.Now().Format(constant.TimeFormat2) + "复制"
		file.FolderId = req.FolderId
	}

	if _, err = engine.Insert(&files); err != nil {
		return errors.New("发生错误！..." + err.Error())
	}
	return nil
}
