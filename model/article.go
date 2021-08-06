package model

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Title            string    `gorm:"not null" json:"title"`
	Content          string    `gorm:"not null" json:"content"`
	FeaturedImageUrl string    `gorm:"not null" json:"featured_image_url"`
	Comments         []Comment `gorm:"-" json:"comments"`
}
