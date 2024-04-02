package common

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"io"
	"lc/netdisk/common/constant"
	"lc/netdisk/common/minio"
	"lc/netdisk/common/xorm"
	"lc/netdisk/model"
	"os"
	"strconv"
)

type MergeStruct struct {
	FsId       int64
	SId        int64
	ObjectName string
	Hash       string
	ChunkNum   int64
}

func MergeLogic() {
	//mqs.LogSend(context.Background(), nil, "MergeLogic", time.Now().Local())

	var mss []*MergeStruct
	cols := "a.id as sId, b.id as fsId, b.object_name, b.hash, b.chunk_num"
	if err := xorm.Xorm.Select(cols).
		Table(&model.FileSchedule{}).Alias("a").
		Join("LEFT", []string{"file_fs", "b"}, "a.fs_id = b.id").
		Where("a.stage = ?", constant.StageNeedMerge).
		Limit(1000).Find(&mss); err != nil {
		logx.Errorf("MergeLogic，查询fileSchedule出错，ERR：[%v]", err)
		return
	}

	q := make(chan struct{}, 5)
	for _, ms1 := range mss {
		ms2 := ms1
		go func() {
			q <- struct{}{}
			Merge(ms2, func(int64) {})
			<-q
		}()
	}
}

func Merge(ms *MergeStruct, errCallBack func(int64)) {
	var err error
	defer func() {
		if err != nil {
			errCallBack(ms.SId)
		}
	}()

	minioSvc := minio.Minio.NewService()
	bigFile, err := os.CreateTemp("", "netdisk")
	if err != nil {
		logx.Errorf("MergeLogic，创建临时文件出错，ERR：[%v]", err)
		return
	}
	defer DeleteTemp(bigFile)

	var chunks []*os.File
	defer deleteChunks(chunks)
	for i := 0; int64(i) < ms.ChunkNum; i++ {
		objectName := ms.ObjectName + "@" + strconv.Itoa(i)
		fileName := ms.Hash + "@" + strconv.Itoa(i)
		logx.Info(objectName, "    ", fileName)
		chunk, err2 := minioSvc.DownloadChunk(context.TODO(), objectName, fileName)
		if err2 != nil {
			logx.Errorf("MergeLogic，文件%v 下载分片[%v] 失败，ERR：[%v]", ms.Hash, i, err2)
			err = err2
			return
		}

		fileInfo, err2 := chunk.Stat()
		if err2 != nil {
			logx.Errorf("MergeLogic，文件%v 下载分片[%v]有误，ERR：[%v]", ms.Hash, i, err2)
			err = err2
			return
		}
		chunks = append(chunks, chunk)
		buffer := make([]byte, fileInfo.Size())
		_, err2 = io.CopyBuffer(bigFile, chunk, buffer)
		if err2 != nil {
			logx.Errorf("MergeLogic，文件%v 合并分片[%v]出错，ERR：[%v]", ms.Hash, i, err2)
			err = err2
			return
		}
	}

	_, err = xorm.Xorm.DoTransaction(func(session *xorm.Session) (interface{}, error) {
		fs1 := &model.FileFs{Status: constant.StatusFsUploaded}
		if _, err = session.ID(ms.FsId).Update(fs1); err != nil {
			logx.Errorf("MergeLogic，文件%v 更新fs状态出错，ERR：[%v]", ms.Hash, err)
			return nil, err
		}

		file := &model.File{Status: constant.StatusFileUploaded}
		if _, err = session.Where("fs_id = ?", ms.FsId).
			Update(file); err != nil {
			logx.Errorf("MergeLogic，文件%v 更新file状态出错，ERR：[%v]", ms.Hash, err)
			return nil, err
		}

		fs2 := &model.FileSchedule{Stage: constant.StageMergeDone}
		if _, err = session.ID(ms.SId).Update(fs2); err != nil {
			logx.Errorf("MergeLogic，文件%v 更新fileSchedule状态出错，ERR：[%v]", ms.Hash, err)
			return nil, err
		}

		if err = minioSvc.Upload(context.TODO(), ms.ObjectName, bigFile); err != nil {
			logx.Errorf("MergeLogic，上传大文件：[%v]，出错，ERR：[%v]", ms.ObjectName, err)
			return nil, err
		}
		return nil, nil
	})
}

func DeleteTemp(temp *os.File) {
	name := temp.Name()
	if err := temp.Close(); err != nil {
		logx.Errorf("DeleteTemp，关闭临时文件 %v 出错，ERR：[%v]", name, err)
	}
	if err := os.Remove(name); err != nil {
		logx.Errorf("DeleteTemp，删除临时文件 %v 出错，ERR：[%v]", name, err)
	}
}

func deleteChunks(chunks []*os.File) {
	for _, chunk := range chunks {
		if chunk == nil {
			continue
		}
		if err := os.Remove(chunk.Name()); err != nil {
			logx.Errorf("deleteChunks，删除分片临时文件出错，ERR：[%v]", err)
			return
		}
	}
}
