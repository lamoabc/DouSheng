package database

import (
	"Code/v1.0/models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	username := "root"   // 账号
	password := "382527" // 密码
	host := "127.0.0.1"  // 数据库地址，可以是Ip或者域名
	port := 3306         // 数据库端口
	Dbname := "user"     // 数据库名
	timeout := "10s"     // 连接超时，10秒

	// 拼接下 dsn 参数
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)

	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("could not connect to the database")
	}

	DB = connection

	connection.AutoMigrate(&models.User{})
}
