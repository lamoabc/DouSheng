package commentService

import (
	"douyin/dao/dataImp/commentImp"
	"douyin/module"
	"douyin/module/jsonModule/response"
	"douyin/tools"
	"time"
)

// add comment
func AddComment(userId int64, token string, videoId int64, commentText string, response *response.CommentActionResponse) {
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
	time := time.Now().String()

	// insert CommentMsg
	err = commentImp.InsertCommentMsg(videoId, userId, commentText, time)
	if err != nil {
		response.StatusCode = -1
		response.StatusMsg = err.Error()
		return
	}

	CommentTable := new(module.CommentTable)
	// Query CommentMsg
	err = commentImp.QueryCommentMsgRes(videoId, userId, time, CommentTable)
	if err != nil {
		response.StatusCode = -1
		response.StatusMsg = err.Error()
	}

	nusertable := new(module.UserTable)
	err = commentImp.QueryUserId(userId, nusertable)
	if err != nil {
		response.StatusCode = -1
		response.StatusMsg = err.Error()
	}

	response.StatusCode = 0
	response.StatusMsg = "successful"
	response.CommentId = CommentTable.CommentId
	response.User.Id = nusertable.UserId
	response.Name = nusertable.Username
	response.FollowCount = nusertable.FollowCount
	response.FollowerCount = nusertable.FollowerCount
	response.Content = CommentTable.Content
	response.CreateDate = time
	return
}

// delete comment
func DeleteComment(userId int64, token string, commentId int64, response *response.CommentActionResponse) {
	tmp, err := tools.AnalyseToken(token)
	if err != nil {
		response.StatusCode = -1
		response.StatusMsg = "Encryption failed"
		return
	}
	if userId != tmp.UserId {
		response.StatusCode = -1
		response.StatusMsg = "this is not your content"
	} else {
		CommentTable := new(module.CommentTable)
		exist := commentImp.DeleteCommentImp(commentId, CommentTable)
		if exist != nil {
			response.StatusCode = -1
			response.StatusMsg = exist.Error()
			return
		}

	}

}
