package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

const (
	dsn = `root:Help@110@tcp(10.10.10.100:3306)/code4fun?charset=utf8mb4&parseTime=True&loc=Local`
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func Test_MySQL(t *testing.T) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to open MySQL, err=%v", err)
	}
	// 迁移 Schema
	db.AutoMigrate(&Product{})
	// Create
	db.Create(&Product{Code: "D43", Price: 100})
	// Read
	var product Product
	db.First(&product, 1) // 根据主键查找
}

func Test_Connect(t *testing.T) {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // 数据源
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to open MySQL, err=%v", err)
	}
	// 迁移 Schema
	db.AutoMigrate(&Product{})
}

func Test_Create01(t *testing.T) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to open MySQL, err=%v", err)
	}
	// 迁移 Schema
	db.AutoMigrate(&Product{})
	db.Create(&Product{Code: "D44", Price: 101})
}

func Test_Query01(t *testing.T) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to open MySQL, err=%v", err)
	}
	// 迁移 Schema
	db.AutoMigrate(&Product{})
	var product Product
	db.First(&product, 2) // 根据主键查询
	fmt.Printf("product: %+v\n", product)
	product = Product{}
	db.First(&product, "code = ?", "D43") // 查找 code 字段为 D43 的记录
	fmt.Printf("product: %+v\n", product)
}
