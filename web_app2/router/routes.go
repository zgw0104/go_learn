package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web_app2/controller"
	"web_app2/logger"
	"web_app2/middleware"
)

func Setup() *gin.Engine {

	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	v1 := r.Group("api/v1")

	//注册业务路由
	v1.POST("/signup", controller.SignUpHandler)

	v1.POST("/signin", controller.SignInHandler)

	v1.Use(middleware.JWTAuthMiddleware())

	{
		v1.GET("/community", controller.CommunityHandler)
		v1.GET("/community/:id", controller.CommunityDetailHandler)
		v1.POST("/post", controller.CreatePostHandler)
		v1.GET("/post/:id", controller.GetPostDetailHandler)
		v1.GET("/posts/", controller.GetPostListHandler)
		v1.GET("/posts2/", controller.GetPostListHandler2)
		v1.GET("/posts3", controller.GetCommunityPostListHandler)

		v1.POST("/vote", controller.PostVoteHandler)

	}

	v1.GET("/ping", middleware.JWTAuthMiddleware(), func(c *gin.Context) {
		// 如果是登录用户，判断请求头中是否包含token
		c.String(http.StatusOK, "pong")
	})

	v1.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
