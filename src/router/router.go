package router

import (
	"github.com/gin-gonic/gin"
	"tiktok/src/config"
	"tiktok/src/controller"
	"tiktok/src/middlewares"
)

func InitRouter(r *gin.Engine) {
	//静态资源映射
	// public directory is used to serve static resources
	r.Static("/static", "./public")
	//全局异常捕捉
	r.Use(middlewares.RecoverMiddleware)

	//获取认证插件
	mid, err := middlewares.InitAuthMiddleware()
	config.AuthMiddleware = mid
	if err != nil {
		panic(err)
	}

	//设置douyin分组
	apiRouter := r.Group("/douyin")

	//User：登录、注册、获取用户信息接口
	apiRouter.POST("/user/login/", config.AuthMiddleware.LoginHandler)
	apiRouter.POST("/user/register/", controller.Register)
	user := apiRouter.Group("/user")
	{
		user.Use(config.AuthMiddleware.MiddlewareFunc())
		user.GET("/", controller.GetUserInfo)
	}
	publish := apiRouter.Group("/publish")
	{
		publish.Use(config.AuthMiddleware.MiddlewareFunc())
		publish.POST("/action/", controller.Publish)  //发布视频
		publish.GET("/list/", controller.PublishList) //发布列表
	}
	favorite := apiRouter.Group("/favorite")
	{
		favorite.Use(config.AuthMiddleware.MiddlewareFunc())
		favorite.POST("/action/", controller.FavoriteAction)
		favorite.GET("/list/", controller.FavoriteList)
	}
	comment := apiRouter.Group("/comment")
	{
		comment.Use(config.AuthMiddleware.MiddlewareFunc())
		comment.POST("/action/", controller.CommentAction)
		comment.GET("/list/", controller.CommentList)
	}
	relation := apiRouter.Group("/relation")
	{
		relation.Use(config.AuthMiddleware.MiddlewareFunc())
		relation.POST("/action/", controller.RelationAction)
		relation.GET("/follow/list/", controller.FollowList)
		relation.GET("/follower/list/", controller.FollowerList)
	}
	// basic apis
	apiRouter.GET("/feed/", controller.Feed) //feed流接口
}
