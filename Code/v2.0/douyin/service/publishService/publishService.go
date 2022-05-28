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
		response.StatusMsg = " user is not login"
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
func PublishList(token string, userId int64, response *response.PublishList) {
	var videoList []module.VideoWithAuthor
	if err := publishImp.QueryVideoListByUserId(userId, &videoList); err != nil {
		response.StatusCode = -1
		response.StatusMsg = err.Error()
		return
	}

	// 查询用户是否关注了此作者
	// 先把token里面的userId解析出来，不同于作者的userId
	isFollowList := make([]bool, len(videoList))
	isFavList := make([]bool, len(videoList))
	if token != "" {
		userClaims, err := tools.AnalyseToken(token)
		if err != nil {
			response.StatusCode = -1
			response.StatusMsg = err.Error()
			return
		}
		// 在根据用户id和videoList中的作者id查询是否关注
		isFollowList, err = publishImp.IsFollow(userClaims.UserId, videoList)
		if err != nil {
			response.StatusCode = -1
			response.StatusMsg = err.Error()
			return
		}

		// 查询该用户是否给video点赞
		isFavList, err = publishImp.IsFavorite(userClaims.UserId, videoList)
		if err != nil {
			response.StatusCode = -1
			response.StatusMsg = err.Error()
			return
		}
	}

	// 该作者和其所有的 video 成功查询后，填装response
	videoListResp := make([]module.Video, len(videoList))
	for i := 0; i < len(videoList); i++ {
		videoListResp[i].Id = videoList[i].VideoId
		videoListResp[i].Author.Id = videoList[i].AuthorId
		videoListResp[i].Author.Name = videoList[i].Username
		videoListResp[i].Author.FollowCount = videoList[i].FollowCount
		videoListResp[i].Author.FollowerCount = videoList[i].FollowerCount
		videoListResp[i].Author.IsFollow = isFollowList[i]
		videoListResp[i].PlayUrl = videoList[i].PlayUrl
		videoListResp[i].CoverUrl = videoList[i].CoverUrl
		videoListResp[i].FavoriteCount = videoList[i].FavCount
		videoListResp[i].CommentCount = videoList[i].ComCount
		videoListResp[i].IsFavorite = isFavList[i]
		videoListResp[i].VideoTitle = videoList[i].VideoTitle
	}

	response.StatusCode = 0
	response.StatusMsg = "successful"
	response.VideoList = videoListResp
	return
}
