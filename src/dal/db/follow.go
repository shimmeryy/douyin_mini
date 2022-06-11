package db

import (
	"context"
	"errors"
	"gorm.io/gorm"
)

type Follow struct {
	gorm.Model
	UserId       int64 `gorm:"column:user_id;type:bigint not null;comment:用户id"`
	FollowUserId int64 `gorm:"column:follow_user_id;type:bigint not null;comment:用户id"`
}

func (f Follow) TableName() string {
	return "follow"
}

// CreateFollow 增加关注
func CreateFollow(ctx context.Context, userId int64, followUserId int64) error {
	follow := Follow{
		UserId:       userId,
		FollowUserId: followUserId,
	}
	err := DB.WithContext(ctx).Create(&follow).Error
	return err
}

// DeleteFollow 取消关注
func DeleteFollow(ctx context.Context, userId int64, followUserId int64) error {
	err := DB.WithContext(ctx).Where(&Follow{
		UserId:       userId,
		FollowUserId: followUserId,
	}).Delete(&Follow{}).Error
	return err
}

// IsFollow 判断是否关注
func IsFollow(ctx context.Context, userId int64, followUserId int64) (bool, error) {
	err := DB.WithContext(ctx).Where(&Follow{
		UserId:       userId,
		FollowUserId: followUserId,
	}).First(&Follow{}).Error
	//当数据库查找不到相关记录，则用户userId 没有关注 用户followUserId
	isFollow := !errors.Is(err, gorm.ErrRecordNotFound)
	return isFollow, err
}

// QueryFollowerListById 根据用户Id查询该用户的关注列表
func QueryFollowerListById(ctx context.Context, userId int64) ([]*User, error) {
	var users []*User
	var res []*User
	err := DB.WithContext(ctx).
		Joins("RIGHT JOIN follow ON user.id = follow.follow_user_id AND follow.user_id = ? AND follow.deleted_at IS NULL", userId).
		Find(&users).Error
	for _, tmp := range users {
		if tmp.ID != 0 {
			res = append(res, tmp)
		}
	}
	return res, err
}

// QueryFansListById 根据用户Id查询该用户的粉丝列表
func QueryFansListById(ctx context.Context, userId int64) ([]*User, error) {
	var users []*User
	var res []*User
	err := DB.WithContext(ctx).
		Joins("RIGHT JOIN follow ON user.id = follow.user_id AND follow.follow_user_id = ? AND follow.deleted_at IS NULL", userId).
		Find(&users).Error
	for _, tmp := range users {
		if tmp.ID != 0 {
			res = append(res, tmp)
		}
	}
	return res, err
}
