package svc

import (
	"github.com/yitter/idgenerator-go/idgen"
	"github.com/zeromicro/go-zero/rest"
	"lc/netdisk/common/minio"
	"lc/netdisk/common/redis"
	"lc/netdisk/common/xorm"
	"lc/netdisk/internal/config"
	"lc/netdisk/internal/middleware"
)

type ServiceContext struct {
	Config config.Config
	Minio  *minio.Client
	Xorm   *xorm.Engine
	Redis  *redis.Client
	Auth   rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	idgenops := idgen.NewIdGeneratorOptions(c.Idgen.WorkerId)
	idgen.SetIdGenerator(idgenops)

	minioClient := minio.Init(&c.Minio)

	xormEngine := xorm.Init(&c.Xorm)
	redisClient := redis.Init(&c.Redis)
	return &ServiceContext{
		Config: c,
		Minio:  minioClient,
		Xorm:   xormEngine,
		Redis:  redisClient,
		Auth:   middleware.NewAuthMiddleware().Handle,
	}
}
