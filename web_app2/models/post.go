package models

import "time"

// 内存对齐：
type Post struct {
	ID           int64     `json:"id,string" gorm:"id;AUTO_INCREMENT"`
	PostId       int64     `json:"post_id,string" gorm:"column:post_id"`
	AuthorId     int64     `json:"author_id,string" gorm:"column:author_id"`
	Community_id int64     `json:"community_id" gorm:"column:community_id" binding:"required"`
	Status       int32     `json:"status" gorm:"column:status"`
	PostTitle    string    `json:"post_title" gorm:"column:title" binding:"required"`
	Content      string    `json:"content" gorm:"column:content" binding:"required"`
	CreateTime   time.Time `json:"create_time" gorm:"column:create_time"`
	UpdateTime   time.Time `json:"update_time" gorm:"column:update_time"`
}

type ApiPostDetail struct {
	AuthorName string `json:"author_name" gorm:"column:author_name"`
	Score      int64  `json:"score" gorm:"column:score "`
	*Post
	*Community
}
