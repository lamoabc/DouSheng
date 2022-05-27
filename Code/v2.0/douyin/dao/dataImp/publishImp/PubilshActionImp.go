package publishImp

import (
	"douyin/dao"
	"douyin/module"
)

// Judge user is it exist
func QueryUserId(userId int64, usertable *module.UserTable) (err error) {
	err = dao.Db.Where("user_id = ?", userId).Find(&usertable).Error
	return
}

// insert data to video
func InsertData(userId int64, play_url string, cover_url string, title string) error {
	deres := dao.Db.Select("author_id", "play_url", "cover_url", "video_title").Create(module.VideoTable{AuthorId: userId, PlayUrl: play_url, CoverUrl: cover_url, VideoTitle: title})
	err := deres.Error
	return err
}
