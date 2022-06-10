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
	"time"
)

func Feed(c *gin.Context) {
	// latestTime
	var lastTime int64
	//token
	claims := jwt.ExtractClaims(c)
	userID := int64(claims[constants.IdentityKey].(float64))
	strTime := c.Query("latest_time")
	if strTime == "" {
		lastTime = time.Now().Unix()
	} else {
		lastTime, _ = strconv.ParseInt(strTime, 10, 64)
	}
	videoList, err := service.FeedServiceInstance().Feed(c, lastTime)
	if err != nil {
		panic(err)
	}
	//取video列表的最后一个元素的下标
	lastIndex := len(videoList) - 1
	lastTime = videoList[lastIndex].CreatedAt.Unix()

	feedInfoList := make([]handlers.FeedInfo, len(videoList))

	for i := range videoList {
		videoId := int64(videoList[i].ID)
		authorInfo, err := service.UserServiceInstance().GetUserInfo(c, videoList[i].AuthorId)
		if err != nil {
			panic(errno.ServiceErr.WithMessage(err.Error()))
		}
		favored, err := service.FavorServiceInstance().CheckIsFavored(c, handlers.FavorCheckParam{
			UserId:  userID,
			VideoId: videoId,
		})
		if err != nil {
			panic(errno.ServiceErr.WithMessage(err.Error()))
		}
		commentCount, err := service.CommentServiceInstance().CountCommentByVideoId(c, handlers.CommentQueryParam{VideoId: videoId})
		if err != nil {
			panic(errno.ServiceErr.WithMessage(err.Error()))
		}
		countFavor, err := service.FavorServiceInstance().CountFavorByVideoId(c, handlers.FavorCountParam{VideoId: videoId})
		if err != nil {
			panic(errno.ServiceErr.WithMessage(err.Error()))
		}
		info := &handlers.FeedInfo{
			ID:            videoId,
			Author:        *authorInfo,
			PlayUrl:       videoList[i].PlayUrl,
			CoverUrl:      videoList[i].CoverUrl,
			FavoriteCount: countFavor,
			IsFavorite:    favored,
			CommentCount:  commentCount,
			Title:         videoList[i].Title,
		}
		feedInfoList[i] = *info
	}

	res := &handlers.FeedResponse{
		Response: handlers.Response{
			StatusCode: 0,
			StatusMsg:  "查询成功",
		},
		NextTime:  lastTime,
		VideoList: feedInfoList,
	}
	c.JSON(http.StatusOK, res)
}
