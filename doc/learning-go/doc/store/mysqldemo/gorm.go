package mysqldemo

import (
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// gorm 将结构体映射到数据库表来简化数据库交互
/*
// gorm.Model 预定义结构体, 包含常用字段 ID, CreatedAt, UpdatedAt, DeletedAt
// ID 为默认主键
// 结构体名转换为 snake_case 并为表名加上复数形式
// 结构体字段名转换为 snake_case 作为表的列名
// gorm 使用字段 CreatedAt 和 UpdatedAt 自动跟踪记录的创建和更新时间
type Model struct {
	ID        uint           `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
*/
type Author struct {
	gorm.Model
	Name     string    `gorm:"<-:create;size:256" json:"name"` // 允许读和创建
	Email    string    `gorm:"<-:update;unique" json:"email"`  // 允许读和更新
	Age      uint8     `gorm:"<-;check:age<65" json:"age"`     // 允许读和写
	Count    uint      `gorm:"<-" json:"count,omitempty"`      // 允许读和写
	Total    uint      `gorm:"<-:false" json:"total"`          // 允许读, 禁止写
	Label    []string  `gorm:"->" json:"label"`                // 只读
	Birthday int64     `gorm:"serializer:unixtime;type:time"`
	F1       string    `gorm:"->;<-:create"`       // 允许读和写
	F2       time.Time `gorm:"->:false;<-:create"` // 仅创建
	F3       string    `gorm:"-"`                  // 通过 struct 读写会忽略该字段
	F4       string    `gorm:"-:all"`              // 通过 struct 读写、迁移会忽略该字段
	F5       string    `gorm:"-:migration"`        // 通过 struct 迁移会忽略该字段
}

// 自定义钩子
func (a *Author) BeforeCreate(tx *gorm.DB) error {
	return nil
}

type Blog struct {
	ID     uint
	Author Author `gorm:"embeded;embeddedPrefix:author_"` // 嵌入结构体并添加字段名前缀
}

var dateTimePrecision = 2
var gdb *gorm.DB

func init() {
	g, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       "root:123456@tcp(172.31.218.169:3306)/test?charset=utf8mb4&parseTime=True&loc=Local",
		DefaultStringSize:         256,                // add default size for string fields, by default, will use db type `longtext` for fields without size, not a primary key, no index defined and don't have default values
		DisableDatetimePrecision:  true,               // disable datetime precision support, which not supported before mysql 5.6
		DefaultDatetimePrecision:  &dateTimePrecision, // default datetime precision
		DontSupportRenameIndex:    true,               // drop & create index when rename index, rename index not supported before mysql 5.7, MariaDB
		DontSupportRenameColumn:   true,               // use change when rename column, rename column not supported before mysql 8, MariaDB
		SkipInitializeWithVersion: false,              // smart configure based on used version
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Error connecting to the database: %v\n", err)
	}
	gdb = g
}

