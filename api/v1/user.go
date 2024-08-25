package v1

import (
	"ginblog/model"
	"ginblog/utils/errmsg"
	"ginblog/utils/validator"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddUser(ctx *gin.Context) {
	var user model.User
	// 将HTTP请求报文中JSON格式的Body数据解析到结构体Struct或字典Map数据结构中
	_ = ctx.ShouldBindJSON(&user)
	msg, code := validator.Validate(user)
	if code != errmsg.SUCCESS {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": msg,
		})
		ctx.Abort()
		return
	}
	code = model.CheckUserName(user.UserName)

	if code == errmsg.SUCCESS {
		code = model.CreateUser(&user)
	}
	ctx.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

func DeleteUser(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	code := model.DeleteUser(id)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

func UpdateUser(ctx *gin.Context) {
	var user model.User
	id, _ := strconv.Atoi(ctx.Param("id"))
	_ = ctx.ShouldBindJSON(&user)

	code := model.CheckUserName(user.UserName)
	if code == errmsg.SUCCESS {
		model.EditUser(id, &user)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

func GetUser(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	user, code := model.GetUser(id)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"data":    user,
	})
}

func GetUsers(ctx *gin.Context) {
	pageSize, _ := strconv.Atoi(ctx.Query("pagesize"))
	pageNum, _ := strconv.Atoi(ctx.Query("pagenum"))

	if pageSize <= 0 {
		pageSize = 10
	}
	if pageNum == 0 {
		pageNum = 1
	}

	users, total := model.GetUsers(pageSize, pageNum)
	code := errmsg.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"data":    users,
		"total":   total,
	})
}
