package models

import "gorm.io/gorm"
import "gorm.io/driver/mysql"

var DB *gorm.DB

func NewDB(dsn string) error {
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		return err
	}

	DB = db
	err = DB.AutoMigrate(&UserBasic{})
	if err != nil {
		return err
	}
	return nil
}
