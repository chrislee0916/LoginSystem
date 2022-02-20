package main

import (
	"chris_project/models"
	"chris_project/router"
)

func init() {
	models.Setup()
}

func main() {

	r := router.SetupRouter()

	r.Run()
}
