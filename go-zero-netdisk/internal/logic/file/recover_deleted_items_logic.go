package file

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"lc/netdisk/common/constant"
	"lc/netdisk/common/xorm"
	"lc/netdisk/internal/svc"
	"lc/netdisk/internal/types"
	"lc/netdisk/model"
)

type RecoverDeletedItemsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRecoverDeletedItemsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RecoverDeletedItemsLogic {
	return &RecoverDeletedItemsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RecoverDeletedItemsLogic) RecoverDeletedItems(req *types.RecoverReq) error {

	l.ctx = context.WithValue(l.ctx, constant.CtxFolderIdsKey, req.FolderIds)
	l.ctx = context.WithValue(l.ctx, constant.CtxFileIdsKey, req.FileIds)
	_, err := l.svcCtx.Xorm.DoTransactions(nil, l.recoverFolderStatus, l.recoverFilesStatus)
	return err
}

func (l *RecoverDeletedItemsLogic) recoverFolderStatus(session *xorm.Session) (interface{}, error) {
	var (
		ctx       = l.ctx
		userId    = ctx.Value(constant.UserIdKey).(int64)
		folderIds = ctx.Value(constant.CtxFolderIdsKey).([]int64)
		folders   []*model.Folder
	)

	length := len(folderIds)
	if err := session.In("id", folderIds).And("user_id = ?", userId).
		And("del_flag", constant.StatusFolderDeleted).
		Find(&folders); err != nil || length != len(folders) {
		return nil, errors.New("还原过程出错！")
	}

	bean := &model.Folder{DelFlag: constant.StatusFolderUndeleted}
	if affected, err := session.In("id", folderIds).And("user_id = ?", userId).
		Update(bean); err != nil || affected != int64(length) {
		return nil, errors.New("还原过程出错！")
	}
	return nil, nil
}

func (l *RecoverDeletedItemsLogic) recoverFilesStatus(session *xorm.Session) (interface{}, error) {
	var (
		ctx     = l.ctx
		userId  = ctx.Value(constant.UserIdKey).(int64)
		fileIds = ctx.Value(constant.CtxFileIdsKey).([]int64)
		files   []*model.File
	)

	length := len(fileIds)
	if err := session.In("id", fileIds).And("user_id = ?", userId).
		And("del_flag", constant.StatusFileDeleted).
		Find(&files); err != nil || length != len(files) {
		return nil, errors.New("还原过程出错！")
	}

	bean := &model.File{DelFlag: constant.StatusFileUndeleted}
	if affected, err := session.In("id", fileIds).And("user_id = ?", userId).
		Update(bean); err != nil || affected != int64(len(fileIds)) {
		return nil, errors.New("还原过程出错！")
	}
	return nil, nil
}
