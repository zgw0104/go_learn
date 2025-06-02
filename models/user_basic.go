package models

import "gorm.io/gorm"

type UserBasic struct {
	gorm.Model
	Identity string `gorm:"column:identity;type:varchar(36);" json:"identity"` //唯一标识
	Username string `gorm:"column:username;type:varchar(255);" json:"username"`
	Password string `gorm:"column:password;type:varchar(255);" json:"password"`
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}
