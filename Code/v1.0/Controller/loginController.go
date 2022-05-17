package controller

import (
	"Code/v1.0/models"

	"github.com/gin-gonic/gin"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin

var userIdSequence = int64(1)

type UserLoginResponse struct {
	models.Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	models.Response
	User models.User `json:"user"`
}

func Login(c *gin.Context) {

	username := c.Query("username")
	password := c.Query("password")

	token := username + password

}

func UserInfo(c *gin.Context) {

}
