package service

import (
	"sync"
)

type FeedService interface {
	CreateVideo()
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

//func (this *FeedServiceImpl) GetFeed(ctx context.Context, lastTime time.Time) *handlers.FeedResponse {
//
//	feedlist, err := db.QueryByTime(ctx, lastTime)
//
//	if err != nil {
//		return nil
//	}
//
//	feedlistResp := make([]handlers.FeedInfo, len(feedlist))
//
//	for i, _ := range feedlist {
//		user, err := db.QueryUserById(ctx, feedlist[i].AuthorId)
//		if err != nil {
//			return nil
//		}
//		video, err := db.QueryVideoById(ctx, int64(feedlist[i].ID))
//		if video != nil {
//			return nil
//		}
//
//		// 是否点赞
//		checkFavor := &handlers.FavorCheckParam{
//			VideoId: int64(video.ID),
//			UserId:  int64(user.ID),
//		}
//		checkCount := &handlers.FavorCountParam{
//			VideoId: int64(video.ID),
//		}
//
//		isfavor, err := FavorServiceInstance().CheckIsFavored(ctx, *checkFavor)
//		countfavor, err := FavorServiceInstance().CountFavorByVideoId(ctx, *checkCount)
//
//		// todo 构造参数返回
//		tmpUser := &handlers.UserInfo{
//			ID:            int64(user.ID),
//			UserName:      user.UserName,
//			FollowCount:   user.FollowCount,
//			FollowerCount: countfavor,
//			IsFollow:      isfavor,
//		}
//		feedlistResp[i] = handlers.FeedInfo{
//			ID:           int64(feedlist[i].ID),
//			AuthorId:     *tmpUser,
//			PlayUrl:      feedlist[i].PlayUrl,
//			CoverUrl:     feedlist[i].CoverUrl,
//			Title:        feedlist[i].Title,
//			CommentCount: feedlist[i].CommentCount,
//		}
//
//	}
//
//	// todo 返回resp格式的数据
//	return feedlistResp
//
//}
