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

// Publish 发布视频
func Publish(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	userID := int64(claims[constants.IdentityKey].(float64))
	title := c.PostForm("title")
	file, err := c.FormFile("data")
	if err != nil {
		panic(errno.ParamErr)
	}
	err = service.VideoServiceInstance().CreateVideo(c, handlers.PublishParams{
		Data:   file,
		UserId: userID,
		Title:  title,
	})
	if err != nil {
		panic(errno.ServiceErr.WithMessage("创建失败"))
	}
	c.JSON(http.StatusOK, &handlers.UploadResponse{
		Response: handlers.Response{
			StatusCode: 200,
			StatusMsg:  "成功",
		},
	})
}

// PublishList 发布列表
func PublishList(c *gin.Context) {
	//claims := jwt.ExtractClaims(c)
	//userID := int64(claims[constants.IdentityKey].(float64))
	userID := c.Query("user_id")
	ID, err2 := strconv.Atoi(userID)
	if err2 != nil {
		panic(errno.ServiceErr.WithMessage("参数异常"))
	}

	list, err := service.VideoServiceInstance().GetVideosByAuthor(c, int64(ID))
	if err != nil {
		panic(errno.ServiceErr.WithMessage(err.Error()))
	}

	publishList := make([]handlers.VideoInfo, len(list))
	for i := range list {

		videoId := int64(list[i].ID)
		authorInfo, err := service.UserServiceInstance().GetUserInfo(c, list[i].AuthorId)
		if err != nil {
			panic(errno.ServiceErr.WithMessage(err.Error()))
		}
		favored, err := service.FavorServiceInstance().CheckIsFavored(c, handlers.FavorCheckParam{
			UserId:  int64(ID),
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
		info := &handlers.VideoInfo{
			ID:            videoId,
			Author:        *authorInfo,
			PlayUrl:       list[i].PlayUrl,
			CoverUrl:      list[i].CoverUrl,
			FavoriteCount: countFavor,
			IsFavorite:    favored,
			CommentCount:  commentCount,
			Title:         list[i].Title,
		}
		publishList[i] = *info
	}
	//fmt.Printf("%+v\n", publishList)
	res := &handlers.PublishListResponse{
		Response: handlers.Response{
			StatusCode: 0,
			StatusMsg:  "获取成功",
		},
		VideoList: publishList,
	}
	c.JSON(http.StatusOK, res)
}
