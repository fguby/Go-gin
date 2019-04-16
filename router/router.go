package router

import (
	jwt "Gin_demo/middleware"
	"Gin_demo/pkg/util"
	"Gin_demo/router/api"
	v1 "Gin_demo/router/api/v1"
	"net/http"

	_ "Gin_demo/docs"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(util.ServerSetting.RunMode)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		apiv1.GET("/tags", v1.GetTags)
		apiv1.POST("/tags", v1.AddTag)
		apiv1.PUT("/tags/:id", v1.EditTag)
		apiv1.DELETE("/tags/:id", v1.DeleteTag)
	}
	//配置页面返回
	r.LoadHTMLGlob("views/*")
	//配置静态文件路径
	r.Static("/static", "./static")
	//r.Static("/swagger", "./docs")
	//配置首页
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"title": "首页"})
	})
	//配置获取token的路径
	r.GET("/auth", api.GetAuth)
	//配置swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
