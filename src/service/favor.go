package service

import (
	"context"
	"errors"
	"sync"
	"tiktok/src/dal/db"
	"tiktok/src/handlers"
)

type FavorService interface {
}

type FavorServiceImpl struct {
}

var (
	favorService     *FavorServiceImpl
	favorServiceOnce sync.Once
)

func FavorServiceInstance() *FavorServiceImpl {
	favorServiceOnce.Do(
		func() {
			favorService = &FavorServiceImpl{}
		})
	return favorService
}

// OperateFavor 点赞或取消点赞
func (this *FavorServiceImpl) OperateFavor(ctx context.Context, req handlers.FavorOperateParam) error {
	// check UserId is valid or not
	if flag := UserServiceInstance().CheckUserById(ctx, req.UserId); flag != true {
		return errors.New("invalid userId")
	}

	// check VideoId is valid or not
	if flag := VideoServiceInstance().CheckVideoById(ctx, req.VideoId); flag != true {
		return errors.New("invalid videoId")
	}

	// req.ActionType == 2 取消点赞
	if req.ActionType == 2 {
		if err := db.DeleteFavorInfo(ctx, req.UserId, req.VideoId); err != nil {
			return err
		}
		return nil
	}

	// req.ActionType == 1 点赞
	favor := db.Favor{
		UserId:  req.UserId,
		VideoId: req.VideoId,
	}
	if err := db.CreateFavorInfo(ctx, favor); err != nil {
		return err
	}
	return nil
}

// QueryFavorByUserId 根据用户id获取点赞列表
func (this *FavorServiceImpl) QueryFavorByUserId(ctx context.Context, req handlers.FavorQueryParam) ([]*db.Favor, error) {
	var res []*db.Favor
	res, err := db.QueryFavorByUserId(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// TODO 根据用户id获取点赞视频详情
// QueryFavorVideoByUserId 根据用户id获取点赞视频详情
//func (this *FavorServiceImpl) QueryFavorVideoByUserId(ctx context.Context, req handlers.FavorQueryParam) ([]*db.Video, error) {
//	tmp, err := this.QueryFavorByUserId(ctx, req)
//	if err != nil {
//		return nil, err
//	}
//	var res []*db.Video
//	for _, favor := range tmp {
//		videoId := favor.VideoId
//		video, err := VideoServiceInstance().GetVideoById(ctx, videoId)
//		if err != nil {
//			return nil, err
//		}
//		res = append(res, &video)
//	}
//
//}

func (this *FavorServiceImpl) CountFavorByVideoId(ctx context.Context, req handlers.FavorCountParam) (int64, error) {
	if flag := VideoServiceInstance().CheckVideoById(ctx, req.VideoId); flag != true {
		return -1, errors.New("invalid videoId")
	}

	count, err := db.CountFavorByVideoId(ctx, req.VideoId)
	if err != nil {
		return -1, errors.New("error occurs countFavorByVideoId")
	}
	return count, nil
}

func (this *FavorServiceImpl) CheckIsFavored(ctx context.Context, req handlers.FavorCheckParam) (bool, error) {
	if flag := UserServiceInstance().CheckUserById(ctx, req.UserId); flag != true {
		return false, errors.New("invalid userId")
	}

	if flag := VideoServiceInstance().CheckVideoById(ctx, req.VideoId); flag != true {
		return false, errors.New("invalid videoId")
	}

	res, err := db.CheckIsFavored(ctx, req.UserId, req.VideoId)
	if err != nil {
		return false, err
	}
	return res, nil
}
