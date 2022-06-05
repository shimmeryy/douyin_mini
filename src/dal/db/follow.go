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

// CreatFollow 增加关注
func CreatFollow(ctx context.Context, userId int64, followUserId int64) (err error) {
	follow := Follow{
		UserId:       userId,
		FollowUserId: followUserId,
	}
	err = DB.WithContext(ctx).Create(&follow).Error
	return
}

// DeleteFollow 取消关注
func DeleteFollow(ctx context.Context, userId int64, followUserId int64) (err error) {
	err = DB.WithContext(ctx).Where(&Follow{
		UserId:       userId,
		FollowUserId: followUserId,
	}).Delete(&Follow{}).Error
	return
}

// IsFollow 判断是否关注
func IsFollow(ctx context.Context, userId int64, followUserId int64) (isFollow bool, err error) {
	err = DB.WithContext(ctx).Where(&Follow{
		UserId:       userId,
		FollowUserId: followUserId,
	}).First(&Follow{}).Error
	//当数据库查找不到相关记录，则用户userId 没有关注 用户followUserId
	isFollow = !errors.Is(err, gorm.ErrRecordNotFound)
	return
}

// QueryFollowerListById 根据用户Id查询该用户的关注列表
func QueryFollowerListById(ctx context.Context, userId int64) (users []*User, err error) {
	err = DB.WithContext(ctx).
		Joins("RIGHT JOIN follow ON user.id = follow.follow_user_id AND follow.user_id = ? AND follow.deleted_at IS NULL", userId).
		Find(&users).Error
	return
}

// QueryFansListById 根据用户Id查询该用户的粉丝列表
func QueryFansListById(ctx context.Context, userId int64) (users []*User, err error) {
	err = DB.WithContext(ctx).
		Joins("RIGHT JOIN follow ON user.id = follow.user_id AND follow.follow_user_id = ? AND follow.deleted_at IS NULL", userId).
		Find(&users).Error
	return
}
