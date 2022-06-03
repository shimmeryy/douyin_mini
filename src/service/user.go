package service

import (
	"context"
	"sync"
	"tiktok/src/dal/db"
	"tiktok/src/errno"
	"tiktok/src/handlers"
	"tiktok/src/utils"
)

type UserService interface {
	CreateUser()
	CheckUser(ctx context.Context, req handlers.UserLoginParam) (uint, error)
}

var (
	userService     *UserServiceImpl
	userServiceOnce sync.Once
)

func UserServiceInstance() *UserServiceImpl {
	userServiceOnce.Do(
		func() {
			userService = &UserServiceImpl{}
		})
	return userService
}

//用户业务层实现类
type UserServiceImpl struct {
}

//检测用户是否存在
func (this *UserServiceImpl) CheckUser(ctx context.Context, req handlers.UserLoginParam) (int64, error) {
	name := req.UserName
	//1、用户是否存在
	users, err := db.QueryUser(ctx, name)
	if err != nil || len(users) == 0 {
		return 0, errno.UserNotExistErr
	}
	user := users[0]
	//2、密码是否正确
	if !utils.ComparePasswords(user.PassWord, req.PassWord) {
		return 0, errno.LoginErr
	}
	return int64(user.ID), nil
}

func (this *UserServiceImpl) CreateUser(ctx context.Context) {
	panic("implement me")
}

func (this *UserServiceImpl) GetUserInfo(ctx context.Context, ID int64) (*handlers.UserInfo, error) {
	//1、根据用户id查询用户
	user, err := db.QueryUserById(ctx, ID)
	if err != nil {
		return nil, err
	}
	if user != nil && user.ID == 0 {
		return nil, errno.ServiceErr.WithMessage("用户不存在")
	}
	//结构体转换
	userInfo := &handlers.UserInfo{
		ID:            int64(user.ID),
		UserName:      user.UserName,
		FollowCount:   user.FollowCount,
		FollowerCount: user.FollowerCount,
		IsFollow:      false,
	}
	//2、todo 查看是否关注对方：根据userID、targetUserId来判断，返回true或者false
	return userInfo, nil
}

// CheckUserById user exist return true ,other return false
func (this *UserServiceImpl) CheckUserById(ctx context.Context, ID int64) bool {
	user, err := db.QueryUserById(ctx, ID)
	if err != nil {
		return false
	}
	if user != nil && user.Model.ID == 0 {
		return false
	}
	return true
}
