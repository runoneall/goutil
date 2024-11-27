package goutil

import (
	mysql "database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Mysql_ToUrl(host string, port int, user string, password string, dbname string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, password, host, port, dbname)
}

func Mysql_Connect(connUrl string) *mysql.DB {
	db, err := mysql.Open("mysql", connUrl)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func Mysql_Exec(db *mysql.DB, sql string) mysql.Result {
	result, err := db.Exec(sql)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func Mysql_Query(db *mysql.DB, sql string) *mysql.Rows {
	rows, err := db.Query(sql)
	if err != nil {
		log.Fatal(err)
	}
	return rows
}

func Mysql_Close(db *mysql.DB) {
	db.Close()
}
