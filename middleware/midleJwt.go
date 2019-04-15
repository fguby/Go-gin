package middleware

import (
	cs "Gin_demo/pkg/constant"
	"Gin_demo/pkg/util"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

//设置自定义中间件
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = cs.SUCCESS
		token := c.Query("token")
		if token == "" {
			//无效参数
			code = cs.INVALID_PARAMS
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				//token鉴权失败
				code = cs.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				//token已经超时
				code = cs.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}
		if code != cs.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  cs.GetMsg(code),
				"data": data,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
