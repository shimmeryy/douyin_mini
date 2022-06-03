package db

import (
	"context"
	"gorm.io/gorm"
)

type Favor struct {
	gorm.Model
	UserId  int64 `gorm:"column:user_id;type:int not null;comment:用户id"`
	VideoId int64 `gorm:"column:video_id;type:int not null;comment:视频id"`
}

func (v Favor) TableName() string {
	return "favor"
}

// QueryFavorByUserId 根据用户Id查询所有点赞列表
func QueryFavorByUserId(ctx context.Context, userId int64) ([]*Favor, error) {
	var res []*Favor
	if err := DB.WithContext(ctx).Where("user_id = ?", userId).Order("created_at desc").Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// CreateFavorInfo 点赞
func CreateFavorInfo(ctx context.Context, favor Favor) error {
	if err := DB.WithContext(ctx).Create(&favor).Error; err != nil {
		return err
	}
	return nil
}

// DeleteFavorInfo 根据userId和videoId删除点赞信息
func DeleteFavorInfo(ctx context.Context, userId int64, videoId int64) error {
	if err := DB.WithContext(ctx).Where("user_id = ? and video_id = ?", userId, videoId).Delete(&Favor{}).Error; err != nil {
		return err
	}
	return nil
}
