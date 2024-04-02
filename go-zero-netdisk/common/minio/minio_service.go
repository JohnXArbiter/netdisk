package minio

import (
	"context"
	"fmt"
	"github.com/minio/minio-go"
	"github.com/zeromicro/go-zero/core/logx"
	"io"
	"net/url"
	"os"
	"strings"
	"time"
)

type (
	Service struct {
		BucketName string
		client     *minio.Client
	}
)

func (c *Client) NewService() *Service {
	return &Service{c.BucketName, c.client}
}

// Upload 上传文件
func (s *Service) Upload(ctx context.Context, objectName string, file io.Reader) error {
	_, err := s.client.PutObjectWithContext(ctx, s.BucketName, objectName, file, -1, minio.PutObjectOptions{})
	fmt.Println(s.BucketName)
	if err != nil {
		logx.Errorf("minio-上传文件出错，err: %v", err)
		return err
	}
	return nil
}

func (s *Service) IfExist(objectName string) (bool, error) {
	_, err := s.client.StatObject(s.BucketName, objectName, minio.StatObjectOptions{})
	if err != nil {
		if strings.Contains(err.Error(), "The specified key does not exist") {
			return false, nil
		}
		logx.Errorf("minio-判断文件是否存在出错，err: %v", err)
		return false, err
	}
	return true, nil
}

func (s *Service) GenUrl(objectName string, filename string, download bool) (string, error) {
	var (
		u   *url.URL
		err error
	)

	if download {
		kvs := url.Values{}
		kvs["response-content-disposition"] = []string{"attachment; filename=" + objectName}
		if filename != "" {
			kvs["response-content-disposition"] = []string{"attachment; filename=" + filename}
		}
		u, err = s.client.PresignedGetObject(s.BucketName, objectName, 7*24*time.Hour, kvs)
	} else {
		u, err = s.client.PresignedGetObject(s.BucketName, objectName, 7*24*time.Hour, nil)
	}

	if err != nil {
		logx.Errorf("minio-获取下载url出错，err: %v", err)
		return "", err
	}
	return fmt.Sprintf("%v", u), nil
}

// DownloadChunk 下载文件切片
func (s *Service) DownloadChunk(ctx context.Context, objectName, name string) (*os.File, error) {
	filename := os.TempDir() + "/" + name
	file, err := os.Create(filename)
	if err != nil {
		logx.Errorf("DownloadChunk，minio下载文件出错，ERR: [%v]", err)
		return nil, err
	}

	if err = s.client.FGetObjectWithContext(ctx, s.BucketName, objectName,
		filename, minio.GetObjectOptions{}); err != nil {
		logx.Errorf("DownloadChunk，minio下载出错，ERR: [%v]", err)
		return nil, err
	}

	return file, nil
}

//
//// DeleteFile 删除文件
//func (s *Service) DeleteFile(bucketName, objectName string) (bool, error) {的miniokehuduan1
//	// 删除存储桶中的文件
//	err := s.Client.RemoveObject(bucketName, objectName)
//	if err != nil {
//		log.Println("remove object fail: ", err)
//		return false, err
//	}
//
//	fmt.Println("Successfully deleted", objectName)
//	return true, err
//}
