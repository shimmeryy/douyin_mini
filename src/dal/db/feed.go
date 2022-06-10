package db

import (
	"context"
	"tiktok/src/constants"
	"time"
)

//QueryByTime 返回按投稿时间倒序的视频列表,最多30个
func QueryByTime(ctx context.Context, lastTime int64) ([]*Video, error) {
	var VideoList []*Video

	if err := DB.WithContext(ctx).Where("created_at <= ?",
		time.Unix(lastTime, 0).Format("2006-01-02 15-04-05")).
		//Preload("author_id").
		Order("created_at DESC").
		Limit(constants.MaxVideoNum).
		Find(&VideoList).Error; err != nil {
		return nil, err
	}
	return VideoList, nil
}
