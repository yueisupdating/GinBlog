package router

import (
	v1 "ginblog/api/v1"
	"ginblog/middleware"
	"ginblog/utils"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	engine := gin.Default()

	authV1 := engine.Group("api/v1")
	authV1.Use(middleware.JwtToken())
	{
		// 用户路由接口
		authV1.DELETE("user/delete/:id", v1.DeleteUser) // 测试时将:id改为具体的id
		authV1.PUT("user/update/:id", v1.UpdateUser)

		// 分类路由接口
		authV1.POST("category/add", v1.AddCate)
		authV1.DELETE("category/delete/:id", v1.DeleteCate) // 测试时将:id改为具体的id
		authV1.PUT("category/update/:id", v1.UpdateCate)

		// 文章路由接口
		authV1.POST("article/add", v1.AddArt)
		authV1.DELETE("article/delete/:id", v1.DeleteArticle) // 测试时将:id改为具体的id
		authV1.PUT("article/update/:id", v1.UpdateArticle)

		authV1.POST("upload", v1.UpLoad)
	}
	routerV1 := engine.Group("api/v1")
	{
		routerV1.POST("user/add", v1.AddUser)

		routerV1.GET("users", v1.GetUsers)
		routerV1.GET("categorys", v1.GetCates)
		routerV1.GET("articles", v1.GetArticles) // 所有文章列表查询
		routerV1.GET("article/category/:id", v1.GetArticleByCategory)
		// 登录路由接口
		routerV1.POST("login", v1.Login)
	}
	engine.Run(utils.HttpPort)
}
