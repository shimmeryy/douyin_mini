package db

import (
	"context"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName      string `gorm:"column:username;type:varchar(32)not null;comment:用户名"`
	PassWord      string `gorm:"column:password;type:varchar(255)not null;comment:密码"`
	FollowCount   int64  `gorm:"column:follow_count;default:0;comment:关注总数"`
	FollowerCount int64  `gorm:"column:follower_count;default:0;comment:粉丝总数"`
}

func (v User) TableName() string {
	return "user"
}

//QueryUser 根据用户名查询用户
func QueryUser(ctx context.Context, userName string) ([]*User, error) {
	var res []*User
	if err := DB.WithContext(ctx).Where("username = ?", userName).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

//QueryUserById 根据用户ID查询用户
func QueryUserById(ctx context.Context, ID int64) (*User, error) {
	var res *User
	if err := DB.WithContext(ctx).Where("id = ?", ID).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
