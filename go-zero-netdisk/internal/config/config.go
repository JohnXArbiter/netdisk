package config

import (
	"github.com/zeromicro/go-zero/rest"
	"lc/netdisk/common"
	"lc/netdisk/common/minio"
	"lc/netdisk/common/redis"
	"lc/netdisk/common/xorm"
)

type Config struct {
	rest.RestConf
	Minio minio.Conf
	Xorm  xorm.DbConf
	Redis redis.Conf
	Email common.Email
	Idgen struct {
		WorkerId uint16
	}
	KqPusherConfs []*KqPusherConf
}

type KqPusherConf struct {
	Type    string
	Brokers []string
	Topic   string
}
