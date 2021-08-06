package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	CommentText string `gorm:"not null" json:"comment_text"`
	ArticleID   int    `json:"article_id"`
	UserID      int    `json:"user_id"`
}
