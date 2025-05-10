package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"web_app2/logic"
	"web_app2/models"
)

//Controller层：负责处理路由，参数校验，请求转发

func SignUpHandler(c *gin.Context) {
	//1 获取参数和参数校验
	p := new(models.ParamSighUp)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("SignUp with  invalid param", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"请求参数有误": err.Error()})
		return
	}
	//手动对请求参数进行详细的业务规则校验
	//if len(p.Password) == 8 || len(p.Username) == 0 || len(p.RePasswd) == 0 || p.Password != p.RePasswd {
	//	zap.L().Error("SignUp with  invalid param")
	//	c.JSON(http.StatusBadRequest, gin.H{"msg": "请求参数有误"})
	//	return
	//}
	//2 业务处理
	logic.SignUp(p)
	//3 返回响应
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
