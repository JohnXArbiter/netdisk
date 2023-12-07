package file

import (
	"context"
	"errors"
	"lc/netdisk/common/constant"
	"lc/netdisk/common/xorm"
	"lc/netdisk/internal/svc"
	"lc/netdisk/internal/types"
	"lc/netdisk/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type MoveFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMoveFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MoveFileLogic {
	return &MoveFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MoveFileLogic) MoveFile(req *types.MoveFileReq) error {
	var (
		userId       = l.ctx.Value(constant.UserIdKey).(int64)
		engine       = l.svcCtx.Xorm
		fileFolder   model.FileFolder
		fileNetdisks []*model.FileNetdisk
	)

	if has, err := engine.ID(req.FolderId).Get(&fileFolder); err != nil {
		return err
	} else if !has {
		return errors.New("未找到该文件夹！")
	}

	if err := engine.In("id", req.FileIds).And("user_id = ?", userId).
		Find(&fileNetdisks); err != nil {
		return err
	}

	if len(fileNetdisks) != len(req.FileIds) {
		return errors.New("出错了！")
	}

	for _, fileNetdisk := range fileNetdisks {
		fileNetdisk.FolderId = req.FolderId
		engine.Update(fileNetdisks)

	}

	_, err := engine.DoTransaction(l.createFsAndNetdiskRecord(fileNetdisks))

	return err
}

func (l *MoveFileLogic) createFsAndNetdiskRecord(fileNetdisks []*model.FileNetdisk) xorm.TxFn {
	return func(session *xorm.Session) (interface{}, error) {
		for _, fileNetdisk := range fileNetdisks {
			_, err := session.Cols("folder_id").Update(fileNetdisk)
			if err != nil {
				return nil, err
			}
		}

		return nil, nil
	}
}
