package controller

import (
	"douyin/module"
	"douyin/module/jsonModule/response"
	"douyin/service/publishService"
	"github.com/gin-gonic/gin"
	"net/http"
)

type VideoListResponse struct {
	module.Response
	VideoList []module.Video `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, module.Response{
			StatusCode: 1,
			StatusMsg:  "File receive error",
		})
		return
	}
	token := c.PostForm("token")
	title := c.PostForm("title")
	var response response.PublishAction
	publishService.PublishAction(data, token, title, &response)
	c.JSON(http.StatusOK, response)
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	token := c.Query("token")
	userId := c.Query("user_id")
	var response response.PublishList
	publishService.PublishList(token, userId, &response)
	c.JSON(http.StatusOK, response)
}
