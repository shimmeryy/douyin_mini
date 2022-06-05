package db

import (
	"context"
	"gorm.io/gorm"
	"tiktok/src/constants"
	"time"
)

type Feed struct {
	gorm.Model
	AuthorId     int64  `gorm:"column:author_id;type:bigint(20)not null;comment:作者id"`
	PlayUrl      string `gorm:"column:play_url;type:varchar(255)not null;comment:视频播放地址"`
	CoverUrl     string `gorm:"column:cover_url;type:varchar(255)not null;comment:视频封面地址"`
	CommentCount int64  `gorm:"column:comment_count;type:int(20);default:0;comment:评论数"`
	Title        string `gorm:"column:title;type:varchar(255)not null;comment:视频标题"`
}

// 返回按投稿时间倒序的视频列表,最多30个
func QueryByTime(ctx context.Context, last_time time.Time) ([]*Feed, error) {
	var VideoList []*Feed

	if err := DB.WithContext(ctx).Where("CreatedAt <= ?",
		last_time.Format("2022-06-03 22:54:05")).
		Preload("AuthorId").
		Order("UpdatedAt DESC").
		Limit(constants.MaxVideoNum).
		Find(&VideoList).Error; err != nil {
		return nil, err
	}
	return VideoList, nil
}
