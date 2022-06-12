package service

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"sync"
	"tiktok/src/config"
	"tiktok/src/constants"
	"tiktok/src/dal/db"
	"tiktok/src/errno"
	"tiktok/src/handlers"
	"tiktok/src/utils"
	"tiktok/src/utils/jwt"
)

type UserService interface {
	CreateUser()
	CheckUser()
	GetUserInfo()
	CheckUserById()
	RegisterUser()
	CheckPassword()
	GetToken()
	GetUserFollowers()
	GetUserFans()
	FollowUser()
	CancelFollowUser()
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

// UserServiceImpl 用户业务层实现类
type UserServiceImpl struct {
}

// CheckUser 检测用户是否存在
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

func (this *UserServiceImpl) GetUserInfo(ctx context.Context, ID int64) (*handlers.UserInfo, error) {
	claims := jwt.ExtractClaims(ctx.(*gin.Context))
	userID := int64(claims[constants.IdentityKey].(float64))
	//1、根据用户id查询用户
	targetUser, err := db.QueryUserById(ctx, ID)
	if err != nil {
		return nil, err
	}
	if targetUser != nil && targetUser.ID == 0 {
		return nil, errno.ServiceErr.WithMessage("用户不存在")
	}
	//2、查看当前用户是否关注对方
	isFollow, _ := db.IsFollow(ctx, userID, int64(targetUser.ID))
	//结构体转换
	userInfo := &handlers.UserInfo{
		ID:            int64(targetUser.ID),
		UserName:      targetUser.UserName,
		FollowCount:   targetUser.FollowCount,
		FollowerCount: targetUser.FollowerCount,
		IsFollow:      isFollow,
	}
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

func (this *UserServiceImpl) RegisterUser(c context.Context, username string, password string) (*handlers.UserRegisterResponse, error) {
	if err := this.CheckPassword(password); err != nil {
		return &handlers.UserRegisterResponse{
			StatusCode: 1,
			StatusMsg:  "密码过长或为空",
		}, err
	}

	var usr db.User
	db.DB.WithContext(c).Where("username = ?", username).Find(&usr)
	if usr.ID > 0 { //用户名已注册
		return &handlers.UserRegisterResponse{StatusCode: 1, StatusMsg: "用户名已存在"}, errors.New("用户名已存在")
	} else {
		//创建用户
		saltPassword, err := utils.HashAndSalt(password)
		if err != nil {
			return &handlers.UserRegisterResponse{
				StatusCode: 1,
				StatusMsg:  "加密失败",
			}, err
		}
		db.DB.WithContext(c).Create(&db.User{
			UserName:      username,
			PassWord:      saltPassword,
			FollowCount:   0,
			FollowerCount: 0,
		})
		var u db.User
		err = db.DB.WithContext(c).Where("username = ?", username).Find(&u).Error
		tokenString, err := this.GetToken(&u)
		if err != nil {
			db.DB.WithContext(c).Delete(&u)
			return &handlers.UserRegisterResponse{
				StatusCode: 1,
				StatusMsg:  "token获取失败",
			}, err
		}
		return &handlers.UserRegisterResponse{
			StatusMsg: "用户创建成功",
			UserId:    int64(u.ID),
			Token:     tokenString,
		}, nil
	}
}

func (this *UserServiceImpl) CheckPassword(password string) error {
	length := len(password)
	if length == 0 {
		return errno.ParamErr.WithMessage("密码不能为空")
	}
	if length > constants.MaxPasswordLength {
		return errno.ParamErr.WithMessage("密码长度过长")
	}
	return nil
}

func (this *UserServiceImpl) GetToken(user *db.User) (string, error) {
	s, _, err := config.AuthMiddleware.TokenGenerator(jwt.MapClaims{
		// 对应的id
		constants.IdentityKey: user.ID,
	})
	if err != nil {
		return "", err
	}
	return s, nil
}

func (this *UserServiceImpl) GetUserFollowers(ctx context.Context, ID int64) ([]handlers.UserInfo, error) {
	//1、根据用户id查询用户的关注用户列表
	followers, err := db.QueryFollowerListById(ctx, ID)
	if err != nil || len(followers) <= 0 {
		return nil, err
	}
	//结构体转换
	infos := make([]handlers.UserInfo, 0)
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
	infos := make([]handlers.UserInfo, 0)
	for i := 0; i < len(fans); i++ {
		//1、查询当前用户是否关注对方用户
		isFollow, _ := db.IsFollow(ctx, ID, int64(fans[i].ID))

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
	flag, err := db.IsFollow(ctx, userId, followUserId)
	if flag { // 已关注，直接返回
		return nil
	}
	err = db.CreateFollow(ctx, userId, followUserId)
	if err != nil {
		return err
	}
	err = db.UpdateFollowAndFollowerCount(ctx, userId, followUserId, 1)
	if err != nil {
		return err
	}
	return nil
}

func (this *UserServiceImpl) CancelFollowUser(ctx context.Context, userId int64, followUserId int64) error {
	// 1. 已取关，直接返回
	flag, err := db.IsFollow(ctx, userId, followUserId)
	if !flag { // 已取关，直接返回
		return nil
	}
	// 2.关注粉丝表（Follow表）中创建新数据
	err = db.DeleteFollow(ctx, userId, followUserId)
	if err != nil {
		return err
	}
	// 3.用户表（user表）关注者userId关注数减少、被关注者followUserId粉丝数减少
	err = db.UpdateFollowAndFollowerCount(ctx, userId, followUserId, -1)
	if err != nil {
		return err
	}
	return nil
}
