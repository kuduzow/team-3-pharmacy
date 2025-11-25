package models

import "gorm.io/gorm"

type Cart struct {
	*gorm.Model
	UserID     uint `json:"user_id"`
	Items   []CartItem `json:"items" gorm:"foreignKey:CartID"`
	TotalPrice float64 `json:"total_price"`
}

type CartItem struct {
	*gorm.Model
	MedicineID uint    `json:"medicine_id"`
	Quantity   int    `json:"quantity"`
	Price      float64 `json:"price"`
}

type CartCreateRequest struct {
	UserID     uint          `json:"user_id"`
	Items   []CartItemRequest `json:"items"`
}


type CartItemRequest struct {
	MedicineID uint    `json:"medicine_id"`
	Quantity   int    `json:"quantity"`
}	