package service

import (
	"context"
	"sync"
	"tiktok/src/dal/db"
	"tiktok/src/errno"
)

type FeedService interface {
	Feed()
}

var (
	feedService     *FeedServiceImpl
	feedServiceOnce sync.Once
)

func FeedServiceInstance() *FeedServiceImpl {
	feedServiceOnce.Do(
		func() {
			feedService = &FeedServiceImpl{}
		})
	return feedService

}

type FeedServiceImpl struct {
}

func (this *FeedServiceImpl) Feed(ctx context.Context, lastTime int64) ([]*db.Video, error) {

	videoList, err := db.QueryByTime(ctx, lastTime)
	if err != nil {
		return nil, errno.ServiceErr.WithMessage("根据时间倒序查询失败")
	}

	return videoList, nil
}
