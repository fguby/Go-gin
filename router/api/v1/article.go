package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//跳转文章页面
func ToArticlePage(c *gin.Context) {
	c.HTML(http.StatusOK, "article.html", gin.H{"title": "首页"})
}
