package database

import (
	"fmt"
	"gorm.io/gorm/schema"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Define DB in order to use Connect

var DB = Connect()

// Database Connect

func Connect() *gorm.DB {
	username := "root"   // username
	password := "382527" // password
	host := "127.0.0.1"  // database address
	port := 3306         // database port
	Dbname := "dousheng" // database name
	timeout := "10s"     // timeout

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)

	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
	})

	if err != nil {
		panic("could not connect to the database")
	}

	return connection
}
