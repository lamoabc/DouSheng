package userService

import (
	"douyin/dao/dataImp/registerImp"
    "douyin/dao/favouriteImp"
    "douyin/dao/followImp"
	"douyin/dao/feedImp"
	"douyin/module"
	"douyin/module/jsonModule/response"
	"douyin/tools"
)

func Register(username string, password string, response *response.Register) {
	//判断username是否重复
	err := registerImp.SelectUsername(&username, &password, new(module.UserTable))
	if err != nil {
		response.StatusCode = -1
		response.StatusMsg = "The username already exists"
		return
	}
	//创建用户
	u := module.UserTable{
		Username: username,
		Password: password,
	}
	err = registerImp.InsertUser(&u)
	if err == nil {
		response.StatusCode = -1
		response.StatusMsg = "Registration fails"
	}
	token, err := tools.GenerateToken(u.UserId, username, password)

	response.Token = token
	response.UserId = u.UserId
	response.StatusCode = 0
	response.StatusMsg = "successful"
	return
}
func Feed(latestTime int64, token string, response *response.Feed) {
	//用户登录状态
	//解析token拿到用户信息
	user, err := tools.AnalyseToken(token)
	if err != nil {
		response.StatusCode = -1
		response.StatusMsg = "Token Encryption failed"
		return
	}
	//token无误
	//声明视频流需要到数据库里拿的module,去数据库拿值
	var data []module.VideoWithAuthor
	var message string
	if latestTime > 0 {
		//限制时间戳
		message = feedImp.Feed2(latestTime, &data)
	} else {
		//没有限制时间戳
		message = feedImp.Feed1(&data)
	}
	if message != "" || len(data) < 1 {
		//拿data过程有异常
		response.StatusCode = -1
		response.StatusMsg = message
		return
	}
	//data无误拿到
	//根据userid查用户对data里的视频是否喜欢
	var isFav [5]bool
	for i := 0; i < len(data); i++ {
		isFav[i], message = feedImp.Feed3(user.UserId, data[i].VideoId)
		if message != "" {
			break
		}
	}
	if message != "" {
		//查是否喜欢过程中有异常
		response.StatusCode = -1
		response.StatusMsg = message
		return
	}
	//根据userid查用户对data里视频的作者是否关注
	var isFol [5]bool
	for i := 0; i < len(data); i++ {
		isFol[i], message = feedImp.Feed4(data[i].AuthorId, user.UserId)
		if message != "" {
			break
		}
	}
	if message != "" {
		//查是否喜欢过程中有异常
		response.StatusCode = -1
		response.StatusMsg = message
		return
	}
	//data,isFav,isFol无误拿到,装填response
	var VideoList = [5]module.Video{}
	for i := 0; i < len(data); i++ {
		VideoList[i].Id = data[i].VideoId
		VideoList[i].Author.Id = data[i].UserId
		VideoList[i].Author.Name = data[i].Username
		VideoList[i].Author.IsFollow = isFol[i]
		VideoList[i].Author.FollowCount = data[i].FollowCount
		VideoList[i].Author.FollowerCount = data[i].FollowerCount
		VideoList[i].CommentCount = data[i].ComCount
		VideoList[i].FavoriteCount = data[i].FavCount
		VideoList[i].CoverUrl = data[i].CoverUrl
		VideoList[i].IsFavorite = isFav[i]
		VideoList[i].PlayUrl = data[i].PlayUrl
		VideoList[i].VideoTitle = data[i].VideoTitle
	}
	response.StatusCode = 0
	response.StatusMsg = "successful"
	response.NextTime = data[len(data)-1].UploadDate
	response.List = VideoList
}
func UserFav(userId int64, videoId int64, actionType int64, response *response.Favourite) {
	//根据actionType进行点赞服务或者取消点赞服务
	if actionType == 1 {
		//点赞
		//将点赞记录同步更新到数据库
		mes := favouriteImp.Insert(userId, videoId)
		if mes != "" {
			response.StatusCode = -1
			response.StatusMsg = mes
			return
		}
		response.StatusCode = 0
		response.StatusMsg = "点赞成功"
		return
	}
	if actionType == 2 {
		//取消点赞
		//将点赞记录从数据库里同步删除
		mes := favouriteImp.Delete(userId, videoId)
		if mes != "" {
			response.StatusCode = -1
			response.StatusMsg = mes
			return
		}
		response.StatusCode = 0
		response.StatusMsg = "取消点赞成功"
		return
	}
	//actionType意外的值错误
	response.StatusCode = -1
	response.StatusMsg = "ActionType value is invalid"
	return
}
func UserFol(followId int64, followerId int64, actionType int64, response *response.Follow) {
	//根据actionType提供关注服务或者取消关注服务
	if actionType == 1 {
		//关注
		//将关注记录同步更新到数据库
		mes := followImp.Insert(followId, followerId)
		if mes != "" {
			response.StatusCode = -1
			response.StatusMsg = mes
			return
		}
		response.StatusCode = 0
		response.StatusMsg = "关注成功"
		return
	}
	if actionType == 2 {
		//取消关注
		//将关注记录从数据库里同步删除
		mes := followImp.Delete(followId, followerId)
		if mes != "" {
			response.StatusCode = -1
			response.StatusMsg = mes
			return
		}
		response.StatusCode = 0
		response.StatusMsg = "取消关注成功"
		return
	}
	//actionType意外的值错误
	response.StatusCode = -1
	response.StatusMsg = "ActionType value is invalid"
	return
}