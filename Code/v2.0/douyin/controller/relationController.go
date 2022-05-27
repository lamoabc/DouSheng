package controller

import (
	"douyin/module"
	"douyin/module/jsonModule/response"
	"douyin/service/relationService"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserListResponse struct {
	module.Response
	UserList []module.User `json:"user_list"`
}

func RelationAction(c *gin.Context) {

}

func FollowList(c *gin.Context) {
	token := c.Query("token")
	id := c.Query("user_id")
	userId, _ := strconv.ParseInt(id, 10, 64)
	var response response.FollowList
	relationService.FollowList(token, userId, &response)
	c.JSON(http.StatusOK, response)
}

func FollowerList(c *gin.Context) {
	token := c.Query("token")
	id := c.Query("user_id")
	userId, _ := strconv.ParseInt(id, 10, 64)
	var response response.FollowList
	relationService.FollowerList(token, userId, &response)
	c.JSON(http.StatusOK, response)
}
