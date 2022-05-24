package service

import (
	"douyin/dao/dataImp"
	"douyin/module"
	"douyin/module/jsonModule/response"
	"douyin/tools"
)

func login(u *module.User) {

}

func Register(username string, password string, response *response.Register) {
	//判断username是否重复
	err := dataImp.SelectUsername(&username, &password, new(module.UserTable))
	if err == nil {
		response.StatusCode = -1
		response.StatusMsg = "The username already exists"
		return
	}
	//创建用户
	u := module.UserTable{
		Username: username,
		Password: password,
	}
	err = dataImp.InsertUser(&u)
	if err == nil {
		response.StatusCode = -1
		response.StatusMsg = "Registration fails"
	}
	token, err := tools.GenerateToken(u.UserId, username, password)

	response.Token = token
	response.UserId = u.UserId
	response.StatusCode = 0
	response.StatusMsg = "successful"
	return
}
