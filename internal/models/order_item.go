package models

import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model

	OrderID       uint   `gorm:"not null;index"`
	MedicineID    uint   `gorm:"not null"`
	MedicineName  string `gorm:"type:varchar(255);not null"`
	Quantity      int    `gorm:"not null"`
	PricePerUnit  int    `gorm:"not null"`
	LineTotal     int    `gorm:"not null"`
}