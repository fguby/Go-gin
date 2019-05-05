package router

import (
	"Gin_demo/pkg/util"
	"Gin_demo/router/api"
	v1 "Gin_demo/router/api/v1"

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
	//配置自定义中间件
	//apiv1.Use(jwt.JWT())
	{
		//标签路由配置
		apiv1.GET("/tags", v1.GetTags)
		apiv1.POST("/tags", v1.AddTag)
		apiv1.PUT("/tags/:id", v1.EditTag)
		apiv1.DELETE("/tags/:id", v1.DeleteTag)

		apiv1.GET("/models", v1.GetModelJson)

		//文章路由配置
		apiv1.GET("/article", v1.ToArticlePage)
	}
	//配置页面返回
	//r.LoadHTMLGlob("views/*")
	//配置静态文件路径
	r.Static("/static", "./static")
	//配置静态文件路径
	//r.StaticFile("/static/assets/waifu-tips.min.js", "./assets/waifu-tips.min.js")
	//配置首页
	// r.GET("/", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "index.html", gin.H{"title": "首页"})
	// })
	// r.GET("/test1", func(c *gin.Context) {
	// 	c.HTML(200, "test1.html", gin.H{"title": "干物妹小埋"})
	// })
	// r.GET("/test2", func(c *gin.Context) {
	// 	c.HTML(200, "test2.html", gin.H{"title": "干物妹小埋"})
	// })
	//配置获取token的路径
	r.GET("/auth", api.GetAuth)
	//配置swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
