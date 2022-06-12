package main

import (
	"github.com/gin-gonic/gin"
	"tiktok/src/dal/db"
	"tiktok/src/redis"
	"tiktok/src/router"
	"tiktok/src/utils/ossUtil"
)

func main() {
	r := gin.Default()
	//config.Init() //初始化配置
	router.InitRouter(r) //初始化路由+全局异常捕捉+jwt插件
	ossUtil.InitBucket()
	redis.Init()
	db.Init() //初始化数据库
	r.Run()
}

