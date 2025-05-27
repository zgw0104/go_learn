package models

import "time"

type Community struct {
	CommunityId   int64  `gorm:"column:community_id" json:"CommunityId"`
	CommunityName string `gorm:"column:community_name" json:"CommunityName"`
	Introduction  string `gorm:"column:introduction" json:"Introduction"`
}

type CommunityDetail struct {
	CommunityId   int64     `gorm:"column:community_id"`
	CommunityName string    `gorm:"column:community_name"`
	Introduction  string    `gorm:"column:introduction" json:"introduction,omitempty"`
	CreateTime    time.Time `gorm:"column:create_time" json:"create_time,omitempty"`
}
