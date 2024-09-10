package v1

import (
	"ginblog/model"
	"ginblog/utils/errmsg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetProfile(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	data, code := model.GetProfile(id)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

func UpdateProfile(ctx *gin.Context) {
	var data *model.Profile
	id, _ := strconv.Atoi(ctx.Param("id"))
	_ = ctx.ShouldBindJSON(&data)
	data, code := model.UpdateProfile(data, id)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
