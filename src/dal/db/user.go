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

//UpdateFollowAndFollowerCount 更新用户与关注用户的关注总数与粉丝数
//示例：两者增加1 num=1；两者减1 num=-1
func UpdateFollowAndFollowerCount(ctx context.Context, userId int64, followUserId int64, num int64) error {
	err := DB.WithContext(ctx).Model(&User{}).Where("id = ?", userId).UpdateColumn("follow_count", gorm.Expr("follow_count + ?", num)).Error
	if err != nil {
		return err
	}
	err = DB.WithContext(ctx).Model(&User{}).Where("id = ?", followUserId).UpdateColumn("follower_count", gorm.Expr("follower_count + ?", num)).Error
	if err != nil {
		return err
	}
	return nil
}
