package router

import (
	"Code/v1.0/service"
	"github.com/gin-gonic/gin"
)

//Define routing rules

func Router(r *gin.Engine) {

	apiRouter := r.Group("/douyin")

	// basic apis
	apiRouter.GET("/user/login", service.Login)
}
