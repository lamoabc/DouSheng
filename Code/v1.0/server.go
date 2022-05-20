package main

import (
	"Code/v1.0/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	router.Router(r)

	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
