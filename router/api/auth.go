package api

import (
	"Gin_demo/models"
	cs "Gin_demo/pkg/constant"
	"Gin_demo/pkg/util"
	"log"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

//获取Token
func GetAuth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	valid := validation.Validation{}
	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)

	data := make(map[string]interface{})
	code := cs.INVALID_PARAMS

	if ok {
		//判断用户名、密码是否查询一致
		isExist := models.CheckAuth(username, password)
		if isExist {
			//存在，生成token
			token, err := util.GenerateToken(username, password)
			if err != nil {
				//token生成失败
				code = cs.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token
				code = cs.SUCCESS
			}
		} else {
			//用户名、密码不相符，无法生成token
			code = cs.ERROR_AUTH
		}
	} else {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  cs.GetMsg(code),
		"data": data,
	})
}
