package db

import (
	"context"
	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	AuthorId int64  `gorm:"column:author_id;type:bigint(20)not null;comment:作者id"`
	PlayUrl  string `gorm:"column:play_url;type:varchar(255)not null;comment:视频播放地址"`
	CoverUrl string `gorm:"column:cover_url;type:varchar(255)not null;comment:视频封面地址"`
	Title    string `gorm:"column:title;type:varchar(255)not null;comment:视频标题"`
}

func (v Video) TableName() string {
	return "video"
}

// QueryVideoByAuthor 根据作者ID查询视频集
func QueryVideoByAuthor(ctx context.Context, authorId int64) ([]*Video, error) {
	var res []*Video
	if err := DB.WithContext(ctx).Where("author_id = ?", authorId).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// CreateVideo 插入视频
func CreateVideo(ctx context.Context, video Video) error {
	if err := DB.WithContext(ctx).Create(&video).Error; err != nil {
		return err
	}
	return nil
}

// QueryVideoById 根据ID查询视频
func QueryVideoById(ctx context.Context, videoId int64) (*Video, error) {
	var res *Video
	if err := DB.WithContext(ctx).Where("id = ?", videoId).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
