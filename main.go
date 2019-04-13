package main

import (
	_ "Gin_demo/init"
	"Gin_demo/pkg/util"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/form", func(c *gin.Context) {
		type1 := c.DefaultPostForm("type", "alert") //可设置默认值
		username := c.PostForm("username")
		password := c.PostForm("password")

		//hobbys := c.PostFormMap("hobby")
		//hobbys := c.QueryArray("hobby")
		hobbys := c.PostFormArray("hobby")

		c.String(http.StatusOK, fmt.Sprintf("type is %s, username is %s, password is %s,hobby is %v", type1, username, password, hobbys))

	})
	//读取端口号
	port := util.GetIniValue("server", "port")
	router.Run(":" + port) // listen and serve on 0.0.0.0:8080
}
