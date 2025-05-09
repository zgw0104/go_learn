package mysql

import (
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var sqldb *gorm.DB

func Init() (err error) {
	//dsn := "root:zgw64220392@tcp(127.0.0.1:3306)/go?charset=utf8"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		viper.GetString("mysql.user"),
		viper.GetString("mysql.pwd"),
		viper.GetString("mysql.host"),
		viper.GetInt("mysql.port"),
		viper.GetString("mysql.dbname"),
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("init db err:", err)
	}

	sqldb, err := db.DB()
	if err != nil {
		zap.L().Error("sqlDB err:", zap.Error(err))
	}
	sqldb.SetMaxIdleConns(viper.GetInt("mysql.max_idle"))
	sqldb.SetMaxOpenConns(viper.GetInt("mysql.max_open"))
	sqldb.SetConnMaxLifetime(time.Hour)
	return

}
