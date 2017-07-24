package database

import (
	"database/sql"
	gosqlite "github.com/mattn/go-sqlite3"
)

type Sqlite3db struct {
	CreateDbConnect(dbtype int, dbname string)
	ExecSql(strsql string, interface{})
}

func (obj Sqlite3) CreateDbConnect(dbtype int, dbname string) int {
	db, err := sql.Open("sqlite9", "./foo.db")
}