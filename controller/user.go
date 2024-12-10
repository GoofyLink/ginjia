package controller

import (
	"github.com/go-playground/validator/v10"
	"net/http"

	"go.uber.org/zap"

	"blog.com/models"

	"blog.com/logic"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	// 1. 进行用户参数校验
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("SignUp with invalid params", zap.Error(err))
		// 判断error 是不是validator.ValidationErrors
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": removeTopStruct(errs.Translate(trans)),
		})
		return
	}
	
	// 2. 处理用户相关逻辑
	if err := logic.SignUp(p); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "注册失败",
		})
		return
	}
	// 3. 返回结果
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
