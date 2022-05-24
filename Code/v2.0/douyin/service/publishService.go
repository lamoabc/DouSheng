package service

import (
	"douyin/dao/dataImp"
	"douyin/module"
	"douyin/module/jsonModule/response"
	"douyin/tools"
)

// PublishList all users have same publish video list
func PublishList(token string, userId string, response *response.PublishList) {
	// 查询该作者的所有video
	var videoList []*module.VideoTable
	if err := dataImp.QueryVideoByUserId(userId, &videoList); err != nil {
		response.StatusCode = -1
		response.StatusMsg = err.Error()
		return
	}

	// 查询作者信息
	author := new(module.UserTable)
	if err := dataImp.QueryAuthorByUserId(userId, author); err != nil {
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
	isFollow, err := dataImp.IsFollow(userClaims.UserId, userId, follow)
	if err != nil {
		response.StatusCode = -1
		response.StatusMsg = err.Error()
		return
	}

	// 查询该用户是否给video点赞
	var fav []*module.FavTable
	isFavList, err := dataImp.IsFavorite(userClaims.UserId, videoList, fav)
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
