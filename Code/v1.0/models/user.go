package models

type User_table struct {
	User_id        int64  `json:"user_id,omitempty"`
	Password       string `json:"password,omitempty"`
	Username       string `json:"username,omitempty"`
	Followcount    int    `json:"follow_count"`
	Follower_count int    `json:"follower_count"`
	//StatusCode int32  `json:"status_code"`
	//StatusMsg  string `json:"status_msg,omitempty"`
}
