package routers

import (
	"Java2GO/middlewire/cors"
	"Java2GO/routers/api"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(cors.CorsHandler())
	r.Use(gin.Recovery())
	apiRouter := r.Group("/douyin")

	publish := apiRouter.Group("/publish")
	publish.POST("/action", api.Publish)
	return r
}
