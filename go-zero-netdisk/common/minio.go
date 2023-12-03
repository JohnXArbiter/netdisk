package common

import (
	"github.com/minio/minio-go"
	"log"
)

type MinioConf struct {
	Endpoint        string
	AccessKeyId     string
	SecretAccessKey string
	UseSSL          bool
}

// InitMinio 初使化 minio client对象。
func InitMinio(conf *MinioConf) *minio.Client {
	client, err := minio.New(conf.Endpoint, conf.AccessKeyId, conf.SecretAccessKey, conf.UseSSL)
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	return client
}
