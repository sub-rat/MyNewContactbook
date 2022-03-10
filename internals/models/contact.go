package models

import "gorm.io/gorm"

type Phone struct {
	gorm.Model
	PhoneType   string `json:"type"`
	PhoneNumber string `json:"phone_number"`
	ContactId   uint
}

type Address struct {
	gorm.Model
	Country   string `json:"country"`
	City      string `json:"city"`
	Street    string `json:"street"`
	ContactID uint
}

type Contact struct {
	gorm.Model `json:"id"`
	FirstName  string     `json:"first_name"`
	LastName   string     `json:"last_name"`
	Email      string     `json:"email"`
	Phone      []Phone    `json:"phone" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Address    Address    `json:"address" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Languages  []Language `gorm:"many2many:contact_languages;"`
	UserID     uint       `json:"user_id"`
}

type Language struct {
	gorm.Model
	Name string
}
