package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tiktok/src/constants"
	"tiktok/src/errno"
	"tiktok/src/handlers"
	"tiktok/src/service"
	"tiktok/src/utils/jwt"
)

// Publish 发布视频
func Publish(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	userID := int64(claims[constants.IdentityKey].(float64))
	title := c.PostForm("title")
	file, err := c.FormFile("data")
	if err != nil {
		panic(errno.ServiceErr.WithMessage("文件格式错误"))
	}
	err = service.VideoServiceInstance().CreateVideo(c, handlers.PublishParams{
		Data:   file,
		UserId: userID,
		Title:  title,
	})
	if err != nil {
		panic(errno.ServiceErr.WithMessage(err.Error()))
	}
	c.JSON(http.StatusOK, &handlers.UploadResponse{
		Response: handlers.Response{
			StatusCode: 200,
			StatusMsg:  "成功",
		},
	})
}

//TODO PublishList
//func PublishList(c gin.Context) {
//	token := c.Query("token")
//	userId := c.Query("user_id")
//
//}
