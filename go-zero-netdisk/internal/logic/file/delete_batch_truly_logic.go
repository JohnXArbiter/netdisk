package file

import (
	"context"
	"errors"
	"lc/netdisk/common/constant"
	"lc/netdisk/common/xorm"
	"lc/netdisk/model"

	"lc/netdisk/internal/svc"
	"lc/netdisk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteBatchTrulyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteBatchTrulyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteBatchTrulyLogic {
	return &DeleteBatchTrulyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteBatchTrulyLogic) DeleteBatchTruly(req *types.DeleteBatchTrulyReq) (resp *types.DeleteBatchTrulyResp, err error) {

	l.ctx = context.WithValue(l.ctx, constant.CtxFolderIdsKey, req.FolderIds)
	l.ctx = context.WithValue(l.ctx, constant.CtxFileIdsKey, req.FileIds)
	l.svcCtx.Xorm.DoTransactions(nil, l.deleteFolders, l.deleteFiles)

	return
}

func (l *DeleteBatchTrulyLogic) deleteFolders(session *xorm.Session) (interface{}, error) {
	var (
		ctx       = l.ctx
		userId    = ctx.Value(constant.UserIdKey).(int64)
		folderIds = ctx.Value(constant.CtxFolderIdsKey).([]int64)
		folders   []*model.Folder
	)

	if err := session.Cols("id").In("id", folderIds).And("user_id = ?", userId).
		And("del_flag = ?", constant.StatusFolderDeleted).Find(&folders); err != nil {
		return nil, err
	}

	length := len(folderIds)
	if length != len(folders) {
		return nil, errors.New("发生错误！")
	}

	if affected, err := session.Cols("id").In("id", folderIds).
		Delete(&model.File{}); err != nil || affected != int64(length) {
		return nil, err
	}

	return nil, nil
}

func (l *DeleteBatchTrulyLogic) deleteFiles(session *xorm.Session) (interface{}, error) {
	var (
		ctx     = l.ctx
		userId  = ctx.Value(constant.UserIdKey).(int64)
		fileIds = ctx.Value(constant.CtxFileIdsKey).([]int64)
		files   []*model.File
	)

	if err := session.Cols("id").In("id", fileIds).And("user_id = ?", userId).
		And("del_flag = ?", constant.StatusFileDeleted).Find(&files); err != nil {
		return nil, err
	}

	length := len(fileIds)
	if length != len(files) {
		return nil, errors.New("发生错误！")
	}

	if affected, err := session.Cols("id").In("id", fileIds).
		Delete(&model.File{}); err != nil || affected != int64(length) {
		return nil, err
	}

	return nil, nil
}
