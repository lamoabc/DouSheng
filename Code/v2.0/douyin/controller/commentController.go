package controller

import (
	"douyin/module"
	"douyin/module/jsonModule/response"
	"douyin/service/commentService"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CommentListResponse struct {
	module.Response
	CommentList []module.Comment `json:"comment_list,omitempty"`
}

func CommentAction(c *gin.Context) {
	//uid := c.Query("user_id")
	token := c.Query("token")
	vid := c.Query("video_id")
	actiontype := c.Query("action_type")

	var response response.CommentActionResponse
	//userId, err := strconv.ParseInt(uid, 10, 64)
	//if err != nil {
	//	response.StatusCode = -1
	//	response.StatusMsg = "userId Decoding failure"
	//	c.JSON(http.StatusOK, response)
	//	return
	//}
	videoId, _ := strconv.ParseInt(vid, 10, 64)

	if actiontype == "1" {
		comment_text := c.Query("comment_text")
		//commentService.AddComment(userId, token, videoId, comment_text, &response)
		commentService.AddComment(token, videoId, comment_text, &response)
		c.JSON(http.StatusOK, response)
	} else if actiontype == "2" {
		cid := c.Query("comment_id")
		commentId, _ := strconv.ParseInt(cid, 10, 64)
		commentService.DeleteComment(token, commentId, &response)
		c.JSON(http.StatusOK, response)
	}
}

func CommentList(c *gin.Context) {

}
