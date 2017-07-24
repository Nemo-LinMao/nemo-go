package database

import (
	"database/sql"
)

const {
	DB_SQLITE3 = iota
	DB_MYSQL   = iota
	DB_ORACLE  = iota
}

type hdb_t *DB

type DbInterface interface {
	CreateDbConnect(dbtype int, dbname string)
	ExecSql(strsql string, interface{})
	ExecRpc(rpcname string, interface{})
}