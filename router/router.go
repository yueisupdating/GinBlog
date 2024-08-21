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
	routerV1.DELETE("user/delete/:id", v1.DeleteUser) // 测试时将:id改为具体的id
	routerV1.PUT("user/update/:id", v1.UpdateUser)
	routerV1.GET("users", v1.GetUsers)
	// 分类路由接口
	routerV1.POST("category/add", v1.AddCate)
	routerV1.DELETE("category/delete/:id", v1.DeleteCate) // 测试时将:id改为具体的id
	routerV1.PUT("category/update/:id", v1.UpdateCate)
	routerV1.GET("categorys", v1.GetCates)
	// 文章路由接口
	routerV1.POST("article/add", v1.AddArt)
	routerV1.DELETE("article/delete/:id", v1.DeleteArticle) // 测试时将:id改为具体的id
	routerV1.PUT("article/update/:id", v1.UpdateArticle)
	routerV1.GET("articles", v1.GetArticles) // 所有文章列表查询
	routerV1.GET("article/category/:id", v1.GetArticleByCategory)

	engine.Run(utils.HttpPort)
}
