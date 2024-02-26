package file

import (
	"context"
	"errors"
	"lc/netdisk/common/constant"
	"lc/netdisk/model"
	"time"
	"xorm.io/xorm"

	"lc/netdisk/internal/svc"
	"lc/netdisk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteFoldersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteFoldersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteFoldersLogic {
	return &DeleteFoldersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteFoldersLogic) DeleteFolders(req *types.IdsReq) error {
	var (
		userId = l.ctx.Value(constant.UserIdKey).(int64)
		engine = l.svcCtx.Xorm
	)

	folderIds := req.Ids
	_, err := engine.Transaction(func(session *xorm.Session) (interface{}, error) {
		var (
			folders []*model.Folder
			err     error
			now     = time.Now().Local().Unix()
		)

		for len(folderIds) > 0 {
			// 1.删除当前文件夹下的文件
			fileBean := &model.File{
				DelFlag: constant.StatusFileDeleted,
				DelTime: now,
			}
			if _, err = session.In("folder_id", folderIds).
				And("user_id = ?", userId).
				And("del_flag = ?", constant.StatusFileUndeleted).
				Update(fileBean); err != nil {
				return nil, err
			}

			// 2.删除当前选中的文件夹
			folderBean := &model.Folder{
				DelFlag: constant.StatusFolderDeleted,
				DelTime: now,
			}
			if _, err = session.In("id", folderIds).
				And("user_id = ?", userId).
				And("del_flag = ?", constant.StatusFolderUndeleted).
				Update(folderBean); err != nil {
				return nil, err
			}

			// 3.搜索下一层文件夹
			if err = session.Select("id").In("parent_id", folderIds).
				And("user_id = ?", userId).
				And("del_flag = ?", constant.StatusFolderUndeleted).
				Find(&folders); err != nil {
				return nil, err
			}

			folderIds = []int64{}
			for _, folder := range folders {
				folderIds = append(folderIds, folder.Id)
			}
			folders = []*model.Folder{}
		}
		return nil, nil
	})
	if err != nil {
		return errors.New("删除过程出错！" + err.Error())
	}

	return nil
}
