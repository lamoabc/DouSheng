package module

type User struct {
	Id            int64  `json:"id"`
	Name          string `json:"name"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}
type UserTable struct {
	UserId          int64  `gorm:"column:user_id;primaryKey"`
	Username        string `gorm:"column:user_name"`
	Password        string `gorm:"column:account_password"`
	FollowCount     int64  `gorm:"follow_count"`
	FollowerCount   int64  `gorm:"follower_count"`
	Signature       string `gorm:"signature"`
	Avatar          string `gorm:"avatar"`
	BackgroundImage string `gorm:"background_image"`
}

func (u UserTable) TableName() string {
	// 绑定 Mysql 表名
	return "user_table"
}
