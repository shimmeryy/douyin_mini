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

// FavoriteAction 点赞或取消点赞
func FavoriteAction(c *gin.Context) {
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

	err = service.FavorServiceInstance().OperateFavor(c, handlers.FavorOperateParam{
		UserId:     userID,
		VideoId:    videoId,
		ActionType: int32(actionType),
	})
	if err != nil {
		panic(errno.ServiceErr.WithMessage(err.Error()))
	}

	c.JSON(http.StatusOK, handlers.Response{
		StatusCode: 0,
		StatusMsg:  "success",
	})
}

// FavoriteList 获取点赞列表
func FavoriteList(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		panic(errno.ParamErr.WithMessage("Error UserId"))
	}

	videoList, err := service.FavorServiceInstance().QueryFavorVideoByUserId(c, handlers.FavorQueryParam{UserId: userId})
	if err != nil {
		panic(errno.ServiceErr.WithMessage(err.Error()))
	}

	c.JSON(http.StatusOK, handlers.FavorQueryResponse{
		Response: handlers.Response{
			StatusCode: 0,
			StatusMsg:  "success",
		},
		VideoList: videoList,
	})

}
