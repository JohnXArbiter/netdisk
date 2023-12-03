package common

import (
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

type DbConf struct {
	Dsn string
}

func InitXorm(conf *DbConf) *xorm.Engine {

	engine, err := xorm.NewEngine("mysql", conf.Dsn)
	//logx.Infof("[XORM CONNECTING] Init Xorm DSN: %v", conf.Dsn)
	if err != nil {
		panic("[XORM ERROR] NewServiceContext 获取mysql连接错误 " + err.Error())
	}
	err = engine.Ping()
	if err != nil {
		panic("[XORM ERROR] NewServiceContext ping mysql 失败" + err.Error())
	}
	return engine
}
