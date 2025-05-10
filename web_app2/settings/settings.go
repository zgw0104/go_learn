package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	_ "github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var Conf = new(AppConfig)

type AppConfig struct {
	Name         string `mapstructure:"name"`
	Mode         string `mapstructure:"mode"`
	Version      string `mapstructure:"version"`
	Port         int    `mapstructure:"port"`
	StartTime    string `mapstructure:"start_time"`
	MachineID    int64  `mapstructure:"machine_id"`
	*LogConfig   `mapstructure:"log"`
	*MySqlConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"maxsize"`
	MaxBackups int    `mapstructure:"maxbackups"`
	MaxAge     int    `mapstructure:"maxage"`
}

type MySqlConfig struct {
	Host    string `mapstructure:"host"`
	Port    int    `mapstructure:"port"`
	User    string `mapstructure:"user"`
	Pwd     string `mapstructure:"pwd"`
	DbName  string `mapstructure:"dbname"`
	MaxIdle int    `mapstructure:"max_idle"`
	MaxOpen int    `mapstructure:"max_open"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Pwd      string `mapstructure:"pwd"`
	PoolSize int    `mapstructure:"pool_size"`
	DB       int    `mapstructure:"db"`
}

func Init() (err error) {
	viper.SetConfigFile("config.yaml")
	//viper.SetConfigFile("config")
	//viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err = viper.ReadInConfig()
	if err != nil {
		return err
	}

	//读取到的配置信息反序列化
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Println("viper.Unmarshal err:", err.Error())
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Println("viper.Unmarshal err:", err.Error())
		}
	})
	return
	//r := gin.Default()
	//if err := r.Run(fmt.Sprintf(":%d", viper.Get("Port"))); err != nil {
	//	panic(err)
	//}
}
