package models

import (
	"Gin_demo/pkg/util"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
	e  = func(err error) {
		if err != nil {
			panic(err)
		}
	}
	fm = fmt.Sprintf
)

type Base_model struct {
	Id         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modifyed_on"`
	DeletedOn  int `json:"deleted_on"`
}

//数据库连接初始化
func DbInit() {
	var err error
	host := util.GetIniValue("mysql", "host")
	user := util.GetIniValue("mysql", "user")
	pd := util.GetIniValue("mysql", "password")
	dbName := util.GetIniValue("mysql", "db")
	db, err = gorm.Open("mysql", fm("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		pd,
		host,
		dbName,
	))
	e(err)
	//设置db属性
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	//添加表名前缀
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "blog_" + defaultTableName
	}
}

//关闭数据库
func CloseDB() {
	defer db.Close()
}
