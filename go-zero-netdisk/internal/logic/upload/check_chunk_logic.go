package upload

import (
	"context"
	"fmt"
	"lc/netdisk/common/redis"
	"lc/netdisk/internal/svc"
	"lc/netdisk/internal/types"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckChunkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckChunkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckChunkLogic {
	return &CheckChunkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckChunkLogic) CheckChunk(req *types.CheckChunkReq) error {
	var (
		rdb      = l.svcCtx.Redis
		minio    = l.svcCtx.Minio
		minioSvc = l.svcCtx.Minio.NewService()
	)

	fileIdStr := strconv.FormatInt(req.FileId, 10)
	_, err := rdb.Exists(l.ctx, redis.UploadCheckBigFileKey+fileIdStr).Result()
	if err != nil {
		return err
	}

	objectName := minio.GenChunkObjectName(req.Hash, req.ChunkSeq)
	if err = minioSvc.IfExist(objectName); err != nil {
		return err
	}

	if err = rdb.Set(l.ctx, fmt.Sprintf(redis.UploadCheckChunkKeyF, req.FileId, req.ChunkSeq),
		objectName, redis.UploadCheckChunkExpire).Err(); err != nil {
		return err
	}

	return nil
}
