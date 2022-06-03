package service

import (
	"context"
	"errors"
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
	CheckUser(ctx context.Context, req handlers.UserLoginParam) (uint, error)
}
type UserRegisterResponse struct {
	StatusCode int32  `json:"status___code"`
	StatusMsg  string `json:"status___msg,omitempty"`
	UserId     int64  `json:"user___id"`
	Token      string `json:"token"`
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

func (this *UserServiceImpl) RegisterUser(username string, password string) (*UserRegisterResponse, error) {
	if err := this.CheckPassword(password); err != nil {
		return &UserRegisterResponse{
			1, "密码过长或为空", 0, "",
		}, err
	}

	var usr db.User
	db.DB.Where("username=?", username).Find(&usr)
	if usr.ID > 0 { //用户名已注册
		return &UserRegisterResponse{1, "用户名已存在", 0, ""}, errors.New("用户名已存在")
	} else {
		//创建用户
		saltPassword, err := utils.HashAndSalt(password)
		if err != nil {
			return &UserRegisterResponse{
				1,
				"加密失败",
				0,
				"",
			}, err
		}
		db.DB.Create(&db.User{UserName: username, PassWord: saltPassword, FollowCount: 0, FollowerCount: 0})
		var u db.User
		db.DB.Where("username=?", username).Find(&u)
		tokenString, err := this.GetToken(&u)
		if err != nil {
			return &UserRegisterResponse{
				1,
				"token获取失败",
				0,
				"",
			}, err
		}
		return &UserRegisterResponse{
			1,
			"用户创建成功",
			int64(u.ID),
			tokenString,
		}, nil
	}
}

func (this *UserServiceImpl) CheckPassword(password string) error {
	length := len(password)
	if length == 0 {
		return errors.New("密码不能为空")
	}
	if length > constants.MaxPasswordlength {
		return errors.New("密码长度过长")
	}
	return nil
}

//func (this *UserServiceImpl) GetToken2(user *db.User) (string, error) {
//	expiresTime := time.Now().Add(7 * 24 * time.Hour)
//	claims := jwt.StandardClaims{
//		Audience:  user.UserName,      // 受众
//		ExpiresAt: expiresTime.Unix(), // 失效时间
//		Id:        string(user.ID),    // 编号
//		IssuedAt:  time.Now().Unix(),  // 签发时间
//		Issuer:    "gin hello",        // 签发人
//		NotBefore: time.Now().Unix(),  // 生效时间
//		Subject:   "login",            // 主题
//	}
//	var jwtSecret = []byte("jklasdfi.as")
//	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
//	token, err := tokenClaims.SignedString(jwtSecret)
//	if err != nil {
//		return "", err
//	}
//	return token, nil
//}

func (this *UserServiceImpl) GetToken(user *db.User) (string, error) {
	s, _, err := config.AuthMiddleware.TokenGenerator(jwt.MapClaims{
		//对应的id
		constants.IdentityKey: user.ID,
	})
	if err != nil {
		return "", err
	}
	return s, err
}
