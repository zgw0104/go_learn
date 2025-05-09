package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
)

func main() {
	viper.SetDefault("fileDir", "./")

	//	读取配置文件
	viper.SetConfigName("config")         //配置文件名称(无扩展名)
	viper.SetConfigType("yaml")           //如果配置文件的名称中无扩展名，则需要配置此项
	viper.AddConfigPath("/etc/appname/")  //查找配置文件所在路径
	viper.AddConfigPath("$HOME/.appname") //多次调用以添加多个搜索路径
	viper.AddConfigPath(".")              //还可以在工作目录中查找配置

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	//实时监控配置文件变化
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		//配置文件发生变化会调用的回调函数
		fmt.Println("Config file changed:", e.Name)
	})

	r := gin.Default()
	r.GET("/version", func(c *gin.Context) {
		c.String(http.StatusOK, viper.GetString("version"))
	})

	r.Run()
}
