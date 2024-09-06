package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"*"},
		AllowHeaders:  []string{"Origin", "Authorization", "Content-Type", "X-Requested-With"},
		ExposeHeaders: []string{"Content-Length", "Authorization"},
		// AllowCredentials: true, // 是否允许传cookie
		MaxAge: 12 * time.Hour,
	})
}
