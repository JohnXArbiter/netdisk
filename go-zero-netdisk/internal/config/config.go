package config

import (
	"github.com/zeromicro/go-zero/rest"
	"lc/netdisk/common/minio"
	"lc/netdisk/common/redis"
	"lc/netdisk/common/xorm"
)

type Config struct {
	rest.RestConf
	Minio minio.Conf
	Xorm  xorm.DbConf
	Redis redis.Conf
	Idgen struct {
		WorkerId uint16
	}
}
