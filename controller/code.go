package controller

type ResponseCode int64

const (
	CodeSuccess ResponseCode = 1000 + iota
	CodeInvalidParam
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword
	CodeServerBusy
	CodeSignUpFailed
	CodeSignSuccess
	CodeLoginSuccess
	CodeNeedLogin
	CodeInvalidToken
)

var codeMessageMap = map[ResponseCode]string{
	CodeSuccess:         "success",
	CodeInvalidParam:    "请求参数错误",
	CodeUserExist:       "用户已存在",
	CodeUserNotExist:    "用户不存在",
	CodeInvalidPassword: "用户名或密码错误",
	CodeServerBusy:      "服务繁忙",
	CodeSignUpFailed:    "注册失败",
	CodeSignSuccess:     "注册成功",
	CodeLoginSuccess:    "登录成功",
	CodeNeedLogin:       "需要登录",
	CodeInvalidToken:    "无效的token",
}

func (c ResponseCode) Message() string {
	message, ok := codeMessageMap[c]
	if !ok {
		message = codeMessageMap[CodeServerBusy]
	}
	return message
}
