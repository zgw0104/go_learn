package models

import "gorm.io/gorm"

type RepoUser struct {
	gorm.Model
	Rid  int `gorm:"column:rid;type:int(11);" json:"r_id"` //仓库id
	Uid  int `gorm:"column:uid;type:int(11)" json:"uid"`
	Type int `gorm:"column:type;type:tinyint(1)" json:"type"` //类型，{1：所有者 2：被授权}
}

func (table *RepoUser) TableName() string {
	return "repo_user"
}
