package service

import (
	"douyin/dao/dataImp"
	"douyin/module"
	"douyin/module/jsonModule/response"
	"douyin/tools"
)

func Login(username string, password string, response *response.Login) {

	data := new(module.UserTable)
	message := dataImp.Login(&username, &password, data)
	if message != "" {
		response.StatusCode = -1
		response.StatusMsg = message
		return
	}
	// use tool class generate token
	token, err := tools.GenerateToken(data.UserId, username, password)

	if err != nil {
		response.StatusCode = -1
		response.StatusMsg = "Encryption failed"
		return
	}
	//一切无误,装填信息
	response.Token = token
	response.UserId = data.UserId
	response.StatusCode = 0
	response.StatusMsg = "successful"
	return
}
