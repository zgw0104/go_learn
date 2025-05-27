package controller

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"web_app2/dao/mysql"
	"web_app2/logic"
	"web_app2/models"
)

//Controller层：负责处理路由，参数校验，请求转发

func SignUpHandler(c *gin.Context) {
	//1 获取参数和参数校验
	p := new(models.ParamSighUp)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("SignUp with invalid param", zap.Error(err))

		//判断err是不是validator类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			Response(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}

	//2 业务处理
	if err := logic.SignUp(p); err != nil {
		zap.L().Error("SignUp err:", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist) {
			Response(c, CodeUserExist)
			return
		}
		Response(c, CodeSignUpFailed)
		return
	}
	//3 返回响应
	Response(c, CodeSignUpSuccess)
}

func SignInHandler(c *gin.Context) {
	//1 获取参数和参数校验
	p := new(models.ParamSignIn)

	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("SignIn with invalid param", zap.Error(err))

		//判断err是不是validator类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			Response(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}

	//2 业务处理
	user, err := logic.SignIn(p)
	if err != nil {
		zap.L().Error("SignIn err", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExist) {
			Response(c, CodeUserNotExist)
		}
		Response(c, CodeInvalidPwd)
		return
	}

	//3 返回响应
	ResponseSuccess(c, gin.H{
		"user_id":   fmt.Sprintf("%d", user.UserId),
		"user_name": user.UserName,
		"token":     user.Atoken,
	})
}
