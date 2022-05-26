package response

import "douyin/module"

type Feed struct {
	module.Response
	NextTime int64           `json:"next_time"`
	List     [5]module.Video `json:"video_list"`
}
