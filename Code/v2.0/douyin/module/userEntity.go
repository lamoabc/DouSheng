package module

type User struct {
	Id            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
}
type UserTable struct {
	UserId        int64  `gorm:"column:user_id"`
	Username      string `gorm:"column:user_name"`
	Password      string `gorm:"column:account_password"`
	FollowCount   int64  `gorm:"follow_count"`
	FollowerCount int64  `gorm:"follower_count"`
}

func (u UserTable) TableName() string {
	// 绑定 Mysql 表名
	return "user_table"
}
