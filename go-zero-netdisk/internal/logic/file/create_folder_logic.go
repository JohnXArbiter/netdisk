package file

import (
	"context"
	"errors"
	"lc/netdisk/common/constant"
	"lc/netdisk/internal/logic/mqs"
	"lc/netdisk/internal/svc"
	"lc/netdisk/internal/types"
	"lc/netdisk/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateFolderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateFolderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateFolderLogic {
	return &CreateFolderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateFolderLogic) CreateFolder(req *types.CreateFolderReq) error {
	var (
		userId         = l.ctx.Value(constant.UserIdKey).(int64)
		engine         = l.svcCtx.Xorm
		parentFolderId = req.ParentFolderId
		err            error
	)

	defer mqs.LogSend(l.ctx, err, "CreateFolder", req.Name)

	folder := &model.Folder{}
	if parentFolderId != 0 {
		if exist, err := engine.ID(parentFolderId).
			And("user_id = ?", userId).
			Exist(folder); err != nil {
			logx.Errorf("创建文件夹查询父文件夹是否存在出错，err: %v", err)
			return errors.New("发生错误，" + err.Error())
		} else if !exist {
			return errors.New("发生错误！")
		}
	}

	if exist, err := engine.Where("name = ?", req.Name).
		And("parent_id = ?", parentFolderId).Exist(folder); err != nil {
		logx.Errorf("创建文件夹查询同名文件夹是否存在出错，err: %v", err)
		return errors.New("发生错误，" + err.Error())
	} else if exist {
		return errors.New("文件夹名称已存在！")
	}

	newFolder := &model.Folder{
		ParentId: parentFolderId,
		Name:     req.Name,
		UserId:   userId,
		DelFlag:  constant.StatusFolderUndeleted,
	}
	if _, err = engine.Insert(newFolder); err != nil {
		return errors.New("创建失败了，" + err.Error())
	}

	return nil
}
