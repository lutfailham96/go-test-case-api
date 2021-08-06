package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Email     string    `gorm:"unique_index;not null" json:"email"`
	Username  string    `gorm:"unique_index;not null" json:"username"`
	Password  string    `gorm:"not null" json:"password"`
	Name      string    `gorm:"not null" json:"name"`
	Address   string    `gorm:"not null" json:"address"`
	AvatarUrl string    `json:"avatar_url"`
	Role      string    `gorm:"not null" json:"role"`
	Articles  []Article `gorm:"-" json:"articles"`
	Comments  []Comment `gorm:"-" json:"comments"`
}
