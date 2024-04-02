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

func (l *CheckChunkLogic) CheckChunk(req *types.CheckChunkReq) (*types.CheckChunkResp, error) {
	var (
		rdb      = l.svcCtx.Redis
		minio    = l.svcCtx.Minio
		minioSvc = l.svcCtx.Minio.NewService()
		resp     = &types.CheckChunkResp{Status: 0}
	)

	// 1. 判断之前是否正确写入了文件的数据
	fileIdStr := strconv.FormatInt(req.FileId, 10)
	_, err := rdb.Exists(l.ctx, redis.UploadCheckBigFileKey+fileIdStr).Result()
	if err != nil {
		return nil, err
	}

	// 2. 检查是否已上传
	objectName := minio.GenChunkObjectName(req.Hash, req.ChunkSeq)
	exist, err := minioSvc.IfExist(objectName)
	if err != nil {
		return resp, err
	} else if exist {
		resp.Status = 1
		return resp, nil
	}

	// 3. 为分片创建一个临时 key
	if err = rdb.Set(l.ctx, fmt.Sprintf(redis.UploadCheckChunkKeyF, req.FileId, req.ChunkSeq),
		objectName, redis.UploadCheckChunkExpire).Err(); err != nil {
		return nil, err
	}

	return resp, nil
}
