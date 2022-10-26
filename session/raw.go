package session

import (
	"database/sql"
	"fmt"
	"geeORM/dialect"
	"geeORM/log"
	"geeORM/schema"
	"reflect"
	"strings"
)

// Session 用于实现与数据库的交互
type Session struct {
	db        *sql.DB
	dialect   dialect.Dialect
	refTable  *schema.Schema
	sql       strings.Builder
	sqlValues []interface{}
}

func New(db *sql.DB, dialect dialect.Dialect) *Session {
	return &Session{
		db:      db,
		dialect: dialect,
	}
}

// Clear 用于每次执行sql后清除参数 使之可以复用
func (s *Session) Clear() {
	s.sql.Reset()
	s.sqlValues = nil
}

func (s *Session) DB() *sql.DB {
	return s.db
}

func (s *Session) Raw(sql string, value ...interface{}) *Session {
	s.sql.WriteString(sql)
	s.sql.WriteString(" ")
	s.sqlValues = append(s.sqlValues, value...)
	return s
}

func (s *Session) Exec() (result sql.Result, err error) {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlValues)
	if result, err = s.DB().Exec(s.sql.String(), s.sqlValues); err != nil {
		log.Error(err)
	}
	return
}

func (s *Session) QueryRow() *sql.Row {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlValues)
	row := s.db.QueryRow(s.sql.String(), s.sqlValues)
	return row
}

func (s *Session) QueryRows() (rows *sql.Rows, err error) {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlValues)
	if rows, err = s.DB().Query(s.sql.String(), s.sqlValues); err != nil {
		log.Error(err)
	}
	return
}

func (s *Session) Model(value interface{}) *Session {
	if s.refTable == nil || reflect.TypeOf(value) != reflect.TypeOf(s.refTable.Model) {
		s.refTable = schema.Parse(value, s.dialect)
	}
	return s
}

func (s *Session) RefTable() *schema.Schema {
	if s.refTable == nil {
		log.Error("model is not set")
	}
	return s.refTable
}

func (s *Session) CreateTable() error {
	table := s.refTable
	var columns []string
	for _, v := range table.Fields {
		columns = append(columns, fmt.Sprintf("%s %s %s", v.Name, v.Type, v.Tag))
	}
	desc := strings.Join(columns, ",")
	_, err := s.Raw(fmt.Sprintf("create table %s (%s)", table.Name, desc)).Exec()
	return err
}

func (s *Session) DropTable() error {
	_, err := s.Raw(fmt.Sprintf("drop table if exist %s", s.refTable.Name)).Exec()
	return err
}

func (s *Session) HasTable() bool {
	sql, values := s.dialect.TableExistSQL(s.RefTable().Name)
	result := s.Raw(sql, values...).QueryRow()
	var tableName string
	err := result.Scan(&tableName)
	if err != nil {
		log.Error("get result err")
	}
	return tableName == s.refTable.Name
}
