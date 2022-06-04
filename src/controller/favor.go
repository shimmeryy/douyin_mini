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

func Action(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	userID := int64(claims[constants.IdentityKey].(float64))

	videoId, err := strconv.ParseInt(c.PostForm("video_id"), 10, 64)
	if err != nil {
		panic(errno.ParamErr.WithMessage("Error VideoId"))
	}

	actionType, err := strconv.ParseInt(c.PostForm("action_type"), 10, 32)
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

//func List(c *gin.Context) {
//	userId := c.Query("user_id")
//
//}
