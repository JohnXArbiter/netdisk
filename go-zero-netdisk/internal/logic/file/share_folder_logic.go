package file

import (
	"context"
	"errors"
	"github.com/yitter/idgenerator-go/idgen"
	"lc/netdisk/common/constant"
	"lc/netdisk/common/variable"
	"lc/netdisk/common/xorm"
	"lc/netdisk/internal/logic/mqs"
	"lc/netdisk/internal/svc"
	"lc/netdisk/internal/types"
	"lc/netdisk/model"
	"strconv"
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
		folderIds  = req.FolderIds
		folders    []*model.Folder
		shareFiles []*model.ShareFile
		err        error
	)

	defer mqs.LogSend(l.ctx, err, "ShareFolder", req.FolderIds)

	if req.Pwd == "" {
		return errors.New("出错啦，请重试")
	}

	id := strconv.FormatInt(idgen.NextId(), 10)
	url := req.Prefix + id + "?pwd=" + req.Pwd
	if has, err2 := engine.In("id", folderIds).
		Get(&folders); err2 != nil {
		logx.Errorf("分享文件夹，查询folder失败，ERR: [%v]", err2)
		err = errors.New("出错了，" + err2.Error())
		return err
	} else if !has {
		err = errors.New("信息有误")
		return err
	}

	folderName := folders[0].Name
	for len(folderIds) > 0 {
		// 1.获取当前文件夹下的文件
		var fileIds []int64
		if err = engine.In("folder_id", folderIds).
			And("user_id = ?", userId).
			And("del_flag = ?", constant.StatusFileUndeleted).
			Find(&fileIds); err != nil {
			logx.Errorf("分享文件夹，插入shareFile失败，ERR: [%v]", err)
			return err
		}

		// 2.搜索下一层文件夹
		var folderIds2 []int64
		if err = engine.Select("id").In("parent_id", folderIds).
			And("user_id = ?", userId).
			And("del_flag = ?", constant.StatusFolderUndeleted).
			Find(&folderIds2); err != nil {
			logx.Errorf("分享文件夹，插入shareFile失败，ERR: [%v]", err)
			return err
		}

		folderIds = folderIds2
		for _, fileId := range fileIds {
			shareFiles = append(shareFiles, &model.ShareFile{
				ShareId: id,
				FileId:  fileId,
			})
		}
	}

	_, err = engine.DoTransaction(func(session *xorm.Session) (interface{}, error) {
		if _, err := session.Insert(shareFiles); err != nil {
			logx.Errorf("分享文件夹，插入shareFile失败，ERR: [%v]", err)
			return nil, err
		}

		created := time.Now().Local()
		expired := created.Unix() + variable.ShareExpireType[req.ExpireType]
		share := &model.Share{}
		share.Id = id
		share.Pwd = req.Pwd
		share.Name = folderName
		share.UserId = userId
		share.Created = created
		share.Expired = expired
		share.Type = constant.TypeShareMulti
		share.Url = url
		if _, err := session.Insert(share); err != nil {
			logx.Errorf("分享文件夹，插入share失败，ERR: [%v]", err)
			return nil, err
		}
		return nil, nil
	})
	return err
}
