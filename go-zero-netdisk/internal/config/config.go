package config

import (
	"github.com/zeromicro/go-zero/rest"
	"lc/netdisk/common/minio"
	"lc/netdisk/common/xorm"
)

type Config struct {
	rest.RestConf
	Minio minio.Conf
	Xorm  xorm.DbConf
}
