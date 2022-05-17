package main

import (
	"Code/v1.0/database"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	r := gin.Default()

	initRouter(r)

	r.Run()

}
