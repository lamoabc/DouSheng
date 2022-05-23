package dataImp

import (
	"douyin/dao"
	"douyin/module"
)

func QueryVideoByUserId(userId string, videoList []*module.VideoTable) (err error) {
	err = dao.Db.Where("author_id = ?", userId).Order("upload_date desc").Find(&videoList).Error
	return
}

func QueryAuthorByUserId(userId string, author *module.UserTable) (err error) {
	err = dao.Db.Where("id = ?", userId).First(&author).Error
	return
}

func IsFollow(followerId int64, followId string, follow *module.FollowTable) (isFollow bool, err error) {
	result := dao.Db.Where("follower_id = ? AND follow_id = ?", followerId, followId).First(&follow)
	if result.Error != nil {
		return false, result.Error
	}
	if result.RowsAffected > 0 {
		return true, nil
	} else {
		return false, nil
	}
}

func IsFavorite(followerId int64, videoList []*module.VideoTable, fav []*module.FavTable) (isFavList []bool, err error) {
	for i := 0; i < len(videoList); i++ {
		result := dao.Db.Where("fav_user_id = ? AND fav_video_id = ?", followerId, videoList[i].VideoId).First(&fav)
		if result.Error != nil {
			return nil, result.Error
		}
		if result.RowsAffected > 0 {
			isFavList[i] = true
		} else {
			isFavList[i] = false
		}
	}
	return isFavList, nil
}
