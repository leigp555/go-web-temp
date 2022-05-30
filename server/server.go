package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web/helper"
	"web/middleware"
)

func HandleLogin(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, map[string]string{
		"msg": "完美",
	})
}
func HandleHome(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, map[string]string{
		"msg": "完美",
	})
}

// Serve 创建web服务
func Serve() (*gin.Engine, error) {
	app := gin.Default()
	//允许跨域
	app.Use(helper.HandleCors())

	app.POST("/login", middleware.GenerateToken, HandleLogin)
	app.GET("/home", middleware.ParseToken, HandleHome)

	//设置端口
	err := app.Run(":8000")
	if err != nil {
		return nil, err
	}
	return app, nil
}
