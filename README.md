# r1a utilities for go

## command.go
- `Command_Args()` return type `map[string]string` parameter dictionary
- `Command_GetArg(arg string)` return type `string` value of parameter `arg` if not exist return empty string

## config.go
- `Config_Read_File(filepath string)` return type `map[string]interface{}` config map from file
- `Config_Write_File(filepath string, configMap map[string]interface{})` none return type write config map to file
- `Config_Format(content string, targetFormat string)` return type `map[string]interface{}` load config from string and format to `targetFormat`

## date.go
- `Date_Format(formatQuery string, useLong bool, useAlias bool)` return type `string` format date with `formatQuery`; use `{{year}}` as year, `{{month}}` as month, `{{day}}` as day, `{{hour}}` as hour, `{{minute}}` as minute, `{{second}}` as second, `{{week}}` as weekday; set `useLong` as `true` use long name, set `useAlias` as `true` use alias name

## mysql.go
- `Mysql_MakeDSN(host string, port int, user string, password string, dbname string)` return type `string` make DSN string for mysql connection
- `Mysql_Connect(dsn string)` return type `*sql.DB` connect to mysql database with DSN string
- `Mysql_Exec(db *sql.DB, sql string)` return type `sql.Result` execute SQL statement and return result
- `Mysql_Query(db *sql.DB, sql string)` return type `[]map[string]interface{}` execute SQL statement and return result as map array
- `Mysql_Close(db *sql.DB)` none return type close mysql database connection

## network.go
- `Network_Get(url string, headers ...map[string]string)` return type `*http.Response` send GET request to `url` with optional request headers `headers`
- `Network_Post(url string, data any, dataType string, headers ...map[string]string)` return type `*http.Response` send POST request to `url` with `data` as request body and `dataType` as content-type, optional request headers `headers`; when `dataType` is `text` data must be `string`, when `dataType` is `binary` data must be `[]byte`, when `dataType` is `json` data must be `map[string]interface{}`, when `dataType` is `form-kv` data must be `map[string]string`, when `dataType` is `form-file` data must be `map[string]string` with `key-name:file-path`