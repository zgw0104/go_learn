package models

//定义请求参数的结构体

// 用户注册请求参数
type ParamSighUp struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	RePasswd string `json:"re_passwd" binding:"required,eqfield=Password"`
}

// 用户登录请求参数
type ParamSignIn struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
