package models

type User struct {
	Username    string `json:"username" gorm:"unique"`
	Password    string `json:"password"`
	Status_code uint   `json:"status_code"`
	User_id     uint   `json:"user_id"`
	Token       string `json:"token"`
}
