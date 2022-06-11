package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tiktok/src/constants"
	"tiktok/src/errno"
	"tiktok/src/handlers"
	"tiktok/src/service"
	"tiktok/src/utils/jwt"
)

// Register 新用户注册
func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	userService := service.UserServiceInstance()

	response, err := userService.RegisterUser(c, username, password)
	if err != nil {
		c.JSON(http.StatusBadRequest, &response)
		return
	}
	c.JSON(http.StatusOK, &response)
}

// GetUserInfo 获取用户信息
func GetUserInfo(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		panic(errno.ParamErr)
	}
	userInfo, err := service.UserServiceInstance().GetUserInfo(c, userId)
	if err != nil {
		panic(errno.ServiceErr.WithMessage(err.Error()))
	}
	c.JSON(http.StatusOK, &handlers.UserInfoResponse{
		Response: handlers.Response{
			StatusCode: 0,
			StatusMsg:  "成功",
		},
		User: *userInfo,
	})
}

// FollowList 获取关注列表
func FollowList(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		panic(errno.ParamErr.WithMessage("invalid UserId"))
	}

	if flag := service.UserServiceInstance().CheckUserById(c, userId); flag != true {
		panic(errno.ParamErr.WithMessage("invalid UserId"))
	}

	followers, err := service.UserServiceInstance().GetUserFollowers(c, userId)
	if err != nil {
		panic(errno.ServiceErr.WithMessage(err.Error()))
	}
	c.JSON(http.StatusOK, &handlers.UserFollowersResponse{
		Response: handlers.Response{
			StatusCode: 0,
			StatusMsg:  "成功",
		},
		Followers: followers,
	})
}

// FollowerList 获取粉丝列表
func FollowerList(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		panic(errno.ParamErr.WithMessage("invalid UserId"))
	}

	if flag := service.UserServiceInstance().CheckUserById(c, userId); flag != true {
		panic(errno.ParamErr.WithMessage("invalid UserId"))
	}

	fans, err := service.UserServiceInstance().GetUserFans(c, userId)
	if err != nil {
		panic(errno.ServiceErr.WithMessage(err.Error()))
	}
	c.JSON(http.StatusOK, &handlers.UserFansResponse{
		Response: handlers.Response{
			StatusCode: 0,
			StatusMsg:  "成功",
		},
		Fans: fans,
	})
}

// RelationAction 关注或取关用户
func RelationAction(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	userId := int64(claims[constants.IdentityKey].(float64))

	followUserId, err := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
	if err != nil {
		panic(errno.ParamErr.WithMessage("invalid FollowUserId"))
	}

	if userId == followUserId {
		panic(errno.ParamErr.WithMessage("can't follow yourself!"))
	}

	actionType, err := strconv.ParseInt(c.Query("action_type"), 10, 64)
	if err != nil || (actionType != 1 && actionType != 2) {
		panic(errno.ParamErr.WithMessage("invalid actionType"))
	}

	if actionType == 1 {
		err := service.UserServiceInstance().FollowUser(c, userId, followUserId)
		if err != nil {
			panic(errno.ServiceErr.WithMessage(err.Error()))
		}
	} else {
		err := service.UserServiceInstance().CancelFollowUser(c, userId, followUserId)
		if err != nil {
			panic(errno.ServiceErr.WithMessage(err.Error()))
		}
	}

	c.JSON(http.StatusOK, &handlers.Response{
		StatusCode: 0,
		StatusMsg:  "success",
	})
}
