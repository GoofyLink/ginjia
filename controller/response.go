package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
{
   "code":10000,  // 程序中的错误码
   "message": xx // 提示信息
   "data": {}    // 数据
}
*/

type ResponseData struct {
	Code    ResponseCode `json:"code"`
	Message interface{}  `json:"message"`
	Data    interface{}  `json:"data"`
}

func ResponseError(c *gin.Context, code ResponseCode) {
	rd := &ResponseData{
		Code:    code,
		Message: code.Message(),
		Data:    nil,
	}
	c.JSON(http.StatusOK, rd)
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	rd := &ResponseData{
		Code:    CodeSuccess,
		Message: CodeSuccess.Message(),
		Data:    data,
	}
	c.JSON(http.StatusOK, rd)
}

// ResponseErrorWithMessage 自定一错误信息和code
func ResponseErrorWithMessage(c *gin.Context, code ResponseCode, message interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		Code:    code,
		Message: message,
		Data:    nil,
	})
}
