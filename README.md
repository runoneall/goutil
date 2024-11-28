# r1a utilities for go

## command.go
- `Command_Args()` return type `map[string]string` parameter dictionary
- `Command_GetArg(arg string)` return type `string` value of parameter `arg` if not exist return empty string

## config.go
- `Config_Read_File(filepath string)` return type `map[string]interface{}` config map from file
- `Config_Write_File(filepath string, configMap map[string]interface{})` none return type write config map to file
- `Config_Format(content string, targetFormat string)` return type `map[string]interface{}` load config from string and format to `targetFormat`

## date.go
- `Date_Format(formatQuery string, useLong bool, useAlias bool)` return type `string` format date with `formatQuery` use `{{year}}` as year, `{{month}}` as month, `{{day}}` as day, `{{hour}}` as hour, `{{minute}}` as minute, `{{second}}` as second, `{{week}}` as weekday; set `useLong` as `true` use long name, set `useAlias` as `true` use alias name

## mysql.go
- `Mysql_MakeDSN(host string, port int, user string, password string, dbname string)` return type `string` make DSN string for mysql connection
- `Mysql_Connect(dsn string)` return type `*sql.DB` connect to mysql database with DSN string
- `Mysql_Exec(db *sql.DB, sql string)` return type `sql.Result` execute SQL statement and return result
- `Mysql_Query(db *sql.DB, sql string)` return type `[]map[string]interface{}` execute SQL statement and return result as map array
- `Mysql_Close(db *sql.DB)` none return type close mysql database connection

## network.go
