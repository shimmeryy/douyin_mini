package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tiktok/src/errno"
	"tiktok/src/handlers"
	"tiktok/src/service"
)

//Register User Register
func Register(c *gin.Context) {
}

//Login User Login
//func Login(c *gin.Context) {
//}

//UserInfo Get UserInfo
func GetUserInfo(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		panic(errno.ServiceErr.WithMessage("参数ID有误"))
	}
	userInfo, err := service.UserServiceInstance().GetUserInfo(c, userId)
	if err != nil {
		panic(errno.ServiceErr.WithMessage(err.Error()))
	}
	c.JSON(http.StatusOK, &handlers.UserInfoResponse{
		Response: handlers.Response{
			StatusCode: 200,
			StatusMsg:  "成功",
		},
		User: *userInfo,
	})
}
