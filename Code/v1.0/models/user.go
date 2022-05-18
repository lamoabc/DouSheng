package models

type User struct {
	Id   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	// FollowCount   int64  `json:"follow_count,omitempty"`
	// FollowerCount int64  `json:"follower_count,omitempty"`
	// IsFollow      bool   `json:"is_follow,omitempty"`
	UserId     int64  `json:"user_id,omitempty"`
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
	Token      string `json:"token"`
}
