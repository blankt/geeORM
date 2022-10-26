package main

import (
	"database/sql"
	"geeORM/dialect"
	"geeORM/log"
	"geeORM/session"
)

// Engine 用于控制数据库连接
type Engine struct {
	db      *sql.DB
	dialect dialect.Dialect
}

func NewEngine(driver, source string) (engine *Engine, err error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		log.Error(err)
		return
	}

	dial, ok := dialect.GetDialect(driver)
	if !ok {
		log.Errorf("dialect %s not found", driver)
		return
	}

	if err = db.Ping(); err != nil {
		log.Error(err)
		return
	}
	engine = &Engine{
		db:      db,
		dialect: dial,
	}
	log.Info("Connect database success")
	return
}

func (e *Engine) Close() {
	if err := e.db.Close(); err != nil {
		log.Error(err)
	}
	log.Info("close database success")
}

func (e *Engine) Session() *session.Session {
	return session.New(e.db, e.dialect)
}
