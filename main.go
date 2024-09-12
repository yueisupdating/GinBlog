package main

import (
	"ginblog/model"
	"ginblog/router"
	"ginblog/utils"
)

func main() {
	model.InitDb()
	utils.InitRedis()
	router.InitRouter()
}
