package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tiktok/src/constants"
	"tiktok/src/handlers"
	"tiktok/src/service"
	"tiktok/src/utils/jwt"
	"time"
)

func InitAuthMiddleware() (*jwt.GinJWTMiddleware, error) {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Key:        []byte(constants.SecretKey), //秘钥
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		//设置token载荷②
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(int64); ok {
				return jwt.MapClaims{
					//对应的id
					constants.IdentityKey: v,
				}
			}
			return jwt.MapClaims{}
		},
		//login身份认证①
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVar handlers.UserLoginParam
			if err := c.ShouldBindQuery(&loginVar); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			//校验用户是否存在（获取到用户信息）
			userService := service.UserServiceInstance()
			userId, err := userService.CheckUser(c, loginVar)
			if err == nil {
				c.Set(constants.IdentityKey, userId) //存储下用户id在context，由于登录接口不会走解析token操作，当前context里不会有claims
			}
			return userId, err //用户id,err，在框架中会检测是否是否异常，若是有异常走
		},
		//若是身份未认证成功。情况：①登录失败。②token不正确。③权限不足。
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"status_code": code,
				"status_msg":  message,
			})
		},
		//自定义登录响应内容
		LoginResponse: func(c *gin.Context, code int, token string, time time.Time) {
			userID, _ := c.Get(constants.IdentityKey)
			c.JSON(http.StatusOK, &handlers.UserLoginResponse{
				Response: handlers.Response{
					StatusCode: 0,
					StatusMsg:  "登录成功",
				},
				UserId: userID.(int64),
				Token:  token,
			})
		},
		//token的header的key为Authorization
		TokenLookup: "header: Authorization, query: token, cookie: jwt, form: token",
		//对应Authorization的值前缀为Bearer
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})
	if err != nil {
		return nil, err
	}
	return authMiddleware, nil
}
