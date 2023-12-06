package xorm

import (
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

	TxFn func(session *Session) (interface{}, error)
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

func (e *Engine) DoTransaction(fn TxFn) (interface{}, error) {
	return e.Transaction(func(session *xorm.Session) (interface{}, error) {
		return fn(&Session{session})
	})
}

//func (e *Engine) Transaction(ctx interface{}, session *Session, fns ...func(interface{}, *Session) error) error {
//
//	if session == nil {
//		session = &Session{e.NewSession()}
//	}
//
//	if err := session.Begin(); err != nil {
//		return err
//	}
//
//	defer func() {
//		var err error
//		if r := recover(); r != nil {
//			err = errors.New(fmt.Sprintf("%v", r))
//			// TODO: log
//		}
//
//		if err != nil {
//			_ = session.Rollback()
//		} else {
//			_ = session.Commit()
//		}
//	}()
//
//	for _, fn := range fns {
//		if err := fn(ctx, session); err != nil {
//			return err
//		}
//	}
//
//	return nil
//}
