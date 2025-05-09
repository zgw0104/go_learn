package main

import (
	"context"
	_ "database/sql"
	"fmt"
	_ "github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"web_app/dao/mysql"
	"web_app/dao/redis"
	"web_app/logger"
	"web_app/routes"
	"web_app/settings"
)

//go web 开发脚手架模板

func main() {
	//1.加载配置文件
	if err := settings.Init(); err != nil {
		fmt.Println("init settings err:", err)
		return
	}
	//2.初始化日志
	if err := logger.Init(); err != nil {
		fmt.Println("init logger err:", err)
		return
	}
	defer zap.L().Sync()
	zap.L().Debug("logger init success...")

	//3.初始化mysql
	if err := mysql.Init(); err != nil {
		fmt.Println("init mysql err:", err)
		return
	}

	//4.初始化redis
	if err := redis.Init(); err != nil {
		fmt.Println("init redis err:", err)
		return
	}
	defer redis.Close()
	//5.注册路由
	r := routes.Setup()

	//6.启动服务
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", viper.GetInt("app.port")),
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	zap.L().Info("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server Shutdown:", zap.Error(err))
	}
	zap.L().Info("Server exiting")
}
