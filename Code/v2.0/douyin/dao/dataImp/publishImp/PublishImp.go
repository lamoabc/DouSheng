package publishImp

import (
	"douyin/dao"
	"douyin/module"
	"gorm.io/gorm"
)

func QueryVideoListByUserId(userId int64, videoList *[]module.VideoWithAuthor) (err error) {
	err = dao.Db.Where("author_id = ?", userId).Order("upload_date desc").Find(&videoList).Error
	return
}

func IsFollow(followerId int64, videoList []module.VideoWithAuthor) (isFollowList []bool, err error) {
	var follow module.FollowTable
	for i := 0; i < len(videoList); i++ {
		err = dao.Db.Where("follower_id = ? AND follow_id = ?", followerId, videoList[i].AuthorId).First(&follow).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				isFollowList = append(isFollowList, false)
				err = nil
			} else {
				return nil, err
			}
		} else {
			isFollowList = append(isFollowList, true)
		}
	}
	return
}

func IsFavorite(followerId int64, videoList []module.VideoWithAuthor) (isFavList []bool, err error) {
	var fav module.FavTable
	for i := 0; i < len(videoList); i++ {
		err = dao.Db.Where("fav_user_id = ? AND fav_video_id = ?", followerId, videoList[i].VideoId).First(&fav).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				isFavList = append(isFavList, false)
				err = nil
			} else {
				return nil, err
			}
		} else {
			isFavList = append(isFavList, true)
		}
	}
	return
}
