package config

import (
	"github.com/zeromicro/go-zero/rest"
	"lc/netdisk/common"
)

type Config struct {
	rest.RestConf
	Minio common.MinioConf
}
