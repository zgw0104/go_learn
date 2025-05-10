package models

//定义请求参数的结构体

type ParamSighUp struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	RePasswd string `json:"re_passwd" binding:"required"`
}
