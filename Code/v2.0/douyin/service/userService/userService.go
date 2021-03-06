package userService

import (
	"douyin/dao"
	"douyin/dao/dataImp/userInfoImp"
	"douyin/dao/favListImp"
	"douyin/dao/favouriteImp"
	"douyin/dao/feedImp"
	"douyin/dao/followImp"
	"douyin/module"
	"douyin/module/jsonModule/response"
	"douyin/tools"
)

func Register(username string, password string, response *response.Register) {
	//创建用户
	u := module.UserTable{
		Username:        username,
		Password:        tools.GetMd5(password),
		Signature:       "欢迎使用抖声APP",
		Avatar:          "https://yygh-lamo.oss-cn-beijing.aliyuncs.com/User%20Avatar/3.jpeg",
		BackgroundImage: "https://yygh-lamo.oss-cn-beijing.aliyuncs.com/User%20background/defaultBackGround.png",
	}
	if err := dao.Db.Create(&u).Error; err != nil {
		response.StatusCode = -1
		response.StatusMsg = "The username already exists"
		return
	}
	token, _ := tools.GenerateToken(u.UserId, username, password)

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
	if message != "" {
		//拿data过程有异常
		response.StatusCode = -1
		response.StatusMsg = message
		return
	}
	//data无误拿到
	if len(data) < 1 {
		response.StatusCode = -1
		response.StatusMsg = "没有更多视频了"
		return
	}
	//根据userId查用户对data里的视频是否喜欢
	var flag bool
	var isFav []bool
	for i := 0; i < len(data); i++ {
		flag, message = feedImp.Feed3(user.UserId, data[i].VideoId)
		if message != "" {
			break
		}
		isFav = append(isFav, flag)
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
		flag, message = feedImp.Feed4(data[i].AuthorId, user.UserId)
		if message != "" {
			break
		}
		isFol = append(isFol, flag)
	}
	if message != "" {
		//查是否喜欢过程中有异常
		response.StatusCode = -1
		response.StatusMsg = message
		return
	}
	//data,isFav,isFol无误拿到,装填response
	var videoTemp module.Video
	for i := 0; i < len(data); i++ {
		videoTemp.Id = data[i].VideoId
		videoTemp.Author.Id = data[i].UserId
		videoTemp.Author.Name = data[i].Username
		videoTemp.Author.IsFollow = isFol[i]
		videoTemp.Author.FollowCount = data[i].FollowCount
		videoTemp.Author.FollowerCount = data[i].FollowerCount
		videoTemp.CommentCount = data[i].ComCount
		videoTemp.FavoriteCount = data[i].FavCount
		videoTemp.CoverUrl = data[i].CoverUrl
		videoTemp.IsFavorite = isFav[i]
		videoTemp.PlayUrl = data[i].PlayUrl
		videoTemp.VideoTitle = data[i].VideoTitle
		videoTemp.Author.Signature = data[i].Signature
		videoTemp.Author.BackgroundImage = data[i].BackgroundImage
		videoTemp.Author.Avatar = data[i].Avatar
		response.List = append(response.List, videoTemp)
	}
	response.StatusCode = 0
	response.StatusMsg = "successful"
}
func FavList(userId int64, token string, response *response.FavouriteList) {
	//解析token拿visitorId
	user, err := tools.AnalyseToken(token)
	if err != nil {
		response.StatusCode = -1
		response.StatusMsg = "Token Encryption failed"
		return
	}
	//token无误
	//声明点赞列表和数据库对接的module,去数据库拿值
	var data []module.UserLikeVideoList
	message := favListImp.GetVideoList(userId, &data)
	if message != "" {
		//拿data过程有异常
		response.StatusCode = -1
		response.StatusMsg = message
		return
	}
	//data无误拿到
	if len(data) < 1 {
		response.StatusCode = -1
		response.StatusMsg = "还没有点赞过视频"
		return
	}
	//根据visitorId查用户对data里的视频是否喜欢
	var flag bool
	var isFav []bool
	for i := 0; i < len(data); i++ {
		flag, message = favListImp.IsFav(user.UserId, data[i].VideoId)
		if message != "" {
			break
		}
		isFav = append(isFav, flag)
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
		flag, message = favListImp.IsFollow(data[i].AuthorId, user.UserId)
		if message != "" {
			break
		}
		isFol = append(isFol, flag)
	}
	if message != "" {
		//查是否喜欢过程中有异常
		response.StatusCode = -1
		response.StatusMsg = message
		return
	}
	//data,isFav,isFol无误拿到,装填response
	var videoTemp module.Video
	for i := 0; i < len(data); i++ {
		videoTemp.Id = data[i].VideoId
		videoTemp.Author.Id = data[i].UserId
		videoTemp.Author.Name = data[i].Username
		videoTemp.Author.IsFollow = isFol[i]
		videoTemp.Author.FollowCount = data[i].FollowCount
		videoTemp.Author.FollowerCount = data[i].FollowerCount
		videoTemp.CommentCount = data[i].ComCount
		videoTemp.FavoriteCount = data[i].FavCount
		videoTemp.CoverUrl = data[i].CoverUrl
		videoTemp.IsFavorite = isFav[i]
		videoTemp.PlayUrl = data[i].PlayUrl
		videoTemp.VideoTitle = data[i].VideoTitle
		videoTemp.Author.Signature = data[i].Signature
		videoTemp.Author.BackgroundImage = data[i].BackgroundImage
		videoTemp.Author.Avatar = data[i].Avatar
		response.List = append(response.List, videoTemp)
	}
	response.StatusCode = 0
	response.StatusMsg = "successful"
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

func UserInfo(token string, userId string, response *response.UserInfo) {
	//查询用户信息
	// 先把token里面的userId解析出来，不同于作者的userId
	userClaims, err := tools.AnalyseToken(token)
	if err != nil {
		response.StatusCode = -1
		response.StatusMsg = err.Error()
		return
	}
	// 在根据用户id和要查询的userid查询是否关注（也就是用户id是否关注了要查询的userid）
	var follow *module.FollowTable
	isFollow, _ := userInfoImp.IsFollow(userClaims.UserId, userId, follow)

	userTable := new(module.UserTable)
	err = userInfoImp.SelectAuthorByUserId(userId, userTable)
	if err != nil {
		response.StatusCode = -1
		response.StatusMsg = "The user information for the user id does not exist"
		return
	}
	response.StatusCode = 0
	response.Id = userTable.UserId
	response.Name = userTable.Username
	response.IsFollow = isFollow
	response.FollowCount = userTable.FollowCount
	response.FollowerCount = userTable.FollowerCount
	response.Signature = userTable.Signature
	response.BackgroundImage = userTable.BackgroundImage
	response.Avatar = userTable.Avatar
	response.StatusMsg = "successful"
	return
}
