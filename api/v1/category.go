package v1

import (
	"ginblog/model"
	"ginblog/utils/errmsg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddCate(ctx *gin.Context) {
	var cate model.Category
	_ = ctx.ShouldBindJSON(&cate)
	code := model.CheckCateName(cate.CategoryName)
	if code == errmsg.SUCCESS {
		code = model.CreateCategory(&cate)
	}
	ctx.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    cate,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

func DeleteCate(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	code := model.DeleteCate(id)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

func UpdateCate(ctx *gin.Context) {
	var cate model.Category
	id, _ := strconv.Atoi(ctx.Param("id"))
	_ = ctx.ShouldBindJSON(&cate)

	code := model.CheckCateName(cate.CategoryName)
	if code == errmsg.SUCCESS {
		model.EditCate(id, &cate)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

func GetCates(ctx *gin.Context) {
	pageSize, _ := strconv.Atoi(ctx.Query("pagesize"))
	pageNum, _ := strconv.Atoi(ctx.Query("pagenum"))

	if pageSize <= 0 {
		pageSize = 10
	}
	if pageNum == 0 {
		pageNum = 1
	}

	cates, code := model.GetCates(pageSize, pageNum)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"data":    cates,
	})
}
