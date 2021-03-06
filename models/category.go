package models

import "gorm.io/gorm"

//User model
type Category struct {
	gorm.Model
	Name        string `gorm:"not null" json:"name"`
	Description string `gorm:"not null" json:"description"`
	Courses     []Course
}
