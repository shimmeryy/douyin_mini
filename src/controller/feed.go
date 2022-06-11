package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strconv"
	"tiktok/src/config"
	"tiktok/src/constants"
	"tiktok/src/errno"
	"tiktok/src/handlers"
	"tiktok/src/service"
	"time"
)

func Feed(c *gin.Context) {
	// latestTime
	var lastTime int64
	//token
	//claims := jwtUitl.ExtractClaims(c)
	//fmt.Printf("%+v", claims)

	tokenString := c.Query("token")
	var userID int64
	if token, err := config.AuthMiddleware.ParseTokenString(tokenString); err == nil {
		for key, value := range token.Claims.(jwt.MapClaims) {
			if key == constants.IdentityKey {
				userID = int64(value.(float64))
			}
		}
	}
	fmt.Printf("%+v", userID)
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
	//fmt.Printf("lastIndex:%d", lastIndex)
	//当视频全部播放完之后设置lastTime为当前时间 开始循环播放
	if lastIndex == -1 {
		lastTime = time.Now().Unix()
	} else {
		lastTime = videoList[lastIndex].CreatedAt.Unix()
	}

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
		commentCount, err := service.CommentServiceInstance().CountCommentByVideoId(c, handlers.CommentQueryByVideoIdParam{VideoId: videoId})
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
