package models

import "gorm.io/gorm"

type Cart struct {
	*gorm.Model
	UserID     uint       `json:"user_id"`
	Items      []CartItem `json:"items" gorm:"foreignKey:CartID"`
	TotalPrice float64    `json:"total_price"`
}

type CartItem struct {
	*gorm.Model
	CartID       uint    `json:"cart_id" gorm:"index"`
	MedicineID   uint    `json:"medicine_id"`
	Name         string  `json:"name"`
	Quantity     int     `json:"quantity"`
	PricePerUnit float64 `json:"price_per_unit"`
}

type CartCreateRequest struct {
	UserID uint              `json:"user_id"`
	Items  []CartItemRequest `json:"items"`
}

type CartItemRequest struct {
	MedicineID uint `json:"medicine_id"`
	Quantity   int  `json:"quantity"`
}

type CartUpdateRequest struct {
	Quantity int `json:"quantity"`
}
