package main

import (
	"context"
	_ "database/sql"
	"fmt"
	_ "github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"web_app2/dao/mysql"
	"web_app2/dao/redis"
	"web_app2/logger"
	"web_app2/pkg/snowflake"
	"web_app2/router"
	"web_app2/settings"
)

//go web 开发脚手架模板

func main() {
	//1.加载配置文件
	if err := settings.Init(); err != nil {
		fmt.Println("init settings err:", err)
		return
	}
	//2.初始化日志
	if err := logger.Init(settings.Conf.LogConfig); err != nil {
		fmt.Println("init logger err:", err)
		return
	}
	defer zap.L().Sync()
	zap.L().Debug("logger init success...")

	//3.初始化mysql
	if err := mysql.Init(settings.Conf.MySqlConfig); err != nil {
		fmt.Println("init mysql err:", err)
		return
	}

	//4.初始化redis
	if err := redis.Init(settings.Conf.RedisConfig); err != nil {
		fmt.Println("init redis err:", err)
		return
	}
	defer redis.Close()

	//初始化雪花id
	if err := snowflake.Init(settings.Conf.StartTime, settings.Conf.MachineID); err != nil {
		fmt.Println("init snowflake err:", err)
		return
	}

	//5.注册路由
	r := router.Setup()

	//6.启动服务
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", settings.Conf.Port),
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
