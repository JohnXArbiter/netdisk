package upload

import (
	"context"
	"fmt"
	"github.com/yitter/idgenerator-go/idgen"
	"github.com/zeromicro/go-zero/core/logx"
	"lc/netdisk/common"
	"lc/netdisk/common/constant"
	"lc/netdisk/common/redis"
	"lc/netdisk/common/variable"
	"lc/netdisk/common/xorm"
	"lc/netdisk/internal/svc"
	"lc/netdisk/internal/types"
	"lc/netdisk/model"
	"mime/multipart"
	"strconv"
	"time"
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
		checkKey = fmt.Sprintf(redis.UploadCheckChunkKeyF, req.FileId, req.ChunkSeq)
	)

	objectName, err := rdb.Get(l.ctx, checkKey).Result()
	if err != nil {
		return err
	}

	bigFileKey := redis.UploadCheckBigFileKey + strconv.FormatInt(req.FileId, 10)
	fileInfo, err := rdb.HGetAll(l.ctx, bigFileKey).Result()
	if err != nil {
		return err
	}

	chunkSum, _ := strconv.ParseInt(fileInfo["chunkSum"], 10, 64)
	chunkNum, _ := strconv.ParseInt(fileInfo["chunkNum"], 10, 64)

	if chunkSum+1 == chunkNum {
		_, err = engine.DoTransaction(l.createSchedule(req, fileParam.File, objectName, chunkNum, fileInfo))
	} else {
		if err = l.svcCtx.Minio.NewService().Upload(l.ctx, objectName, fileParam.File); err != nil {
			return err
		}
		_, err = l.incr(bigFileKey, 1)
	}
	_, _ = rdb.Del(l.ctx, checkKey).Result()
	return nil
}

func (l *UploadChunkLogic) createSchedule(req *types.UploadChunkReq, fileData multipart.File,
	objectName string, chunkNum int64, fileInfo map[string]string) xorm.TxFn {
	return func(session *xorm.Session) (interface{}, error) {

		size, _ := strconv.ParseInt(fileInfo["size"], 10, 64)
		fsId := idgen.NextId()
		fileFs := &model.FileFs{}
		fileFs.Id = fsId
		fileFs.Bucket = l.svcCtx.Minio.BucketName
		fileFs.Ext = fileInfo["ext"]
		fileFs.Name = fileInfo["name"]
		fileFs.Hash = fileInfo["hash"]
		fileFs.Size = size
		fileFs.Url = ""
		fileFs.ObjectName = objectName
		fileFs.ChunkNum = chunkNum
		fileFs.Status = constant.StatusFsFileNeedMerge
		if _, err := session.Insert(fileFs); err != nil {
			return nil, err
		}

		userId, _ := strconv.ParseInt(fileInfo["userId"], 10, 64)
		folderId, _ := strconv.ParseInt(fileInfo["folderId"], 10, 64)
		file := model.File{}
		file.Name = fileInfo["name"]
		file.UserId = userId
		file.FsId = fsId
		file.FolderId = folderId
		file.Url = ""
		file.Ext = fileInfo["ext"]
		file.ObjectName = objectName
		file.Size = size
		file.Type = variable.GetTypeByBruteForce(fileInfo["ext"])
		file.Status = constant.StatusFileNeedMerge
		file.IsBig = constant.BigFileFlag
		file.DoneAt = time.Now().Local()
		if _, err := session.Insert(file); err != nil {
			return nil, err
		}

		fileScheduleId := idgen.NextId()
		fileSchedule := &model.FileSchedule{}
		fileSchedule.Id = fileScheduleId
		fileSchedule.FileId = req.FileId
		fileSchedule.FsId = fsId
		fileSchedule.ChunkNum = chunkNum
		fileSchedule.Stage = constant.StageMerging
		if _, err := session.Insert(fileSchedule); err != nil {
			return nil, err
		}

		key := redis.UploadCheckBigFileKey + strconv.FormatInt(req.FileId, 10)
		if _, err := l.incr(key, 1); err != nil {
			return nil, err
		}

		if err := l.svcCtx.Minio.NewService().Upload(l.ctx, objectName, fileData); err != nil {
			l.incr(key, -1)
			return nil, err
		}

		ms := &common.MergeStruct{}
		ms.ObjectName = objectName
		ms.Hash = fileInfo["hash"]
		ms.FsId = fsId
		ms.SId = fileScheduleId
		ms.ChunkNum = chunkNum
		go common.Merge(ms, l.errCallBack)

		return nil, nil
	}
}

func (l *UploadChunkLogic) incr(key string, value int64) (int64, error) {
	return l.svcCtx.Redis.HIncrBy(l.ctx, key, "chunkSum", value).Result()
}

func (l *UploadChunkLogic) errCallBack(sId int64) {
	fs := &model.FileSchedule{Stage: constant.StageNeedMerge}
	if _, err := l.svcCtx.Xorm.ID(sId).Update(fs); err != nil {
		logx.Errorf("createSchedule->errCallBack，退回schedule:%v 状态出错，ERR：[%v]", sId, err)
		return
	}
}
