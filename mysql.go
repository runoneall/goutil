package goutil

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Mysql_Connect(host, port, user, password, dbname string) *sql.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbname)
	db, _ := sql.Open("mysql", dsn)
	return db
}

func Mysql_Exec(db *sql.DB, sql string, args ...interface{}) sql.Result {
	result, _ := db.Exec(sql, args...)
	return result
}

func Mysql_Query(db *sql.DB, sql string, args ...interface{}) *sql.Rows {
	rows, _ := db.Query(sql, args...)
	return rows
}

func Mysql_Close(db *sql.DB) {
	db.Close()
}
