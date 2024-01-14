package file

import (
	"context"
	"lc/netdisk/common/constant"
	"lc/netdisk/model"

	"lc/netdisk/internal/svc"
	"lc/netdisk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RecoverFilesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRecoverFilesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RecoverFilesLogic {
	return &RecoverFilesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RecoverFilesLogic) RecoverFiles(req *types.FileIdsStruct) error {
	var (
		ctx       = l.ctx
		userId    = ctx.Value(constant.UserIdKey).(int64)
		engine    = l.svcCtx.Xorm
		folderIds = ctx.Value(constant.CtxFolderIdsKey).([]int64)
		folders   []*model.Folder
	)

	if err := engine.In("id", req.FileIds).Find(&folders); err != nil {

	}

	return nil
}
