package visitorService

import (
	"douyin/dao/dataImp/loginImp"
	"douyin/dao/feedImp"
	"douyin/module"
	"douyin/module/jsonModule/response"
	"douyin/tools"
    "douyin/dao/favListImp"
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
		var videoTemp module.Video
		for i := 0; i < len(data); i++ {
			videoTemp.Id = data[i].VideoId
			videoTemp.Author.Id = data[i].UserId
			videoTemp.Author.Name = data[i].Username
			videoTemp.Author.IsFollow = false
			videoTemp.Author.FollowCount = data[i].FollowCount
			videoTemp.Author.FollowerCount = data[i].FollowerCount
			videoTemp.CommentCount = data[i].ComCount
			videoTemp.FavoriteCount = data[i].FavCount
			videoTemp.CoverUrl = data[i].CoverUrl
			videoTemp.IsFavorite = false
			videoTemp.PlayUrl = data[i].PlayUrl
			videoTemp.VideoTitle = data[i].VideoTitle
			videoTemp.Author.Signature = data[i].Signature
			videoTemp.Author.BackGround = data[i].BackGround
			videoTemp.Author.Avatar = data[i].Avatar
			response.List = append(response.List, videoTemp)
		}
		response.StatusCode = 0
		response.StatusMsg = "successful"
	}
}
func FavList(userId int64, response *response.FavouriteList) {
	//游客登录状态
	//声明点赞列表和数据库对接的module,去数据库拿值
	var data []module.UserLikeVideoList
	message := favListImp.GetVideoList(userId, &data)
	if message != "" {
		//拿data过程有异常
		response.StatusCode = -1
		response.StatusMsg = message
		return
	}
	//data无误拿到,装填response
	var videoTemp module.Video
	for i := 0; i < len(data); i++ {
		videoTemp.Id = data[i].VideoId
		videoTemp.Author.Id = data[i].UserId
		videoTemp.Author.Name = data[i].Username
		videoTemp.Author.IsFollow = false
		videoTemp.Author.FollowCount = data[i].FollowCount
		videoTemp.Author.FollowerCount = data[i].FollowerCount
		videoTemp.CommentCount = data[i].ComCount
		videoTemp.FavoriteCount = data[i].FavCount
		videoTemp.CoverUrl = data[i].CoverUrl
		videoTemp.IsFavorite = false
		videoTemp.PlayUrl = data[i].PlayUrl
		videoTemp.VideoTitle = data[i].VideoTitle
		videoTemp.Author.Signature = data[i].Signature
		videoTemp.Author.BackGround = data[i].BackGround
		videoTemp.Author.Avatar = data[i].Avatar
		response.List = append(response.List, videoTemp)
	}
	response.StatusCode = 0
	response.StatusMsg = "successful"
}