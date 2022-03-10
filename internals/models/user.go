package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FullName   string    `json:"full_name"`
	Email      string    `json:"email" gorm:"uniqueIndex"`
	UserName   string    `json:"username" gorm:"uniqueIndex"`
	Password   string    `json:"password"`
	IsVerified bool      `json:"is_verified"`
	Contact    []Contact `json:"contact"`
}
