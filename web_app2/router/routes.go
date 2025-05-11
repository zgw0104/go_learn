package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web_app2/controller"
	"web_app2/logger"
)

func Setup() *gin.Engine {

	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	//注册业务路由
	r.POST("/signup", controller.SignUpHandler)

	r.POST("/signin", controller.SignInHandler)

	//r.GET("/ping", func(c *gin.Context) {
	//	if xx {
	//		c.String(http.StatusOK, "pong")
	//	} else {
	//		c.String(http.StatusOK, "请登录")
	//	}
	//})
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
