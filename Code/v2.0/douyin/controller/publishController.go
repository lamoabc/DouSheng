package controller

import (
	"douyin/module"
	"github.com/gin-gonic/gin"
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

}
