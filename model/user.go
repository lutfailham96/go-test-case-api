package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	Email     string         `gorm:"unique_index;not null" json:"email"`
	Username  string         `gorm:"unique_index;not null" json:"username"`
	Password  string         `gorm:"not null" json:"password"`
	Name      string         `gorm:"not null" json:"name"`
	Address   string         `gorm:"not null" json:"address"`
	AvatarUrl string         `json:"avatar_url"`
	Role      string         `gorm:"not null" json:"role"`
	Articles  []Article      `json:"articles"`
	Comments  []Comment      `json:"comments"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
