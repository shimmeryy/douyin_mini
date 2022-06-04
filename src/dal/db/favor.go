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

// CountFavorByVideoId 根据videoId获取赞数量
func CountFavorByVideoId(ctx context.Context, videoId int64) (int64, error) {
	var count int64
	if err := DB.WithContext(ctx).Model(&Favor{}).Where("video_id = ?", videoId).Count(&count).Error; err != nil {
		return count, err
	}
	return count, nil
}

// CheckIsFavored 根据UserId和VideoId来检查是否用户点赞过该视频
func CheckIsFavored(ctx context.Context, userId int64, videoId int64) (bool, error) {
	var cnt int64
	if err := DB.WithContext(ctx).Model(&Favor{}).Where("user_id = ? AND video_id = ?", userId, videoId).Count(&cnt).Error; err != nil {
		return false, err
	}
	if cnt == 0 {
		return false, nil
	}
	return true, nil
}
