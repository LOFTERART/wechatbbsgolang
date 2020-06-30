package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"os"
)

var DB *gorm.DB

// 初始化数据库连接等mysql
func Initialized() {
	// 申明变量
	var (
		err error
	)
	// 连接数据库
	DB, err = gorm.Open("mysql", os.Getenv("MYSQL_DSN"))
	if err != nil {
		log.Fatalf("database connection error: %v", err)
		return
	}

	DB.SingularTable(true)
	DB.AutoMigrate(&TbActivity{},
		&TbWelfare{},
		&TbAddress{},
		&TbBanner{},
		&Classify{},
		&Communitys{},
		&SubTopic{},
		&User{},
		&TbWxtoken{},
		&Ad{},
		&Diary{},
		&Comment{},
	)

	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(20000)
	// debug 模式开启sql日志
	DB.LogMode(true)
}

// 关闭数据库连接
func CloseDb() {
	DB.Close()
}
