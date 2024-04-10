package svc

import (
	"github.com/olivere/elastic/v7"
	"github.com/yitter/idgenerator-go/idgen"
	"github.com/zeromicro/go-zero/rest"
	"lc/netdisk/common"
	"lc/netdisk/common/es"
	"lc/netdisk/common/minio"
	"lc/netdisk/common/redis"
	"lc/netdisk/common/xorm"
	"lc/netdisk/internal/config"
	"lc/netdisk/internal/logic/mqs"
	"lc/netdisk/internal/middleware"
)

type ServiceContext struct {
	Config   config.Config
	Minio    *minio.Client
	MinioSvc *minio.Service
	Xorm     *xorm.Engine
	Redis    *redis.Client
	Email    *common.Email
	Auth     rest.Middleware
	Es       *elastic.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	idgenops := idgen.NewIdGeneratorOptions(c.Idgen.WorkerId)
	idgen.SetIdGenerator(idgenops)

	minioClient := minio.Init(&c.Minio)
	xormEngine := xorm.Init(&c.Xorm)
	redisClient := redis.Init(&c.Redis)
	esClient := es.Init(c.Eshost)

	mqs.NewLogPusher(c.KqPusherConfs)

	return &ServiceContext{
		Config:   c,
		Minio:    minioClient,
		MinioSvc: minioClient.NewService(),
		Xorm:     xormEngine,
		Redis:    redisClient,
		Email:    &c.Email,
		Es:       esClient,
		Auth:     middleware.NewAuthMiddleware().Handle,
	}
}
