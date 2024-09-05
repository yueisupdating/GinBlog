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

	code := model.CheckCateNameForUpdate(cate.CategoryName, int(cate.ID))
	if code == errmsg.SUCCESS {
		model.EditCate(id, &cate)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

func GetCate(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	cate, code := model.GetCate(id)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"data":    cate,
	})
}

func GetCates(ctx *gin.Context) {
	cates, total := model.GetCates()
	code := errmsg.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"data":    cates,
		"total":   total,
	})
}
