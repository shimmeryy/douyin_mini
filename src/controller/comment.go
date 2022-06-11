package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tiktok/src/constants"
	"tiktok/src/errno"
	"tiktok/src/handlers"
	"tiktok/src/service"
	"tiktok/src/utils/jwt"
)

func CommentAction(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	userID := int64(claims[constants.IdentityKey].(float64))

	videoId, err := strconv.ParseInt(c.Query("video_id"), 10, 64)
	if err != nil {
		panic(errno.ParamErr.WithMessage("Error VideoId"))
	}

	actionType, err := strconv.ParseInt(c.Query("action_type"), 10, 32)
	if err != nil {
		panic(errno.ParamErr.WithMessage("Error ActionType"))
	}

	// 评论
	if actionType == 1 {
		commentText := c.Query("comment_text")
		commentInfo, err := service.CommentServiceInstance().OperateComment(c, handlers.CommentOperateParam{
			UserId:     userID,
			VideoId:    videoId,
			ActionType: int32(actionType),
			Text:       commentText,
		})
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, handlers.CommentCreateResponse{
			Response: handlers.Response{
				StatusCode: 0,
				StatusMsg:  "success",
			},
			CommentInfo: *commentInfo,
		})
	} else if actionType == 2 { // 取消评论
		commentId, err := strconv.ParseInt(c.Query("comment_id"), 10, 64)
		if err != nil {
			panic(errno.ParamErr.WithMessage("Error CommentId"))
		}
		if _, err := service.CommentServiceInstance().OperateComment(c, handlers.CommentOperateParam{
			UserId:     userID,
			VideoId:    videoId,
			ActionType: int32(actionType),
			CommentId:  commentId,
		}); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, handlers.Response{
			StatusCode: 0,
			StatusMsg:  "success",
		})
	} else {
		panic(errno.ParamErr.WithMessage("Error ActionType"))
	}
}

func CommentList(c *gin.Context) {
	videoId, err := strconv.ParseInt(c.Query("video_id"), 10, 64)
	if err != nil {
		panic(errno.ParamErr.WithMessage("Error VideoId"))
	}

	commentList, err := service.CommentServiceInstance().QueryCommentByVideoId(c, handlers.CommentQueryByVideoIdParam{
		VideoId: videoId,
	})
	if err != nil {
		panic(errno.ServiceErr.WithMessage(err.Error()))
	}

	c.JSON(http.StatusOK, handlers.CommentQueryResponse{
		Response: handlers.Response{
			StatusCode: 0,
			StatusMsg:  "success",
		},
		CommentList: commentList,
	})
}
