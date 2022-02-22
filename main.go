package main

import (
	"chris_project/models"
	"chris_project/router"
)

func init() {
	//初始化DB
	models.Setup()
}

func main() {
	//初始化路由
	r := router.SetupRouter()

	r.Run()
}
