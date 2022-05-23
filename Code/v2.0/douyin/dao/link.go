package dao

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var Db = Connect()

func Connect() *gorm.DB {
	username := "root"  // username
	password := ""      // password
	host := "127.0.0.1" // database address
	port := 3306        // database port
	Dbname := "douyin"  // database name
	timeout := "10s"    // timeout

	// 通过前面的数据库参数，拼接 Mysql DSN，其实就是数据库连接串（数据源名称）
	// Mysql dsn格式： {username}:{password}@tcp({host}:{port})/{Dbname}?charset=utf8&parseTime=True&loc=Local
	// 类似{username}使用花括号包着的名字都是需要替换的参数
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)

	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		panic("could not connect to the database")
	}

	return connection
}