func gormDemo() {
	log.SetPrefix("gorm ")
	var db, err = gdb.DB() // 维护连接池
	if err != nil {
		log.Fatalf("gdb.DB() Error: %v\n", err)
	}
	db.SetMaxOpenConns(10)
	log.Println("db.Stats().InUse", db.Stats().InUse)

	// // 迁移 schema
	// gdb.AutoMigrate(&Product{})

	// // Create
	// result := gdb.Create(&Product{Code: "D42", Price: 100})
	// reuslt.ID // 插入的主键 ID
	// reuslt.Error // 返回 error
	// Result.RowsAffected // 插入记录的条数
	// 批量创建传入一个切片
	// result := gdb.Create([]*Author{&Author{}, &Author{}})
	// result.Error // 返回 error
	// result.RowsAffected // 插入记录的条数

	// // Read
	// var author Author
	// gdb.First(&author, 1)                 // 根据整型主键查找第一条记录 SELECT * FROM products ORDER BY id LIMIT 1;
	// gdb.First(&author, "code = ?", "D42") // 查找 code 字段值为 D42 的记录
	// gdb.Last(&author) // 获取最后一条记录, 主键降序 SELECT * FROM authors ORDER BY id DESC LIMIT 1;
	// gdb.Take(&author, "name = ?", "zhangsan") // 返回符合条件的一条记录, 没有指定排序字段 SELECT * FROM authors WHERE name = 'zhangsan' LIMIT 1;
	// gdb.Find(&author, []int{1, 2, 3}) // SELECT * FROM authors WHERE id IN (1,2,3);

	// First 和 Last 只有在目标 struct 是指针或者通过 db.Model() 指定 model 时才有效, Take 方法没有此限制
	// // 读取第一条最后一条装入 result
	// var result = map[string]any{}
	// gdb.Model(&Author{}).First(&result) // works because model is specified using `db.Model()`
	// gdb.Model(&Author{}).Last(&result) // works because model is specified using `db.Model()`
	// gdb.Table("authors").First(&result) // doesn't work
	// gdb.Table("authors").Take(&result)  // works with Take

	// 条件 Where
	// gdb.Where("name = ?", "zhangsan").First(&Author{}) // 返回第一条 SELECT * FROM authors WHERE name = 'zhangsan' ORDER BY id LIMIT 1;
	// gdb.Where("name = ?", "zhangsan").Last(&Author{})  // 返回最后一条 SELECT * FROM authors WHERE name = 'zhangsan' ORDER BY id DESC LIMIT 1;
	// gdb.Where("name IN ?", []string{"zhangsan", "lisi"}).Find(&Author{}) // SELECT * FROM authors WHERE name IN ('zhangsan', 'lisi');
	// 条件 Not
	// gdb.Not("name = ?", "zhangsan").First(&Author{}) // SELECT * FROM authors WHERE NOT name = 'zhangsan' ORDER BY id LIMIT 1;
	// gdb.Not("name = ?", "zhangsan").Last(&Author{})  // SELECT * FROM authors WHERE NOT name = 'zhangsan' ORDER BY id DESC LIMIT 1;
	// gdb.Not(Author{Name: "zhangsan", Age: 18}).First(&Author{}) // SELECT * FROM authors WHERE name <> 'zhangsan' AND age <> 18 ORDER BY id LIMIT 1;
	// gdb.Not([]int64{1, 2, 3}).Find(&Author{}) // SELECT * FROM authors WHERE id NOT IN (1,2,3);
	// 条件 Or
	// gdb.Where("role = ?", "admin").Or("role = ?", "super_admin").Find(&Author{})          // SELECT * FROM authors WHERE role = 'admin' OR role = 'super_admin';
	// gdb.Where("role = ?", "admin").Or(Author{Name: "zhangsan", Age: 18}).Find(&Author{}) // SELECT * FROM authors WHERE role = 'admin' OR (name = 'zhangsan' AND age = 18);
	// 条件 排序
	// gdb.Order("age desc, name").Find(&Author{}) // SELECT * FROM authors ORDER BY age desc, name;
	// 条件 LIMIT OFFSET
	// gdb.Limit(3).Find(&Author{})            // SELECT * FROM authors LIMIT 3;
	// gdb.Limit(10).Offset(5).Find(&Author{}) // SELECT * FROM authors OFFSET 5 LIMIT 10;
	// 子查询
	// gdb.Where("amount > (?)", gdb.Select("AVG(amount)").Find(&Author{})).Find(&Author{}) // SELECT * FROM authors WHERE amount > (SELECT AVG(amount) FROM authors);

	// // Update - 将 product 的 price 更新为 200
	// gdb.Model(&product).Update("Price", 200)
	// // Update - 更新多个字段
	// gdb.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
	// gdb.Model(&product).Updates(map[string]any{"Price": 200, "Code": "F42"})
	// var author Author
	// gdb.First(&author)
	// author.Name = "zhangsan"
	// author.Age = 18
	// gdb.Save(&author)
	// gdb.Save(&Author{Name: "zhangsan", Age: 18})

	// // Delete - 删除 product
	// gdb.Delete(&product, 1)
	// gdb.Where("name = ?", "zhangsan").Delete(&Author{}) // DELETE FROM authors WHERE name = 'zhangsan';
}
