package goutil

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Mysql_MakeDSN(host string, port int, user string, password string, dbname string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, password, host, port, dbname)
}

func Mysql_Connect(dsn string) *sql.DB {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func Mysql_Exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func Mysql_Query(db *sql.DB, sql string) []interface{} {
	rows, err := db.Query(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var result []interface{}
	for rows.Next() {
		var row []interface{}
		err := rows.Scan(&row)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, row)
	}
	return result
}

func Mysql_Close(db *sql.DB) {
	db.Close()
}
