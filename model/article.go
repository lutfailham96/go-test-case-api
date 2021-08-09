package model

import (
	"time"

	"gorm.io/gorm"
)

type Article struct {
	ID               uint           `gorm:"primarykey" json:"id"`
	Title            string         `gorm:"not null" json:"title"`
	Content          string         `gorm:"not null" json:"content"`
	FeaturedImageUrl string         `gorm:"not null" json:"featured_image_url"`
	UserID           uint           `gorm:"not null" json:"user_id"`
	Comments         []Comment      `json:"comments"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
