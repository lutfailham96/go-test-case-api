package model

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	CommentText string         `gorm:"not null" json:"comment_text"`
	ArticleID   uint           `gorm:"not null" json:"article_id"`
	UserID      uint           `gorm:"not null" json:"user_id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
