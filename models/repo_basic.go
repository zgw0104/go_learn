package models

import "gorm.io/gorm"

type RepoBasic struct {
	gorm.Model
	Uid      int    `gorm:"column:uid" json:"uid"`
	Identity string `gorm:"column:identity;type:varchar(36);" json:"identity"` //唯一标识
	Name     string `gorm:"column:name;type:varchar(255);" json:"name"`        // name
	Desc     string `gorm:"column:desc;type:varchar(255);" json:"desc"`        // desc
	Star     int    `gorm:"column:star;type:int(11);" json:"star"`             //star
}

func (table *RepoBasic) TableName() string {
	return "repo_basic"
}
