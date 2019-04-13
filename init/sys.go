package init

import (
	"Gin_demo/models"
	"Gin_demo/pkg/util"
)

func init() {
	//初始化ini配置
	util.IniInit()
	//初始化数据库配置
	models.DbInit()
}
