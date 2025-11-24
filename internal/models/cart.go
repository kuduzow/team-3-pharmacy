package models

import "gorm.io/gorm"

type Cart struct {
	*gorm.Model
	MedicineID uint
	Quantity   int
	PricePerUnit float64
	TotalPrice   float64
}