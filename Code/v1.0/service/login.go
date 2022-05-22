package service

import (
	"Code/v1.0/database"
	"Code/v1.0/models"
	"Code/v1.0/tools"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	// Judge whether the username and password are empty
	if username == "" || password == "" {
		c.JSON(http.StatusOK, gin.H{
			"status_code": -1,
			"status_msg":  "Required information is NULL",
		})
		return
	}

	// Judge whether the username and password is it false
	data := new(models.User_table)
	err := database.DB.Where("username = ? AND password = ? ", username, password).First(&data).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusOK, gin.H{
				"status_code": -1,
				"status_msg":  "username or password is false",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status_code": -1,
			"status_msg":  "Get User_table Error:" + err.Error(),
		})
		return
	}

	// use tool class generate token
	token, err := tools.GenerateToken(data.User_id, username, password)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status_code": -1,
			"status_msg":  "GenerateToken Error:" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": map[string]interface{}{
			"status_code": 0,
			"status_msg":  "login success",
			"user_id":     data.User_id,
			"token":       token,
		},
	})

}
