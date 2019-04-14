package router

import (
	"Gin_demo/pkg/util"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	println("运行模式为:%s", util.ServerSetting.RunMode)
	gin.SetMode(util.ServerSetting.RunMode)

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})

	return r
}
