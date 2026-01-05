package models

import "gorm.io/gorm"

type Menu struct {
	gorm.Model
	Name        string  `json:"name" gorm:"not null"`
	Description string  `json:"description"`
	Price       float64 `json:"price" gorm:"not null"`
	Category    string  `json:"category"`
	Image       string  `json:"image"` 
}