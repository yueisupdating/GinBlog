package v1

import (
	"ginblog/middleware"
	"ginblog/model"
	"ginblog/utils/errmsg"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	var user model.User
	var token string
	var err error
	_ = ctx.ShouldBindJSON(&user)
	user, code := model.CheckLogin(user.UserName, user.Password)
	if code == errmsg.SUCCESS {
		jwt := middleware.NewJWT()
		token, err = jwt.GenerateJWT(user.UserName)
		if err != nil {
			code = errmsg.ERROR
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": errmsg.GetErrMsg(code),
		"token":   token,
	})
}
