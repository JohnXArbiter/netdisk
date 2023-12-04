package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"lc/netdisk/common/minio"
	"lc/netdisk/common/xorm"
	"lc/netdisk/internal/config"
	"lc/netdisk/internal/middleware"
)

type ServiceContext struct {
	Config config.Config
	Minio  *minio.Client
	Xorm   *xorm.Engine

	Auth rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	minioClient := minio.Init(&c.Minio)

	xormEngine := xorm.Init(&c.Xorm)

	return &ServiceContext{
		Config: c,
		Minio:  minioClient,
		Xorm:   xormEngine,
		Auth:   middleware.NewAuthMiddleware().Handle,
	}
}
