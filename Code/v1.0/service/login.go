package service

import (
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "pong",
	})

	//username := c.Query("username")
	//password := c.Query("password")
	//
	//token := username + password

}

func UserInfo(c *gin.Context) {

}
