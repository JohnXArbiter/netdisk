package file

import (
	"context"
	"errors"
	"fmt"
	redis2 "github.com/redis/go-redis/v9"
	"lc/netdisk/common/constant"
	"lc/netdisk/common/redis"
	"lc/netdisk/internal/logic/mqs"
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

func (l *DeleteFilesLogic) DeleteFiles(req *types.DeleteFilesReq) error {
	var (
		userId = l.ctx.Value(constant.UserIdKey).(int64)
		engine = l.svcCtx.Xorm
		rdb    = l.svcCtx.Redis
		err    error
	)

	defer mqs.LogSend(l.ctx, err, "DeleteFiles", req.FileIds)

	bean := &model.File{
		DelFlag:  constant.StatusFileDeleted,
		DelTime:  time.Now().Local().Unix(),
		SyncFlag: constant.FlagSyncDelete,
	}
	if _, err = engine.In("id", req.FileIds).
		And("user_id = ?", userId).
		Update(bean); err != nil {
		logx.Errorf("删除文件失败，ERR: [%v]", err)
		return errors.New("删除过程出错，" + err.Error())
	}

	key := fmt.Sprintf(redis.FileFolderDownloadUrlKey, userId, req.FolderId)
	if err = rdb.ZRem(l.ctx, key, req.FileIds).Err(); err != nil {
		if err != redis2.Nil {
			logx.Errorf("删除文件更新redis缓存失败，ERR: [%v]", err)
		}
	}

	return nil
}
