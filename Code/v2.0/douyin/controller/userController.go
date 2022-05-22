package controller

import (
	"douyin/module/jsonModule/response"
	"douyin/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(c *gin.Context) {

}

func Login(c *gin.Context) {
	//拿request里給的信息
	username := c.Query("username")
	password := c.Query("password")
	//声明此次请求需要返回的response
	var response response.Login
	// Judge whether the username and password are
	if username == "" || password == "" {
		//request里的信息为空,无需调用登录服务,直接装填response
		response.StatusCode = -1
		response.StatusMsg = "Required information is NULL"
		c.JSON(http.StatusOK, response)
		return
	} else {
		service.Login(username, password, &response)
		c.JSON(http.StatusOK, response)
	}
}

func UserInfo(c *gin.Context) {
	var response response.UserInfo
	response.StatusCode = 0
	response.StatusMsg = "successful"
	response.Id = 1
	response.Name = "zhiqieryi"
	response.FollowCount = 0
	response.FollowerCount = 0
	response.IsFollow = true
	c.JSON(http.StatusOK, response)
}
