package database

import (
	"fmt"
	"database/sql"
	mysql "github.com/go-sql-driver/mysql"
)

type Mysqldb struct {
	CreateDbConnect(dbtype int, dbname string)
	ExecSql(strsql string, out interface{})
}

func (obj Mysqldb) CreateDbConnect(dbname string, user string, pwd string) hdb_t {
	connectStr := fmt.Sprintf("%s:%s@%s:%d/%s", user, pwd, localhost, 3306, "nemo-go")
	hdb, err := sql.Open("mysql", connectStr)

	if nil == err {
		return hdb
	} else {
		return nil
	}
}


func (obj Mysqldb) ExecSql(strsql string, out interface{}) {
	
}