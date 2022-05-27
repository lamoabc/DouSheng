package relationService

import (
	"douyin/dao/dataImp/relationImp"
	"douyin/module"
	"douyin/module/jsonModule/response"
	"douyin/tools"
)

func FollowList(token string, userId int64, response *response.FollowList) {
	// 根据userId查询出该用户的关注列表
	var userList []module.UserTable
	if err := relationImp.QueryUserById(userId, &userList); err != nil {
		response.StatusCode = -1
		response.StatusMsg = err.Error()
		return
	}

	// 查询当前用户是否关注了列表中的用户
	isFolList := make([]bool, len(userList))
	if token != "" {
		userClaims, err := tools.AnalyseToken(token)
		if err != nil {
			response.StatusCode = -1
			response.StatusMsg = err.Error()
			return
		}
		isFolList, err = relationImp.IsFollow(userClaims.UserId, userList)
		if err != nil {
			response.StatusCode = -1
			response.StatusMsg = err.Error()
			return
		}
	}

	userListResp := make([]module.User, len(userList))
	for i := 0; i < len(userList); i++ {
		userListResp[i].Id = userList[i].UserId
		userListResp[i].Name = userList[i].Username
		userListResp[i].FollowCount = userList[i].FollowCount
		userListResp[i].FollowerCount = userList[i].FollowerCount
		userListResp[i].IsFollow = isFolList[i]
	}

	response.StatusCode = 0
	response.StatusMsg = "successful"
	response.UserList = userListResp
}
