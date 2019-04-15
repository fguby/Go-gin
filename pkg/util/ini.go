package util

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-ini/ini"
)

var cfg *ini.File

type Server struct {
	App          string
	Port         int
	RunMode      string
	ReadTimeOut  time.Duration
	WriteTimeOut time.Duration
	JwtSecret    string
}

var ServerSetting = &Server{}

func IniInit() {
	var err error
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		fmt.Println("加载文件失败: %s", err)
		//非正常运行，退出整个程序
		os.Exit(1)
	}
	mapTo("server", ServerSetting)
	//换算成秒数
	ServerSetting.ReadTimeOut = ServerSetting.ReadTimeOut * time.Second
	ServerSetting.WriteTimeOut = ServerSetting.ReadTimeOut * time.Second
}

//根据配置文件获取属性值
func GetIniValue(mode string, key string) string {
	value := cfg.Section(mode).Key(key).String()
	return value
}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo RedisSetting err: %v", err)
	}
}
