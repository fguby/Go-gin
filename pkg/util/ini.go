package util

import (
	"fmt"
	"os"

	"github.com/go-ini/ini"
)

var cfg, ini_err = ini.Load("conf/app.ini")

func IniInit() {
	if ini_err != nil {
		fmt.Println("加载文件失败: %s", ini_err)
		//非正常运行，退出整个程序
		os.Exit(1)
	}
}

//根据配置文件获取属性值
func GetIniValue(mode string, key string) string {
	value := cfg.Section(mode).Key(key).String()
	return value
}
