package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
	"time"
)

const (
	mysqlDSN = `root@tcp(localhost:3306)/db?charset=utf8mb4&parseTime=True&loc=Local`
)

var DB *gorm.DB

type User struct {
	ID        int       `gorm:"id" json:"id"`
	Sex       int       `gorm:"sex" json:"sex"`
	Username  string    `gorm:"username" json:"username"`
	Password  string    `gorm:"password" json:"password"`
	Email     string    `gorm:"email" json:"email"`
	CreatedAt time.Time `gorm:"created_at" json:"created_at"` // 默认自动跟踪记录创建时间
	UpdatedAt time.Time `gorm:"updated_at" json:"updated_at"` // 默认自动跟踪记录更新时间
}

func Init() {
	var err error
	DB, err = gorm.Open(mysql.Open(mysqlDSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func TestCreateUser(t *testing.T) {
	Init()
	user := &User{
		ID:       0,
		Sex:      1,
		Username: "张三",
		Password: "123456",
		Email:    "zhangsan@163.com",
	}
	if err := DB.Create(&user).Error; err != nil {
		t.Fatalf("err: %v", err)
	}
	t.Logf("user: %+v", user) // 成功创建记录后，会将主键ID，创建时间，更新时间回写到记录中
}

// 批量插入记录
func TestCreateUsers(t *testing.T) {
	Init()
	users := []*User{
		{
			ID:       0,
			Sex:      2,
			Username: "李四",
			Password: "woshilisi",
			Email:    "lisi@126.com",
		},
		{
			ID:       0,
			Sex:      1,
			Username: "王五",
			Password: "woshiwangwu",
			Email:    "wangwu@gamil.com",
		},
		{
			ID:       0,
			Sex:      1,
			Username: "赵六",
			Password: "woshizhaoliu",
			Email:    "zhaoliu@hotmail.com",
		},
	}
	if err := DB.Create(users).Error; err != nil {
		t.Fatalf("err: %v", err)
	}
	for _, user := range users {
		t.Logf("user: %+v", user)
	}
}

// 更新记录的单个列
func TestUpdateUser(t *testing.T) {
	Init()
	if err := DB.Model(&User{}).Where("id = ?", 1).Update("password", 654321).Error; err != nil {
		t.Errorf("err: %v", err)
	}
}
