package middleware

import (
	"errors"
	"ginblog/utils"
	"ginblog/utils/errmsg"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	JwtKey []byte
}

func NewJWT() *JWT {
	return &JWT{[]byte(utils.JWTKEY)}
}

type MyClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// 生成JWT
func (j *JWT) GenerateJWT(username string) (string, error) {
	claims := MyClaims{
		username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // 过期时间24小时
			IssuedAt:  jwt.NewNumericDate(time.Now()),                     // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                     // 生效时间
		},
	}
	// 使用HS256签名算法
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.JwtKey)
}

// 解析JWT
func (j *JWT) ParseJwt(tokenString string) (*MyClaims, int) {
	var jwtClaim = &MyClaims{}
	token, err := jwt.ParseWithClaims(tokenString, jwtClaim, func(token *jwt.Token) (interface{}, error) {
		return j.JwtKey, nil
	})

	if token.Valid {
		return jwtClaim, errmsg.SUCCESS
	} else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
		return nil, errmsg.ERROR_TOKEN_RUNTIME
	} else {
		return nil, errmsg.ERROR_TOKEN_WRONG
	}

}

func JwtToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var code int
		tokenHeader := ctx.Request.Header.Get("Authorization")
		if tokenHeader == "" {
			code = errmsg.ERROR_TOKEN_NOT_EXIST
			ctx.JSON(http.StatusOK, gin.H{
				"status":  code,
				"message": errmsg.GetErrMsg(code),
			})
			ctx.Abort()
			return
		}

		tokenSplit := strings.Split(tokenHeader, " ")
		if len(tokenSplit) != 2 || tokenSplit[0] != "Bearer" {
			code = errmsg.ERROR_TOKEN_TYPE_WRONG
			ctx.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			ctx.Abort()
			return
		}
		jwt := NewJWT()
		_, code = jwt.ParseJwt(tokenSplit[1])

		if code != errmsg.SUCCESS {
			if code == errmsg.ERROR_TOKEN_RUNTIME {
				ctx.JSON(http.StatusOK, gin.H{
					"status":  code,
					"message": errmsg.GetErrMsg(code),
					"data":    nil,
				})
				ctx.Abort()
				return
			}
			ctx.JSON(http.StatusOK, gin.H{
				"status":  code,
				"message": errmsg.GetErrMsg(code),
				"data":    nil,
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
