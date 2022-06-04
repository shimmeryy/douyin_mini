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
	//2、查询当前用户是否关注对方用户
	return userInfo, nil
}

func (this *UserServiceImpl) GetUserFollowers(ctx context.Context, ID int64) ([]handlers.UserInfo, error) {
	//1、根据用户id查询用户的关注用户列表
	followers, err := db.QueryFollowerListById(ctx, ID)
	if err != nil || len(followers) <= 0 {
		return nil, err
	}
	//结构体转换
	infos := make([]handlers.UserInfo, len(followers))
	for i := 0; i < len(followers); i++ {
		userInfo := &handlers.UserInfo{
			ID:            int64(followers[i].ID),
			UserName:      followers[i].UserName,
			FollowCount:   followers[i].FollowCount,
			FollowerCount: followers[i].FollowerCount,
			IsFollow:      true,
		}
		infos = append(infos, *userInfo)
	}

	return infos, nil
}

func (this *UserServiceImpl) GetUserFans(ctx context.Context, ID int64) ([]handlers.UserInfo, error) {
	//1、根据用户id查询用户的粉丝用户列表
	fans, err := db.QueryFansListById(ctx, ID)
	if err != nil || len(fans) <= 0 {
		return nil, err
	}
	//结构体转换
	infos := make([]handlers.UserInfo, len(fans))
	for i := 0; i < len(fans); i++ {
		//1、查询当前用户是否关注对方用户
		isFollow, err := db.IsFollow(ctx, ID, int64(fans[i].ID))
		if err != nil {
			return nil, err
		}

		userInfo := &handlers.UserInfo{
			ID:            int64(fans[i].ID),
			UserName:      fans[i].UserName,
			FollowCount:   fans[i].FollowCount,
			FollowerCount: fans[i].FollowerCount,
			IsFollow:      isFollow,
		}

		infos = append(infos, *userInfo)
	}

	return infos, nil
}

func (this *UserServiceImpl) FollowUser(ctx context.Context, userId int64, followUserId int64) error {
	//1、关注粉丝表（Follow表）中创建新数据
	err := db.CreatFollow(ctx, userId, followUserId)
	if err != nil {
		return err
	}
	//TODO 2、redis缓存数据修改
	//TODO 3、用户表（user表）关注者userId关注数增加、被关注者followUserId粉丝数增加
	return nil
}

func (this *UserServiceImpl) CancelFollowUser(ctx context.Context, userId int64, followUserId int64) error {
	//1、关注粉丝表（Follow表）中创建新数据
	err := db.DeleteFollow(ctx, userId, followUserId)
	if err != nil {
		return err
	}
	//TODO 2、redis缓存数据修改
	//TODO 3、用户表（user表）关注者userId关注数减少、被关注者followUserId粉丝数减少
	return nil
}
