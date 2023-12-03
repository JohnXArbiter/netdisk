package svc

import (
	"github.com/minio/minio-go"
	"lc/netdisk/common"
	"lc/netdisk/internal/config"
)

type ServiceContext struct {
	Config      config.Config
	MinioClient *minio.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	minioClient := common.InitMinio(&c.Minio)

	return &ServiceContext{
		Config:      c,
		MinioClient: minioClient,
	}
}
