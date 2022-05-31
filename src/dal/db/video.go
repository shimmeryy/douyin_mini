package db

import (
	"context"
	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	AuthorId     int64  `gorm:"column:author_id;type:bigint(20,0)not null;comment:作者id"`
	PlayUrl      string `gorm:"column:play_url;type:varchar(255)not null;comment:视频播放地址"`
	CoverUrl     string `gorm:"column:cover_url;type:varchar(255)not null;comment:视频封面地址"`
	CommentCount int64  `gorm:"column:comment_count;type:int(20)default:0;comment:评论数"`
	Title        string `grom:"column:title;type:varchar(255)not null;comment:视频标题"`
}

func (v Video) TableName() string {
	return "video"
}

//
// QueryVideoByAuthor
//  @Description: 根据作者名字查询视频集
//  @author XiaoWenzhuo
//  @param ctx
//  @param author_id
//  @return []*Video
//  @return error
//
func QueryVideoByAuthor(ctx context.Context, authorId int64) ([]*Video, error) {
	var res []*Video
	if err := DB.WithContext(ctx).Where("author_id = ?", authorId).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

//
// CreateVideo
//  @Description: 插入视频
//  @author XiaoWenzhuo
//  @param ctx
//  @param video
//  @return error
//
func CreateVideo(ctx context.Context, video Video) error {
	if err := DB.WithContext(ctx).Create(&video).Error; err != nil {
		return err
	}
	return nil
}
