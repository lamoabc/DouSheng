package module

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}
type FollowTable struct {
	FollowId   int64 `gorm:"column:follow_id"`
	FollowerId int64 `gorm:"column:follower_id"`
}
type FavTable struct {
	FavUserId  int64 `gorm:"column:fav_user_id"`
	FavVideoId int64 `gorm:"column:fav_video_id"`
}
