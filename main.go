package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"runtime/debug"
	"tiktok/src/constants"
	"tiktok/src/controller"
	"tiktok/src/dal/db"
	"tiktok/src/errno"
	"tiktok/src/handlers"
	"tiktok/src/redis"
	"tiktok/src/service"
	"tiktok/src/utils/jwt"
	"time"
)

func main() {
	r := gin.Default()
	//config.Init() //初始化配置
	InitRouter(r) //初始化路由+全局异常捕捉+jwt插件
	redis.Init()
	db.Init() //初始化数据库
	r.Run()
}

//初始化路由
func InitRouter(r *gin.Engine) {
	//静态资源映射
	// public directory is used to serve static resources
	r.Static("/static", "./public")
	//全局异常捕捉
	r.Use(RecoverMiddleware)

	//获取认证插件
	var authMiddleware *jwt.GinJWTMiddleware
	middleware, err := InitAuthMiddleware()
	if err != nil {
		panic(err)
	}
	authMiddleware = middleware

	//设置douyin分组
	apiRouter := r.Group("/douyin")

	//User：登录、注册、获取用户信息接口
	apiRouter.POST("/user/login/", authMiddleware.LoginHandler)
	apiRouter.POST("/user/register/", controller.Register)
	user := apiRouter.Group("/user")
	{
		user.Use(authMiddleware.MiddlewareFunc())
		user.GET("/", controller.GetUserInfo)
	}

	// basic apis
	//apiRouter.GET("/feed/", controller.Feed) //feed流接口
	//apiRouter.POST("/publish/action/", controller.Publish)  //发布视频
	//apiRouter.GET("/publish/list/", controller.PublishList) //发布列表
	//
	//// extra apis - I
	//apiRouter.POST("/favorite/action/", controller.FavoriteAction) //点赞
	//apiRouter.GET("/favorite/list/", controller.FavoriteList)      //查看喜欢列表
	//apiRouter.POST("/comment/action/", controller.CommentAction)   //评论
	//apiRouter.GET("/comment/list/", controller.CommentList)        //查看评论列表
	//
	//// extra apis - II
	//apiRouter.POST("/relation/action/", controller.RelationAction)     //关注用户
	//apiRouter.GET("/relation/follow/list/", controller.FollowList)     //关注列表
	//apiRouter.GET("/relation/follower/list/", controller.FollowerList) //粉丝列表
}

//获取认证插件
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
				"code":    code,
				"message": message,
			})
		},
		//自定义登录响应内容
		LoginResponse: func(c *gin.Context, code int, token string, time time.Time) {
			userID, _ := c.Get(constants.IdentityKey)
			c.JSON(http.StatusOK, &handlers.UserLoginResponse{
				Response: handlers.Response{
					StatusCode: http.StatusOK,
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

//全局异常捕捉插件
func RecoverMiddleware(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			//打印错误堆栈信息
			log.Printf("panic: %v\n", r)
			debug.PrintStack()
			//若是非自定义异常，将其进行转换为自定异常统一返回

			err, ok := r.(errno.ErrNo)
			if !ok {
				r = errno.ConvertErr(err)
			}
			//封装通用JSON返回
			c.JSON(http.StatusOK, gin.H{
				"code": err.ErrCode,
				"msg":  err.ErrMsg,
			})
			//终止后续接口调用，不加的话recover到异常后，还会继续执行接口里后续代码
			c.Abort()
		}
	}()
	//加载完 defer recover，继续后续的插件及代码执行
	c.Next()
}

//错误转字符串
func errorToString(r interface{}) string {
	switch v := r.(type) {
	case error:
		return v.Error()
	default:
		return r.(string)
	}
}
