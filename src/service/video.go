package service

import (
	"context"
	"io/ioutil"
	"sync"
	"tiktok/src/dal/db"
	"tiktok/src/errno"
	"tiktok/src/handlers"
	"tiktok/src/utils/ossUtil"
)

type VideoService interface {
	CreateVideo()
	GetVideosByAuthor()
	GetVideoById()
	CheckVideoById()
}

var (
	videoService     *VideoServiceImpl
	videoServiceOnce sync.Once
)

func VideoServiceInstance() *VideoServiceImpl {
	videoServiceOnce.Do(
		func() {
			videoService = &VideoServiceImpl{}
		})
	return videoService
}

type VideoServiceImpl struct {
}

func (this *VideoServiceImpl) CreateVideo(ctx context.Context, req handlers.PublishParams) (err error) {
	userId := req.UserId
	title := req.Title
	file := req.Data
	fileHandle, err := file.Open()
	if err != nil {
		return errno.ServiceErr.WithMessage("打开文件错误")
	}
	defer fileHandle.Close()
	fileByte, err := ioutil.ReadAll(fileHandle)
	url, err := ossUtil.UploadFile(file.Filename, fileByte)
	if err != nil {
		return errno.ServiceErr.WithMessage("上传失败")
	}

	err = db.CreateVideo(ctx, db.Video{
		AuthorId: userId,
		PlayUrl:  url,
		CoverUrl: url + "?x-oss-process=video/snapshot,t_500,f_jpg,w_600,h_800,m_fast",
		Title:    title,
	})
	if err != nil {
		return errno.ServiceErr.WithMessage("发布失败")
	}
	return nil
}

func (this *VideoServiceImpl) GetVideosByAuthor(ctx context.Context, userId int64) (list []*db.Video, err error) {
	list, err = db.QueryVideoByAuthor(ctx, int64(userId))
	if err != nil {
		return nil, errno.ServiceErr.WithMessage("查询失败")
	}
	return list, nil
}

// GetVideoById 根据id获取Video
func (this *VideoServiceImpl) GetVideoById(ctx context.Context, videoId int64, userId int64) (videoInfo *handlers.VideoInfo, err error) {
	video, err := db.QueryVideoById(ctx, videoId)
	if err != nil {
		return nil, err
	}
	if video != nil && video.ID == 0 {
		return nil, errno.ServiceErr.WithMessage("视频不存在")
	}
	user, err := db.QueryUserById(ctx, video.AuthorId)
	if err != nil {
		return nil, errno.ServiceErr.WithMessage("用户不存在")
	}

	commentCount, err := CommentServiceInstance().CountCommentByVideoId(ctx, handlers.CommentQueryByVideoIdParam{VideoId: videoId})
	if err != nil {
		panic(errno.ServiceErr.WithMessage(err.Error()))
	}

	favorCount, err := FavorServiceInstance().CountFavorByVideoId(ctx, handlers.FavorCountParam{VideoId: videoId})
	if err != nil {
		panic(errno.ServiceErr.WithMessage(err.Error()))
	}

	isFavorite, err := FavorServiceInstance().CheckIsFavored(ctx, handlers.FavorCheckParam{
		UserId:  userId,
		VideoId: videoId,
	})
	if err != nil {
		panic(errno.ServiceErr.WithMessage(err.Error()))
	}

	videoInfo = &handlers.VideoInfo{
		ID: int64(video.ID),
		Author: handlers.UserInfo{
			ID:            int64(user.ID),
			UserName:      user.UserName,
			FollowCount:   user.FollowCount,
			FollowerCount: user.FollowerCount,
			IsFollow:      false,
		},
		PlayUrl:       video.PlayUrl,
		CoverUrl:      video.CoverUrl,
		Title:         video.Title,
		FavoriteCount: favorCount,
		CommentCount:  commentCount,
		IsFavorite:    isFavorite,
	}

	return videoInfo, nil
}

func (this *VideoServiceImpl) CheckVideoById(ctx context.Context, id int64) bool {
	video, err := db.QueryVideoById(ctx, id)
	if err != nil {
		return false
	}
	if video != nil && video.ID == 0 {
		return false
	}
	return true
}
