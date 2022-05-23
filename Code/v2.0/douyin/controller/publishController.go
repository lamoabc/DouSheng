package controller

import (
	"douyin/module"
	"douyin/module/jsonModule/response"
	"douyin/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type VideoListResponse struct {
	module.Response
	VideoList []module.Video `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	token := c.Query("token")
	userId := c.Query("user_id")
	var response response.PublishList
	service.PublishList(token, userId, &response)
	c.JSON(http.StatusOK, response)
}
