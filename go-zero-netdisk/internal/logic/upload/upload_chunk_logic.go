package upload

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"lc/netdisk/common/redis"
	"lc/netdisk/internal/svc"
	"lc/netdisk/internal/types"
	"lc/netdisk/model"
	"strconv"
)

type UploadChunkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadChunkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadChunkLogic {
	return &UploadChunkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadChunkLogic) UploadChunk(req *types.UploadChunkReq, fileParam *types.FileParam) error {
	var (
		rdb      = l.svcCtx.Redis
		engine   = l.svcCtx.Xorm
		minioSvc = l.svcCtx.Minio.NewService()
	)

	objectName, err := rdb.Get(l.ctx, fmt.Sprintf(redis.UploadCheckChunkKeyF, req.FileId, req.ChunkSeq)).Result()
	if err != nil {
		return err
	}

	if err = minioSvc.Upload(l.ctx, objectName, fileParam.File); err != nil {
		return err
	}

	fileInfo, err := rdb.HGetAll(l.ctx, redis.UploadCheckBigFileKey+strconv.FormatInt(req.FileId, 10)).Result()
	if err != nil {
		return err
	}

	chunkSum, _ := strconv.ParseInt(fileInfo["chunkSum"], 10, 64)
	chunkNum, _ := strconv.ParseInt(fileInfo["chunkNum"], 10, 64)

	// 还有fs

	if chunkSum+1 == chunkNum {
		fileSchedule := &model.FileSchedule{}
		fileSchedule.FileId = req.FileId
		fileSchedule.FsId =
			engine.Insert()
	}

	return nil
}
