package controller

import (
	"errors"

	"blog.com/dao/mysql"
	"github.com/go-playground/validator/v10"

	"go.uber.org/zap"

	"blog.com/models"

	"blog.com/logic"

	"github.com/gin-gonic/gin"
)

func SignUpHandler(c *gin.Context) {
	// 1. 进行用户参数校验
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("SignUp with invalid params", zap.Error(err))
		// 判断error 是不是validator.ValidationErrors
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			//c.JSON(http.StatusOK, gin.H{
			//	"message": err.Error(),
			//})
			return
		}
		//c.JSON(http.StatusOK, gin.H{
		//	"message": removeTopStruct(errs.Translate(trans)), // 翻译错误
		//})
		ResponseErrorWithMessage(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}

	// 2. 处理用户相关逻辑
	if err := logic.SignUp(p); err != nil {
		//c.JSON(http.StatusOK, gin.H{
		//	"message": "注册失败",
		//})
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3. 返回结果
	//c.JSON(http.StatusOK, gin.H{
	//	"message": "ok",
	//})
	ResponseSuccess(c, nil)
}

func LoginHandler(c *gin.Context) {
	// 1. 进行用户参数校验
	p := new(models.ParamLogin)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("Login with invalid params", zap.Error(err))
		// 判断error 是不是validator.ValidationErrors
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			//c.JSON(http.StatusOK, gin.H{
			//	"message": err.Error(),
			//})
			ResponseError(c, CodeInvalidParam)
			return
		}
		//c.JSON(http.StatusOK, gin.H{
		//	"message": removeTopStruct(errs.Translate(trans)), // 翻译错误
		//})
		ResponseErrorWithMessage(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	// 2.处理登录相关逻辑
	token, err := logic.Login(p)
	if err != nil {
		zap.L().Error("logic.login with invalid params", zap.Error(err))
		//c.JSON(http.StatusOK, gin.H{
		//	"message": "用户名或密码错误",
		//})
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotExist)
			return
		}
		ResponseError(c, CodeInvalidPassword)
		return
	}
	// 最后返回数据
	//c.JSON(http.StatusOK, gin.H{
	//	"message": "登录成功",
	//})
	ResponseSuccess(c, token)
}
