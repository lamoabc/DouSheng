package service

import (
	"Code/v1.0/database"
	"Code/v1.0/models"
	"Code/v1.0/tools"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserLoginResponse struct {
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	// Judge whether the username and password are empty

	if username == "" || password == "" {
		c.JSON(http.StatusOK, models.User{
			StatusCode: -1,
			StatusMsg:  "Required information is NULL",
		})
		return
	}

	data := new(models.User)
	// use tool class generate token
	token, err := tools.GenerateToken(data.Id, data.Name)

	if err != nil {
		c.JSON(http.StatusOK, models.User{
			StatusCode: -1,
			StatusMsg:  "GenerateToken Error:" + err.Error(),
		})
		return
	}

	user := models.User{}
	exist := database.DB.Model(&user).Where("username = ? and password = ?", username, password)
	var userid int64
	database.DB.Model(&user).Where("username", username).Find("use_id", userid)

	if exist != nil {
		c.JSON(http.StatusOK, models.User{
			StatusCode: 0,
			StatusMsg:  "login success",
			UserId:     userid,
			Token:      token,
		})
	} else {
		c.JSON(http.StatusOK, models.User{
			StatusCode: 1,
			StatusMsg:  "User doesn't exist",
		})
	}

}
