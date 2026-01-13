package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID        uint        `json:"user_id"`
	User          User        `json:"user" gorm:"foreignKey:UserID"` 
	Total         float64     `json:"total"`
	Status        string      `json:"status"`
	PaymentMethod string      `json:"payment_method"`
	OrderMethod   string      `json:"order_method"`
	Address       string      `json:"address"`
	Note          string      `json:"note"`
	Items         []OrderItem `json:"items" gorm:"foreignKey:OrderID"`
}

type OrderItem struct {
	gorm.Model
	OrderID uint    `json:"order_id"`
	MenuID  uint    `json:"menu_id"`
	Name    string  `json:"name"`
	Price   float64 `json:"price"`
	Qty     int     `json:"qty"`
}