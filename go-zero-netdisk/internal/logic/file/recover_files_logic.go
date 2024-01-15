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
		ctx    = l.ctx
		userId = ctx.Value(constant.UserIdKey).(int64)
		engine = l.svcCtx.Xorm
		files  []*model.File
	)

	if err := engine.In("id", req.FileIds).Find(&files); err != nil {
		return err
	}

	if len(req.FileIds) != len(files) {
		return errors.New("选择信息有误")
	}

	var (
		folderIds   []int64
		folderIdMap = make(map[int64]struct{})
	)
	for _, file := range files {
		if file.UserId != userId {
			return errors.New("不是你的文件")
		}
		if file.DelFlag != constant.StatusFileDeleted {
			return errors.New("文件信息错误")
		}
		if _, ok := folderIdMap[file.FolderId]; !ok {
			folderIdMap[file.FolderId] = struct{}{}
			folderIds = append(folderIds, file.FolderId)
		}
	}

	_, err := engine.DoTransaction(func(session *xorm.Session) (interface{}, error) {
		if affected, err := session.Cols("del_flag").In("id", req.FileIds).
			Update(&model.File{DelFlag: constant.StatusFileUndeleted}); err != nil {
			return nil, err
		} else if affected != int64(len(files)) {
			return nil, errors.New("ASD")
		}

		// 所有涉及的文件夹应都不为被删除状态
		if affected, err := session.Cols("del_flag").In("id", folderIds).
			Update(&model.Folder{DelFlag: constant.StatusFolderUndeleted}); err != nil {
			return nil, err
		} else if affected != int64(len(folderIds)) {
			return nil, errors.New("ASD")
		}
		return nil, nil
	})
	return err
}
