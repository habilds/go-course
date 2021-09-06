package models

import "gorm.io/gorm"

//User model
type User struct {
	gorm.Model
	Username string `gorm:"unique_index;not null" json:"username"`
	Email    string `gorm:"unique_index;not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
	Name     string `json:"name"`
}
