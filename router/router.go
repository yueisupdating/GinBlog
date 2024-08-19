package router

import (
	v1 "ginblog/api/v1"
	"ginblog/utils"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	engine := gin.Default()
	routerV1 := engine.Group("api/v1")

	// 用户路由接口
	routerV1.POST("user/add", v1.AddUser)
	//TODO:删
	//TODO:改
	routerV1.GET("users", v1.GetUsers)

	engine.Run(utils.HttpPort)
}
