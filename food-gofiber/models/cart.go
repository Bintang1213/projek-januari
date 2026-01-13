package models

type Cart struct {
	ID       uint `gorm:"primaryKey" json:"id"`
	UserID   uint `gorm:"not null" json:"user_id"`
	MenuID   uint `gorm:"not null" json:"menu_id"`
	Quantity int  `gorm:"default:1" json:"quantity"`
	Menu     Menu `gorm:"foreignKey:MenuID" json:"menu"` 
}