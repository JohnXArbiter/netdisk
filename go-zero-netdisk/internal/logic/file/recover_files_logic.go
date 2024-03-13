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
	"strconv"
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

func (l *RecoverFilesLogic) RecoverFiles(req *types.RecoverFilesReq) error {
	var (
		ctx    = l.ctx
		userId = ctx.Value(constant.UserIdKey).(int64)
		engine = l.svcCtx.Xorm
	)

	fileIds, folderIds := req.FileIds, req.FolderIds
	_, err := engine.DoTransaction(func(session *xorm.Session) (interface{}, error) {
		// 1.恢复文件delFlag字段
		fileBean := &model.File{
			DelFlag: constant.StatusFileUndeleted,
			DelTime: 0,
		}
		if affected, err := session.In("id", fileIds).
			And("user_id = ?", userId).
			Update(fileBean); err != nil {
			logx.Errorf("恢复文件信息出错，err: %v", err)
			return nil, err
		} else if affected != int64(len(fileIds)) {
			return nil, errors.New("恢复文件信息出错！")
		}

		for len(folderIds) > 0 {
			// 2.恢复文件夹delFlag字段
			folderBean := &model.Folder{
				DelFlag: constant.StatusFolderUndeleted,
				DelTime: 0,
			}
			if _, err := session.In("id", folderIds).
				And("del_flag = ?", constant.StatusFolderDeleted).
				Update(folderBean); err != nil {
				logx.Errorf("恢复文件夹信息出错，err: %v", err)
				return nil, err
			}

			// 3.查询父文件夹
			var ids []int64
			subQuery := "select distinct(parent_id) from folder where id in ("
			for i, folderId := range folderIds {
				subQuery += strconv.FormatInt(folderId, 10)
				if i != len(folderIds)-1 {
					subQuery += ","
				}
			}
			if err := session.Select("id").Table(&model.Folder{}).
				Where("id in ( "+subQuery+") )").
				And("del_flag = ?", constant.StatusFolderDeleted).
				Find(&ids); err != nil {
				logx.Errorf("查询父文件夹出错，err: %v", err)
				return nil, err
			}
			folderIds = ids
		}
		return nil, nil
	})

	return err
}
