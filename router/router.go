package router

import (
	v1 "ginblog/api/v1"
	"ginblog/middleware"
	"ginblog/utils"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(middleware.Log())
	engine.Use(middleware.Cors())

	authV1 := engine.Group("api/v1")
	authV1.Use(middleware.JwtToken())
	{
		// 用户路由接口
		authV1.DELETE("admin/user/delete/:id", v1.DeleteUser) // 测试时将:id改为具体的id
		authV1.PUT("admin/user/update/:id", v1.UpdateUser)

		// 分类路由接口
		authV1.POST("admin/cate/add", v1.AddCate)
		authV1.DELETE("admin/cate/delete/:id", v1.DeleteCate) // 测试时将:id改为具体的id
		authV1.PUT("admin/cate/update/:id", v1.UpdateCate)

		// 文章路由接口
		authV1.POST("article/add", v1.AddArt)
		authV1.DELETE("article/delete/:id", v1.DeleteArticle) // 测试时将:id改为具体的id
		authV1.PUT("article/update/:id", v1.UpdateArticle)

		authV1.POST("upload", v1.UpLoad)
	}
	routerV1 := engine.Group("api/v1")
	{
		routerV1.POST("admin/user/add", v1.AddUser)

		routerV1.GET("admin/user/get/:id", v1.GetUser)
		routerV1.GET("admin/userList", v1.GetUsers)
		routerV1.GET("admin/cate/get/:id", v1.GetCate)
		routerV1.GET("admin/cateList", v1.GetCates)
		routerV1.GET("articles", v1.GetArticles) // 所有文章列表查询
		routerV1.GET("article/category/:id", v1.GetArticleByCategory)
		// 登录路由接口
		routerV1.POST("login", v1.Login)
	}
	engine.Run(utils.HttpPort)
}
