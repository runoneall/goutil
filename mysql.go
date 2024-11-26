package goutil

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Mysql_Connect(host, port, user, password, dbname string) *sql.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbname)
	db, _ := sql.Open("mysql", dsn)
	return db
}

func Mysql_Exec(db *sql.DB, sql string, args ...interface{}) (sql.Result, error) {
	result, err := db.Exec(sql, args...)
	if err != nil {
		log.Fatal(err)
	}
	return result, err
}

func Mysql_Query(db *sql.DB, sql string, args ...interface{}) (*sql.Rows, error) {
	rows, err := db.Query(sql, args...)
	if err != nil {
		log.Fatal(err)
	}
	return rows, err
}

func Mysql_QueryRow(db *sql.DB, sql string, args ...interface{}) *sql.Row {
	row := db.QueryRow(sql, args...)
	return row
}
func Mysql_Close(db *sql.DB) {
	db.Close()
}
