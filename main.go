package main

import (
	"fmt"
	"goblog/app/http/middlewares"
	"goblog/bootstrap"

	"goblog/pkg/config"
	"net/http"
)

func init() {
	// 初始化配置信息
	// config.Initialize()
}

func main() {
	a := config.StrMap{}
	fmt.Println(a)
	// 初始化 SQL
	bootstrap.SetupDB()

	// 初始化路由绑定
	router := bootstrap.SetupRoute()

	http.ListenAndServe(":3000", middlewares.RemoveTrailingSlash(router))
}
