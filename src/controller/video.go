package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tiktok/src/errno"
	"tiktok/src/handlers"
	"tiktok/src/service"
)

func Publish(c *gin.Context) {
	file, err := c.FormFile("data")
	if err != nil {
		panic(errno.ServiceErr.WithMessage("文件格式错误"))
	}
	url, err := service.VideoServiceInstance().CreateVideo(c, file)
	if err != nil {
		panic(errno.ServiceErr.WithMessage(err.Error()))
	}
	c.JSON(http.StatusOK, &handlers.OSSUploadResponse{
		Response: handlers.Response{
			StatusCode: 200,
			StatusMsg:  "成功",
		},
		Url: url,
	})
}
