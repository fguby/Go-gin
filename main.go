package main

import (
	_ "Gin_demo/init"
	"Gin_demo/pkg/util"
	"fmt"
	"net/http"

	"Gin_demo/router"
)

func main() {
	router := router.InitRouter()

	fmt.Println("端口为:%s", util.ServerSetting.Port)

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", util.ServerSetting.Port),
		Handler:        router,
		ReadTimeout:    util.ServerSetting.ReadTimeOut,
		WriteTimeout:   util.ServerSetting.WriteTimeOut,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
