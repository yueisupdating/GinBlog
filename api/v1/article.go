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
			"status": code,
			"data":   article,
			"msg":    errmsg.GetErrMsg(code),
		},
	)
}

func DeleteArticle(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	code := model.DeleteArt(id)
	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errmsg.GetErrMsg(code),
	})
}

func UpdateArticle(ctx *gin.Context) {
	var article model.Article
	id, _ := strconv.Atoi(ctx.Param("id"))
	_ = ctx.ShouldBindJSON(&article)

	code := model.EditArt(id, &article)
	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errmsg.GetErrMsg(code),
	})
}

func GetArticles(ctx *gin.Context) {
	pageSize, _ := strconv.Atoi(ctx.Query("pagesize"))
	pageNum, _ := strconv.Atoi(ctx.Query("pagenum"))

	if pageSize <= 0 {
		pageSize = 10
	}
	if pageNum == 0 {
		pageNum = 1
	}

	articles, code := model.GetArticles(pageSize, pageNum)
	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errmsg.GetErrMsg(code),
		"data":   articles,
	})
}
