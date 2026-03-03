package mysqldemo

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type DBPool struct {
	Master *sql.DB
	Slave  *sql.DB
}

func (db *DBPool) Exec(query string, args ...any) (sql.Result, error) {
	return db.Master.Exec(query, args...)
}
func (db *DBPool) Query(query string, args ...any) (*sql.Rows, error) {
	return db.Slave.Query(query, args...)
}

func init() {
	var dsn = "root:123456@tcp(172.31.218.169:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"

	d, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error connecting to the database: %c\n", err)
	}
	db = d
	db.SetConnMaxIdleTime(0) // 设置连接可能处于空闲状态的最长时间, <= 0 连接不会因连接的空闲时间而关闭
	db.SetConnMaxLifetime(0) // 设置可以重复使用连接的最长时间, <= 0 连接不会因连接的超时而关闭
	db.SetMaxIdleConns(10)   // 设置空闲连接池中的最大连接数, <= 0 不保留空闲连接, 默认为 2
	db.SetMaxOpenConns(10)   // 设置数据库的最大打开连接数, <= 0 打开的连接数没有限制, 默认为 0

	// 读写分离配置
	// 创建两个分别连接到 mysql 主库和从库的客户端
}
func init() {
	// 读写分离配置
	// 使用中间件自动路由
	// pool := &DBPool{
	// 	Master: "",
	// 	Slave: "",
	// }
}

func init() {
	// 读写分离配置
	// 使用 gorm 框架配置
}

func printRows(rows *sql.Rows) {
	for rows.Next() {
		var id int
		var name string
		var age int
		cols, err := rows.Columns()
		if err != nil {
			log.Fatalf("Error rows.Columns() %s\n", err)
		}
		log.Println("columns", cols)
		colsType, err := rows.ColumnTypes()
		if err != nil {
			log.Fatalf("Error rows.ColumnTypes() %s\n", err)
		}
		for _, ty := range colsType {
			log.Println("columnTypes Name", ty.Name())
			log.Println("columnTypes ScanType", ty.ScanType().Name())
			log.Println("columnTypes DatebaseTypeName", ty.DatabaseTypeName())
		}

		if err := rows.Scan(&id, &name, &age); err != nil {
			log.Fatalf("Error scanning row: %v\n", err)
		}
		log.Printf("id: %d, name: %s, age: %d\n", id, name, age)
	}
}
