package file

import (
	"context"
	"errors"
	"lc/netdisk/common/constant"
	"lc/netdisk/common/xorm"
	"lc/netdisk/model"
	"time"

	"lc/netdisk/internal/svc"
	"lc/netdisk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteBatchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteBatchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteBatchLogic {
	return &DeleteBatchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteBatchLogic) DeleteBatch(req *types.DeleteBatchReq) (resp *types.DeleteBatchResp, err error) {
	var (
		engine = l.svcCtx.Xorm
	)

	l.ctx = context.WithValue(l.ctx, constant.CtxFolderIdsKey, req.FolderIds)
	l.ctx = context.WithValue(l.ctx, constant.CtxFileIdsKey, req.FileIds)
	_, err = engine.DoTransactions(nil, l.updateFoldersStatus, l.updateFileNetdisksStatus)
	return
}

func (l *DeleteBatchLogic) updateFoldersStatus(session *xorm.Session) (interface{}, error) {
	var (
		ctx       = l.ctx
		userId    = ctx.Value(constant.UserIdKey).(int64)
		folderIds = ctx.Value(constant.CtxFolderIdsKey).([]int64)
	)

	bean := &model.Folder{
		DelFlag: constant.StatusFileDeleted,
		DelTime: time.Now().Local().Unix(),
	}
	if affected, err := session.In("id", folderIds).And("user_id = ?", userId).
		Update(bean); err != nil || affected != int64(len(folderIds)) {
		return nil, errors.New("删除过程出错！")
	}
	return nil, nil
}

func (l *DeleteBatchLogic) updateFileNetdisksStatus(session *xorm.Session) (interface{}, error) {
	var (
		ctx     = l.ctx
		userId  = ctx.Value(constant.UserIdKey).(int64)
		fileIds = ctx.Value(constant.CtxFileIdsKey).([]int64)
	)

	bean := &model.File{
		DelFlag: constant.StatusFileDeleted,
		DelTime: time.Now().Local().Unix(),
	}
	if affected, err := session.In("id", fileIds).And("user_id = ?", userId).
		Update(bean); err != nil || affected != int64(len(fileIds)) {
		return nil, errors.New("删除过程出错！")
	}
	return nil, nil
}
