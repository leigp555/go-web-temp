package helper

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func HandleCors() gin.HandlerFunc {
	t := cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4500"},
		AllowMethods:     []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "AccessToken", "X-CSRF-Token", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:4500"
		},
		MaxAge: 12 * time.Hour,
	})
	return t
}
