package visitorService

import (
	"douyin/dao/dataImp/loginImp"
	"douyin/dao/feedImp"
	"douyin/module"
	"douyin/module/jsonModule/response"
	"douyin/tools"
)

func Login(username string, password string, response *response.Login) {

	data := new(module.UserTable)
	message := loginImp.Login(&username, &password, data)
	if message != "" {
		response.StatusCode = -1
		response.StatusMsg = message
		return
	}
	// use tool class generate token
	token, err := tools.GenerateToken(data.UserId, username, password)

	if err != nil {
		response.StatusCode = -1
		response.StatusMsg = "Encryption failed"
		return
	}
	//一切无误,装填信息
	response.Token = token
	response.UserId = data.UserId
	response.StatusCode = 0
	response.StatusMsg = "successful"
	return
}
func Feed(latestTime int64, response *response.Feed) {
	//游客登录状态
	//声明Table去数据库拿值,装填进response
	var data []module.VideoWithAuthor
	var message string
	if latestTime > 0 {
		//限制时间戳
		message = feedImp.Feed2(latestTime, &data)
	} else {
		//没有限制时间戳
		message = feedImp.Feed1(&data)
	}
	if message != "" {
		//有异常,装填response
		response.StatusCode = -1
		response.StatusMsg = message
		return
	} else {
		var VideoList = [5]module.Video{}
		for i := 0; i < 5; i++ {
			VideoList[i].Id = data[i].VideoId
			VideoList[i].Author.Id = data[i].UserId
			VideoList[i].Author.Name = data[i].Username
			VideoList[i].Author.IsFollow = false
			VideoList[i].Author.FollowCount = data[i].FollowCount
			VideoList[i].Author.FollowerCount = data[i].FollowerCount
			VideoList[i].CommentCount = data[i].ComCount
			VideoList[i].FavoriteCount = data[i].FavCount
			VideoList[i].CoverUrl = data[i].CoverUrl
			VideoList[i].IsFavorite = false
			VideoList[i].PlayUrl = data[i].PlayUrl
			VideoList[i].VideoTitle = data[i].VideoTitle
		}
		response.StatusCode = 0
		response.StatusMsg = "successful"
		response.NextTime = data[4].UploadDate
		response.List = VideoList
		return
	}
}
