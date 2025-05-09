package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	_ "github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Init() (err error) {
	viper.SetConfigFile("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err = viper.ReadInConfig()
	if err != nil {
		return err
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
	return
	//r := gin.Default()
	//if err := r.Run(fmt.Sprintf(":%d", viper.Get("Port"))); err != nil {
	//	panic(err)
	//}
}
