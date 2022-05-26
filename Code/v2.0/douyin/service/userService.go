package service

import (
	"douyin/dao/dataImp"
	"douyin/module"
	"douyin/module/jsonModule/response"
	"douyin/tools"
    "douyin/dao/feedImp"
)

func login(u *module.User) {

}

func Register(username string, password string, response *response.Register) {
	//判断username是否重复
	err := dataImp.SelectUsername(&username, &password, new(module.UserTable))
	if err == nil {
		response.StatusCode = -1
		response.StatusMsg = "The username already exists"
		return
	}
	//创建用户
	u := module.UserTable{
		Username: username,
		Password: password,
	}
	err = dataImp.InsertUser(&u)
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
	var isFav []bool
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
	var isFol []bool
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
	response.NextTime = data[4].UploadDate
	response.List = VideoList
}