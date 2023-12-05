package xorm

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

type (
	DbConf struct {
		Dsn string
	}

	Engine struct {
		*xorm.Engine
	}

	Session struct {
		*xorm.Session
	}
)

func Init(conf *DbConf) *Engine {

	engine, err := xorm.NewEngine("mysql", conf.Dsn)
	//logx.Infof("[XORM CONNECTING] Init Xorm DSN: %v", conf.Dsn)
	if err != nil {
		panic("[XORM ERROR] NewServiceContext 获取mysql连接错误 " + err.Error())
	}
	err = engine.Ping()
	if err != nil {
		panic("[XORM ERROR] NewServiceContext ping mysql 失败" + err.Error())
	}
	return &Engine{engine}
}

func Transaction(ctx context.Context, engine *Engine,
	fn func(context.Context, *Session) (interface{}, error)) (interface{}, error) {
	session := engine.NewSession()
	defer session.Close()

	if err := session.Begin(); err != nil {
		return nil, err
	}

	i, err := fn(ctx, &Session{session})
	if err != nil {
		return nil, err
	}

	err = session.Commit()
	if err != nil {
		return nil, err
	}

	return i, nil
}
