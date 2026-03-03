package sqllitedemo

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

/*
SQLite是一个进程内的库，实现了自给自足的、无服务器的、零配置的、事务性的 SQL 数据库引擎。它是一个零配置的数据库，这意味着与其他数据库不一样，您不需要在系统中配置。

就像其他数据库，SQLite 引擎不是一个独立的进程，可以按应用程序需求进行静态或动态连接。SQLite 直接访问其存储文件。

data types
+------------------------------+
|go        | sqlite3           |
|----------|-------------------|
|nil       | null              |
|int       | integer           |
|int64     | integer           |
|float64   | float             |
|bool      | integer           |
|[]byte    | blob              |
|string    | text              |
|time.Time | timestamp/datetime|
+------------------------------+
*/

var lite3 *sql.DB

func init() {
	db, err := sql.Open("sqlite3", "./database.db")

	if err != nil {
		log.Fatalf("Error sql.Open(\"sqlite3\") %s\n", err)
	}
	log.Println(db)
	lite3 = db
}
