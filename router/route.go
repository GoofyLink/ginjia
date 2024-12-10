package router

import (
	"net/http"

	"blog.com/controller"

	"blog.com/logger"
	"github.com/gin-gonic/gin"
)

// SetupRouter 路由
func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "ok")
	})

	//注册业务路由
	r.POST("/signup", controller.SignUp)

	//没有找到页面返回这个路由
	r.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
