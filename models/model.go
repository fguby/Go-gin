package models

import (
	"Gin_demo/pkg/util"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

//数据库连接初始化
func DbInit() {
	var err error
	host := util.GetIniValue("mysql", "host")
	user := util.GetIniValue("mysql", "user")
	pd := util.GetIniValue("mysql", "password")
	dbName := util.GetIniValue("mysql", "db")
	db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		pd,
		host,
		dbName,
	))
	if err != nil {
		fmt.Println("数据库初始化连接失败: %s", err)
		//退出程序
		os.Exit(1)
	}
	//设置db属性
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}
