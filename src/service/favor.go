package service

import (
	"context"
	"errors"
	"sync"
	"tiktok/src/dal/db"
	"tiktok/src/errno"
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
		return errno.ServiceErr.WithMessage("invalid userId")
	}

	// check VideoId is valid or not
	if flag := VideoServiceInstance().CheckVideoById(ctx, req.VideoId); flag != true {
		return errno.ServiceErr.WithMessage("invalid videoId")
	}

	if req.ActionType != 1 && req.ActionType != 2 {
		return errno.ServiceErr.WithMessage("invalid actionType")
	}

	// req.ActionType == 2 取消点赞
	if req.ActionType == 2 {
		if err := db.DeleteFavorInfo(ctx, req.UserId, req.VideoId); err != nil {
			return err
		}
		return nil
	}

	// req.ActionType == 1 点赞
	flag, err := this.CheckIsFavored(ctx, handlers.FavorCheckParam{
		UserId:  req.UserId,
		VideoId: req.VideoId,
	})
	if err != nil {
		return err
	}
	// 如果已经点赞过了那么直接返回
	if flag {
		return nil
	}
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

// QueryFavorVideoByUserId 根据用户id获取点赞视频详情
func (this *FavorServiceImpl) QueryFavorVideoByUserId(ctx context.Context, req handlers.FavorQueryParam) ([]*handlers.VideoInfo, error) {
	list, err := this.QueryFavorByUserId(ctx, req)
	if err != nil {
		return nil, err
	}
	var res []*handlers.VideoInfo
	for _, favor := range list {
		videoId := favor.VideoId
		video, err := VideoServiceInstance().GetVideoById(ctx, videoId, req.UserId)
		if err != nil {
			return nil, err
		}
		res = append(res, video)
	}
	return res, nil
}

func (this *FavorServiceImpl) CountFavorByVideoId(ctx context.Context, req handlers.FavorCountParam) (int64, error) {
	if flag := VideoServiceInstance().CheckVideoById(ctx, req.VideoId); flag != true {
		return -1, errors.New("invalid videoId")
	}

	count, err := db.CountFavorByVideoId(ctx, req.VideoId)
	if err != nil {
		return -1, errno.ServiceErr.WithMessage("error occurs countFavorByVideoId")
	}
	return count, nil
}

func (this *FavorServiceImpl) CheckIsFavored(ctx context.Context, req handlers.FavorCheckParam) (bool, error) {
	if flag := UserServiceInstance().CheckUserById(ctx, req.UserId); flag != true {
		return false, errors.New("invalid userId")
	}

	if flag := VideoServiceInstance().CheckVideoById(ctx, req.VideoId); flag != true {
		return false, errno.ServiceErr.WithMessage("invalid videoId")
	}

	res, err := db.CheckIsFavored(ctx, req.UserId, req.VideoId)
	if err != nil {
		return false, err
	}
	return res, nil
}
