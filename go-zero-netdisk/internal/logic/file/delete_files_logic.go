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

type DeleteFilesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteFilesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteFilesLogic {
	return &DeleteFilesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteFilesLogic) DeleteFiles(req *types.IdsReq) error {
	var (
		userId = l.ctx.Value(constant.UserIdKey).(int64)
		engine = l.svcCtx.Xorm
	)

	cond := &model.File{
		DelFlag: constant.StatusFileDeleted,
		DelTime: time.Now().Local().Unix(),
	}
	if affected, err := engine.In("id", req.Ids).And("user_id = ?", userId).
		Update(cond); err != nil || affected != int64(len(req.Ids)) {
		return errors.New("删除过程出错！" + err.Error())
	}
	return nil
}
