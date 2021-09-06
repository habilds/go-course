package models

import "gorm.io/gorm"

//User model
type Course struct {
	gorm.Model
	Title       string   `gorm:"not null" json:"title"`
	Description string   `gorm:"not null" json:"description"`
	Price       int      `gorm:"not null" json:"price"`
	CategoryID  int      `json:"category_id"`
	Category    Category `json:"category"`
}
