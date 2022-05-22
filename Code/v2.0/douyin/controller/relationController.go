package controller

import (
	"douyin/module"
	"github.com/gin-gonic/gin"
)

type UserListResponse struct {
	module.Response
	UserList []module.User `json:"user_list"`
}

func RelationAction(c *gin.Context) {

}

func FollowList(c *gin.Context) {

}

func FollowerList(c *gin.Context) {

}
