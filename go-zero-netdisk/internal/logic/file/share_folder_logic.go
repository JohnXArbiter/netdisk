package file

import (
	"context"
	"lc/netdisk/common/constant"
	"lc/netdisk/common/variable"
	"lc/netdisk/common/xorm"
	"lc/netdisk/internal/svc"
	"lc/netdisk/internal/types"
	"lc/netdisk/model"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShareFolderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareFolderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareFolderLogic {
	return &ShareFolderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareFolderLogic) ShareFolder(req *types.ShareFolderReq) error {
	var (
		userId     = l.ctx.Value(constant.UserIdKey).(int64)
		engine     = l.svcCtx.Xorm
		folderIds  = []int64{req.FolderId}
		shareFiles []*model.ShareFile
	)

	for len(folderIds) > 0 {
		// 1.获取当前文件夹下的文件
		var fileIds []int64
		if err := engine.In("folder_id", folderIds).
			And("user_id = ?", userId).
			And("del_flag = ?", constant.StatusFileUndeleted).
			Find(&fileIds); err != nil {
			logx.Errorf("分享文件夹，插入shareFile失败，ERR: [%v]", err)
			return err
		}

		// 2.搜索下一层文件夹
		var folderIds2 []int64
		if err := engine.Select("id").In("parent_id", folderIds).
			And("user_id = ?", userId).
			And("del_flag = ?", constant.StatusFolderUndeleted).
			Find(&folderIds2); err != nil {
			logx.Errorf("分享文件夹，插入shareFile失败，ERR: [%v]", err)
			return err
		}

		folderIds = folderIds2
		for _, fileId := range fileIds {
			shareFiles = append(shareFiles, &model.ShareFile{
				ShareId: req.Id,
				FileId:  fileId,
			})
		}
	}

	_, err := engine.DoTransaction(func(session *xorm.Session) (interface{}, error) {
		if _, err := session.Insert(shareFiles); err != nil {
			logx.Errorf("分享文件夹，插入shareFile失败，ERR: [%v]", err)
			return nil, err
		}

		created := time.Now().Unix()
		expired := created + variable.ShareExpireType[req.ExpireType]
		share := &model.Share{}
		share.Id = req.Id
		share.UserId = userId
		share.Created = created
		share.Expired = expired
		if _, err := session.Insert(share); err != nil {
			logx.Errorf("分享文件夹，插入share失败，ERR: [%v]", err)
			return nil, err
		}
		return nil, nil
	})
	return err
}
