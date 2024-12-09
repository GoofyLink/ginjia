package router

import (
	"net/http"

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
	return r
}
