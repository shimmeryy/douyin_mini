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
func (this *FavorServiceImpl) OperateFavor(ctx context.Context, req handlers.FavorUpdateParam) error {

	if flag := UserServiceInstance().CheckUserById(ctx, req.UserId); flag != true {
		return errors.New("invalid userId")
	}

	// TODO Check VideoId 的合法性

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
