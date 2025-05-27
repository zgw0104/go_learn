package models

import "time"

type User struct {
	Id         int64     `gorm:"primary_key;column:id"`
	UserId     int64     `gorm:"column:user_id"`
	UserName   string    `gorm:"column:username"`
	Password   string    `gorm:"column:password"`
	Email      string    `gorm:"column:email"`
	Gender     int       `gorm:"column:gender"`
	CreateTime time.Time `gorm:"column:create_time"`
	UpdateTime time.Time `gorm:"column:update_time"`
	Atoken     string    `json:"atoken"`
	Rtoken     string    `json:"rtoken"`
}
