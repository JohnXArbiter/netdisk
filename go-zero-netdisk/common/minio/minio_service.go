package minio

import (
	"context"
	"fmt"
	"github.com/minio/minio-go"
	"log"
	"os"
)

type (
	Service struct {
		client *minio.Client
	}
)

func (c *Client) NewService() *Service {
	return &Service{c.client}
}

// UploadFile 上传文件
func (s *Service) UploadFile(ctx context.Context, bucketName, objectName string, file *os.File) error {
	_, err := s.client.PutObjectWithContext(ctx, bucketName, objectName, file, -1, minio.PutObjectOptions{})
	if err != nil {
		log.Println("putObject fail: ", err)
		return err
	}

	fmt.Println("Successfully uploaded", objectName)

	return nil
}

//// DownloadFile 下载文件
//func (s *Service) DownloadFile(bucketName, objectName, filePath string) error {
//	// 创建本地文件
//	file, err := os.Create(filePath)
//	if err != nil {
//		return err
//	}
//	defer file.Close()
//
//	// 下载存储桶中的文件到本地
//	err = s.Client.FGetObject(bucketName, objectName, filePath, minio.GetObjectOptions{})
//	if err != nil {
//		return err
//	}
//
//	fmt.Println("Successfully downloaded", objectName)
//	return nil
//}
//
//// DeleteFile 删除文件
//func (s *Service) DeleteFile(bucketName, objectName string) (bool, error) {
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
//
//// ListObjects 列出文件
//func (s *Service) ListObjects(bucketName, prefix string) ([]string, error) {
//	var objectNames []string
//
//	for object := range s.Client.ListObjects(bucketName, prefix, true, nil) {
//		if object.Err != nil {
//			return nil, object.Err
//		}
//
//		objectNames = append(objectNames, object.Key)
//	}
//
//	return objectNames, nil
//}
