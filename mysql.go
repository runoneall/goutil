package goutil

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Mysql_MakeDSN(host string, port int, user string, password string, dbname string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, password, host, port, dbname)
}

func Mysql_Connect(dsn string) *sql.DB {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}

func Mysql_Exec(db *sql.DB, sql string) sql.Result {
	defer func() {
		if err := recover(); err != nil {
			return
		}
	}()
	result, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
	return result
}

func Mysql_Query(db *sql.DB, sql string) []map[string]interface{} {
	defer func() {
		if err := recover(); err != nil {
			return
		}
	}()
	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		panic(err)
	}
	results := make([]map[string]interface{}, 0)
	values := make([]interface{}, len(columns))
	for rows.Next() {
		for i := range values {
			values[i] = new(interface{})
		}
		if err := rows.Scan(values...); err != nil {
			panic(err)
		}
		rowMap := make(map[string]interface{})
		for i, col := range columns {
			col_content := *(values[i].(*interface{}))
			if bytes, ok := col_content.([]byte); ok {
				col_content = string(bytes)
			}
			rowMap[col] = col_content
		}
		results = append(results, rowMap)
	}
	return results
}

func Mysql_Close(db *sql.DB) {
	db.Close()
}
