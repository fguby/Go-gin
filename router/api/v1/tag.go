package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTags(c *gin.Context) {
	c.JSON(http.StatusOK, "获取成功")
}

func AddTag(c *gin.Context) {
	c.JSON(http.StatusOK, "添加成功")
}

func EditTag(c *gin.Context) {
	c.JSON(http.StatusOK, "修改成功")
}

func DeleteTag(c *gin.Context) {
	c.JSON(http.StatusOK, "删除成功")
}
