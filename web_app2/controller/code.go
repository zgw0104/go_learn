package controller

type ResCode int64

const (
	CodeSuccess = 1000 + iota
	CodeInvalidParam
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPwd
	CodeServerBusy
	CodeSignUpFailed
	CodeSignUpSuccess
	CodeLoginFailed
	CodeLoginSuccess

	CodeInvalidToken
	CodeNeedLogin
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:       "success",
	CodeInvalidParam:  "请求参数错误",
	CodeUserExist:     "用户已存在",
	CodeUserNotExist:  "用户不存在",
	CodeInvalidPwd:    "密码错误",
	CodeServerBusy:    "服务繁忙",
	CodeSignUpFailed:  "注册失败",
	CodeSignUpSuccess: "注册成功",
	CodeLoginFailed:   "登录失败",
	CodeLoginSuccess:  "登录成功",
	CodeInvalidToken:  "无效token",
	CodeNeedLogin:     "需要登录",
}

func (c ResCode) GetMsg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}
	return msg
}
