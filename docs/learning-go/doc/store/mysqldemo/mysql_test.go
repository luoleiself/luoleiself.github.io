package mysqldemo

import (
	"database/sql"
	"testing"
)

// Ping
func TestMysqlPing(t *testing.T) {
	t.Run("Ping", func(t *testing.T) {
		if err := db.Ping(); err != nil {
			t.Fatalf("Error pinging the database: %v\n", err)
		}
		t.Log("Successfully connected to the database.")
		// 已注册的驱动名称
		t.Log("sql.Drivers()", sql.Drivers())
		// 驱动
		t.Log("db.Driver()", db.Driver())
	})
}

// Stats
func TestMySqlStat(t *testing.T) {
	t.Run("Stats", func(t *testing.T) {
		// 数据库的状态
		stat := db.Stats()
		t.Log("空闲连接数 stat.Idle", stat.Idle)
		t.Log("当前正在使用的连接数 stat.InUse", stat.InUse)
		t.Log("建立连接数 stat.OpenConnections", stat.OpenConnections)
		t.Log("数据库打开连接的最大数量 stat.MaxOpenConnections", stat.MaxOpenConnections)
		t.Log("连接等待数 stat.WaitCount", stat.WaitCount)
		t.Log("新连接等待的时长 stat.WaitDuration", stat.WaitDuration)
		t.Log("因为 SetMaxIdleConns 而关闭的总数 stat.MaxIdleClosed", stat.MaxIdleClosed)
		t.Log("因为 SetConnMaxIdleTime 而关闭的总数 stat.MaxIdleTimeClosed", stat.MaxIdleTimeClosed)
		t.Log("因为 SetConnMaxLifetime 而关闭的总数 stat.MaxLifetimeClosed", stat.MaxLifetimeClosed)
	})
}

// Insert
func TestMysqlInsert(t *testing.T) {
	t.Run("Exec(InsertCmd)", func(t *testing.T) {
		// 插入语句
		rows, err := db.Query("SELECT * FROM user WHERE name = ?", "Tom")
		if err != nil {
			t.Fatalf("Error Query by condition: %v\n", err)
		}
		if rows.Next() {
			t.Log("name = Tom is exists.")
		} else {
			t.Log("INSERT INTO...")
			result, err := db.Exec("INSERT INTO user (name) VALUES (?)", "Tom")
			if err != nil {
				t.Fatalf("Error Exec insert: %v\n", err)
			}
			lastInsertId, _ := result.LastInsertId()
			t.Logf("inserted record with id: %d \n", lastInsertId)
			affected, _ := result.RowsAffected()
			t.Logf("affected record with len: %d\n", affected)
		}
	})
}

// Query
func TestMysqlQuery(t *testing.T) {
	t.Run("Query", func(t *testing.T) {
		t.Log("db.Query...")
		rows, err := db.Query("SELECT * FROM user")
		if err != nil {
			t.Fatalf("error querying data: %v\n", err)
		}
		defer rows.Close()
		printRows(rows)
	})
}

// Update
func TestMysqlUpdate(t *testing.T) {
	t.Run("Exec(UpdateCmd)", func(t *testing.T) {
		t.Log("UPDATE SET...")
		result, err := db.Exec("UPDATE user SET age = ? WHERE name = ?", 18, "Tom")
		if err != nil {
			t.Fatalf("Error Updating data: %v\n", err)
		}
		lastInsertId, _ := result.LastInsertId()
		affected, _ := result.RowsAffected()
		t.Log("result.LastInsertId", lastInsertId)
		t.Log("result.RowsAffected", affected)

		t.Log("db.QueryRow...")
		row := db.QueryRow("SELECT * FROM user WHERE name = ? and age = ?", "Tom", 18)
		if row.Err() == nil {
			var id, age int
			var name string
			row.Scan(&id, &name, &age)
			t.Logf("id: %d, name: %s, age: %d\n", id, name, age)
		}
	})
}

// Delete
func TestMysqlDelete(t *testing.T) {
	t.Run("", func(t *testing.T) {
		t.Log("DELETE...")
		result, err := db.Exec("DELETE FROM user WHERE age < ?", 18)
		if err != nil {
			t.Fatalf("Error Delete data: %v\n", err)
		}
		lastInsertId, _ := result.LastInsertId()
		affected, _ := result.RowsAffected()
		t.Log("result.LastInsertId", lastInsertId)
		t.Log("result.RowsAffected", affected)
		rows, err := db.Query("SELECT * FROM user")
		if err != nil {
			t.Fatalf("Error Select * : %s\n", err)
		}
		defer rows.Close()
		printRows(rows)
	})
}

// Prepare
func TestMysqlPrepare(t *testing.T) {
	t.Run("Prepare", func(t *testing.T) {
		t.Log("db.Prepare...")
		stmt, err := db.Prepare("INSERT INTO user(name, age) VALUES(?, ?)")
		if err != nil {
			t.Fatalf("Error db.Prepare: %v\n", err)
		}
		defer stmt.Close()
		result, err := stmt.Exec("Jerry", 3)
		if err != nil {
			t.Fatalf("Error stmt.Exec %v\n", err)
		}
		lastInsertId, _ := result.LastInsertId()
		affected, _ := result.RowsAffected()
		t.Logf("result.LastInsertId %d result.RowsAffected %d\n", lastInsertId, affected)

		// prepare queryRow
		t.Log("db.prepare... stmt.QueryRow")
		stmt, err = db.Prepare("SELECT * FROM user WHERE name = ?")
		if err != nil {
			t.Fatalf("Error db.Prepare: %v\n", err)
		}
		defer stmt.Close()
		row := stmt.QueryRow("Jerry")
		if row.Err() == nil {
			var id, age int
			var name string
			row.Scan(&id, &name, &age)
			t.Logf("id: %d, name: %s, age: %d\n", id, name, age)
		}
	})
}
