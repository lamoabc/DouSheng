package dataImp

import (
	"douyin/dao"
	"douyin/module"
)

func SelectUsername(username *string, password *string, data *module.UserTable) (err error) {
	err = dao.Db.Where("user_name = ?", username).Find(&data).Error
	return
}

func InsertUser(u *module.UserTable) (err error) {
	err = dao.Db.Create(&u).Error
	return
}
