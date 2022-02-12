package main

import (
	"net/http"

	"goblog/app/http/middlewares"
	"goblog/pkg/logger"

	"goblog/bootstrap"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	bootstrap.SetupDB()
	router := bootstrap.SetupRoute()

	err := http.ListenAndServe(":3000", middlewares.RemoveTrailingSlash(router))
	logger.LogError(err)
}
