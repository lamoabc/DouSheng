package main

import (
	"Code/v1.0/controller"

	"github.com/gin-gonic/gin"
)

func initRouter(r *gin.Engine) {
	// public directory is used to serve static resources

	apiRouter := r.Group("/douyin")

	// basic apis

	apiRouter.POST("/user/login/", controller.Login)

}
