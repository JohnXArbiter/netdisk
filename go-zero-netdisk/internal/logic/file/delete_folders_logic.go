package file

import (
	"context"
	"errors"
	"lc/netdisk/common/constant"
	"lc/netdisk/model"
	"time"

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

	bean := &model.Folder{
		DelFlag: constant.StatusFileDeleted,
		DelTime: time.Now().Local().Unix(),
	}
	if affected, err := engine.In("id", req.Ids).
		And("user_id = ?", userId).Update(bean); err != nil {
		return errors.New("删除过程出错！" + err.Error())
	} else if affected != int64(len(req.Ids)) {
		return errors.New("无法删除，可能信息有误？😵‍💫")
	}
	return nil
}
