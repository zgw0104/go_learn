package mysql

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
	"web_app2/settings"
)

var db *gorm.DB

func Init(cfg *settings.MySqlConfig) (err error) {
	//dsn := "root:zgw64220392@tcp(127.0.0.1:3306)/go?charset=utf8"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true",
		cfg.User,
		cfg.Pwd,
		cfg.Host,
		cfg.Port,
		cfg.DbName,
	)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("init db err:", err)
	}

	sqldb, err := db.DB()
	if err != nil {
		zap.L().Error("sqlDB err:", zap.Error(err))
	}
	sqldb.SetMaxIdleConns(cfg.MaxIdle)
	sqldb.SetMaxOpenConns(cfg.MaxOpen)
	sqldb.SetConnMaxLifetime(time.Hour)
	return

}
