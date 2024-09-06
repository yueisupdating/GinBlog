package v1

import (
	"ginblog/model"
	"ginblog/utils/errmsg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddArt(ctx *gin.Context) {
	var article model.Article
	_ = ctx.ShouldBindJSON(&article)
	code := model.CreateArticle(&article)

	ctx.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    article,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

func DeleteArticle(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	code := model.DeleteArt(id)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

func UpdateArticle(ctx *gin.Context) {
	var article model.Article
	id, _ := strconv.Atoi(ctx.Param("id"))
	_ = ctx.ShouldBindJSON(&article)

	code := model.EditArt(id, &article)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

func GetArticles(ctx *gin.Context) {
	title := ctx.Query("title")
	articles, total := model.GetArticles(title)
	code := errmsg.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"data":    articles,
		"total":   total,
	})
}

func GetArticleByCategory(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	articles, code, cnt := model.GetArticleByCategory(id)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"data":    articles,
		"total":   cnt,
	})
}

func GetArticleByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	user, code := model.GetArticleByID(id)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"data":    user,
	})
}
