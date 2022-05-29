package api

import (
	"Java2GO/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type VideoListResponse struct {
	models.Response
	VideoList []models.Video `json:"video_list"`
}

func Publish(c *gin.Context) {
	c.JSON(http.StatusOK, models.Response{
		StatusCode: 0,
		StatusMsg:  "ok",
	})
}
