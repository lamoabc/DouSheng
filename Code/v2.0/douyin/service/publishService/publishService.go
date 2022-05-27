package publishService

import (
	"douyin/dao/dataImp/publishImp"
	"douyin/module"
	"douyin/module/jsonModule/response"
	"douyin/tools"
	"mime/multipart"
)

func PublishAction(data *multipart.FileHeader, token string, title string, response *response.PublishAction) {
	//根据token解析出来的内容判断内容是否存在
	tmp, err := tools.AnalyseToken(token)
	if err != nil {
		response.StatusCode = -1
		response.StatusMsg = "Encryption failed"
		return
	}
	usertale := new(module.UserTable)
	exist := publishImp.QueryUserId(tmp.UserId, usertale)
	if exist != nil {
		response.StatusCode = -1
		response.StatusMsg = " no is it exist"
		return
	}
	fileContent, _ := data.Open()

	play_url := tools.GetPlayUrl(data.Filename, fileContent)
	cover_url := tools.GetCoverUrl(data.Filename)
	err = publishImp.InsertData(tmp.UserId, play_url, cover_url, title)
	if err != nil {
		response.StatusCode = -1
		response.StatusMsg = err.Error()
		return
	}
	// 成功的返回值
	response.StatusCode = 0
	response.StatusMsg = "successful"
	return
}

// PublishList all users have same publish video list
func PublishList(token string, userId string, response *response.PublishList) {
	// 查询该作者的所有video
	var videoList []*module.VideoTable
	if err := publishImp.QueryVideoByUserId(userId, &videoList); err != nil {
		response.StatusCode = -1
		response.StatusMsg = err.Error()
		return
	}

	// 查询作者信息
	author := new(module.UserTable)
	if err := publishImp.QueryAuthorByUserId(userId, author); err != nil {
		response.StatusCode = -1
		response.StatusMsg = err.Error()
		return
	}

	// 查询用户是否关注了此作者
	// 先把token里面的userId解析出来，不同于作者的userId
	userClaims, err := tools.AnalyseToken(token)
	if err != nil {
		response.StatusCode = -1
		response.StatusMsg = err.Error()
		return
	}
	// 在根据用户id和作者id查询是否关注
	var follow *module.FollowTable
	isFollow, err := publishImp.IsFollow(userClaims.UserId, userId, follow)
	if err != nil {
		response.StatusCode = -1
		response.StatusMsg = err.Error()
		return
	}

	// 查询该用户是否给video点赞
	var fav []*module.FavTable
	isFavList, err := publishImp.IsFavorite(userClaims.UserId, videoList, fav)
	if err != nil {
		response.StatusCode = -1
		response.StatusMsg = err.Error()
		return
	}

	// 该作者和其所有的 video 成功查询后，填装response
	authorResp := new(module.User)
	authorResp.Id = author.UserId
	authorResp.Name = author.Username
	authorResp.FollowCount = author.FollowCount
	authorResp.FollowerCount = author.FollowerCount
	authorResp.IsFollow = isFollow

	videoListResp := make([]module.Video, len(videoList))
	for i := 0; i < len(videoList); i++ {
		videoListResp[i].Id = videoList[i].VideoId
		videoListResp[i].Author = *authorResp
		videoListResp[i].PlayUrl = videoList[i].PlayUrl
		videoListResp[i].CoverUrl = videoList[i].CoverUrl
		videoListResp[i].FavoriteCount = videoList[i].FavCount
		videoListResp[i].CommentCount = videoList[i].ComCount
		videoListResp[i].IsFavorite = isFavList[i]
		videoListResp[i].VideoTitle = videoList[i].VideoTitle
	}

	response.StatusCode = 0
	response.StatusMsg = "success"
	response.VideoList = videoListResp
	return
}
