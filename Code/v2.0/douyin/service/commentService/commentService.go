package commentService

import (
	"douyin/dao"
	"douyin/dao/dataImp/commentImp"
	"douyin/module"
	"douyin/module/jsonModule/response"
	"douyin/tools"
	"time"
)

// add comment
func AddComment(token string, videoId int64, commentText string, response *response.CommentActionResponse) {
	tmp, err := tools.AnalyseToken(token)
	if err != nil {
		response.StatusCode = -1
		response.StatusMsg = "Encryption failed"
		return
	}
	usertable := new(module.UserTable)
	exist := commentImp.QueryUserId(tmp.UserId, usertable)
	if exist != nil {
		response.StatusCode = -1
		response.StatusMsg = " user is not login"
		return
	}
	createtime := time.Now().Format("2006-01-02")

	// insert CommentMsg
	//err = commentImp.InsertCommentMsg(videoId, userId, commentText, time)
	err = commentImp.InsertCommentMsg(videoId, tmp.UserId, commentText, createtime)
	if err != nil {
		response.StatusCode = -1
		response.StatusMsg = err.Error()
		return
	}

	CommentTable := new(module.CommentTable)
	// Query CommentMsg
	err = commentImp.QueryCommentMsgRes(videoId, tmp.UserId, createtime, CommentTable)
	if err != nil {
		response.StatusCode = -1
		response.StatusMsg = err.Error()
	}

	nusertable := new(module.UserTable)
	err = commentImp.QueryUserId(tmp.UserId, nusertable)
	if err != nil {
		response.StatusCode = -1
		response.StatusMsg = err.Error()
	}

	response.StatusCode = 0
	response.StatusMsg = "successful"
	response.CommentId = CommentTable.CommentId
	response.User.Id = usertable.UserId
	response.Name = nusertable.Username
	response.FollowCount = nusertable.FollowCount
	response.FollowerCount = nusertable.FollowerCount
	response.Content = CommentTable.Content
	response.CreateDate = createtime
	return
}

// delete comment
func DeleteComment(token string, commentId int64, response *response.CommentActionResponse) {
	_, err := tools.AnalyseToken(token)
	if err != nil {
		response.StatusCode = -1
		response.StatusMsg = "Encryption failed"
		return
	}
	CommentTable := new(module.CommentTable)
	// 目前的删除策略是根据评论id来进行删除(当前情况默认考虑评论唯一)
	exist := commentImp.DeleteCommentImp(commentId, CommentTable)
	if exist != nil {
		response.StatusCode = -1
		response.StatusMsg = exist.Error()
		return
	}
}

func ComList(videoId int64, response *response.CommentList) {
	var data []module.CommentTable
	err := commentImp.GetCommentList(videoId, &data)
	if err != nil {
		response.StatusCode = -1
		response.StatusMsg = err.Error()
		return
	}
	//装填response
	var commentTemp module.Comment
	for i := 0; i < len(data); i++ {
		commentTemp.Id = data[i].CommentId
		commentTemp.Content = data[i].Content
		commentTemp.CreateDate = data[i].CreateDate
		//查询评论对应的用户
		var user module.UserTable
		userId := data[i].ComUserId
		if err := dao.Db.Where("user_id = ?", userId).Find(&user).Error; err != nil {
			response.StatusCode = -1
			response.StatusMsg = "The username already exists"
			return
		}
		commentTemp.User.Id = user.UserId
		commentTemp.User.Name = user.Username
		commentTemp.User.FollowCount = user.FollowCount
		commentTemp.User.FollowerCount = user.FollowerCount
        commentTemp.User.Avatar = user.Avatar
		response.List = append(response.List, commentTemp)
	}
	response.StatusCode = 0
	response.StatusMsg = "successful"
}
