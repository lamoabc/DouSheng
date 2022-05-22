package controller

import (
	"douyin/module"
	"github.com/gin-gonic/gin"
)

type CommentListResponse struct {
	module.Response
	CommentList []module.Comment `json:"comment_list,omitempty"`
}

func CommentAction(c *gin.Context) {

}

func CommentList(c *gin.Context) {

}
