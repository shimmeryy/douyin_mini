package db

import (
	"context"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	UserId  int64  `gorm:"column:user_id;type:int not null;comment:用户id"`
	VideoId int64  `gorm:"column:video_id;type:int not null;comment:视频id"`
	Text    string `gorm:"column:text;type:varchar(255) not null; comment:评论内容"`
}

func (v Comment) TableName() string {
	return "comment"
}

// QueryCommentByVideoId 根据videoId查询所有的评论列表
func QueryCommentByVideoId(ctx context.Context, videoId int64) ([]*Comment, error) {
	var res []*Comment
	if err := DB.WithContext(ctx).Where("video_id = ?", videoId).Order("created_at desc").Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// CreateCommentInfo 创建新的评论信息
func CreateCommentInfo(ctx context.Context, comment Comment) error {
	if err := DB.WithContext(ctx).Create(&comment).Error; err != nil {
		return err
	}
	return nil
}

// DeleteCommentInfo 根据commentId删除评论
func DeleteCommentInfo(ctx context.Context, commentId int64) error {
	if err := DB.WithContext(ctx).Delete(&Comment{}, commentId).Error; err != nil {
		return err
	}
	return nil
}

// CountCommentByVideoId 根据videoId查询评论数量
func CountCommentByVideoId(ctx context.Context, videoId int64) (int64, error) {
	var cnt int64
	if err := DB.WithContext(ctx).Model(&Comment{}).Where("video_id = ?", videoId).Count(&cnt).Error; err != nil {
		return cnt, err
	}
	return cnt, nil
}
