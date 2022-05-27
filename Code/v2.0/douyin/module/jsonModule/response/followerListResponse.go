package response

import "douyin/module"

type FollowerList struct {
	module.Response
	UserList []module.User `json:"user_list"`
}
