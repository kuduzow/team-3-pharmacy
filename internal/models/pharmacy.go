package models

import "gorm.io/gorm"

type Pharmacy struct {
	gorm.Model
	Name                 string  `json:"name"`
	Description          string  `json:"description"`
	Price                float64 `json:"price"`
	InStock              bool    `json:"in_stock"`
	StockQuantity        int     `json:"stock_quantity"`
	CategoryID           uint    `json:"category_id"`
	SubcategoryID        uint    `json:"Subcategory_id"`
	Manufacturer         string  `json:"manufacturer"`
	PrescriptionRequired bool    `json:"prescription_required"`
	AvgRating            float64 `json:"avg_rating"`
	Category             *Category
}

type PharmacyCreateRequest struct {
	Name                 string  `json:"name"`
	Description          string  `json:"description"`
	Price                float64 `json:"price"`
	InStock              bool    `json:"in_stock"`
	StockQuantity        int     `json:"stock_quantity"`
	CategoryID           uint    `json:"category_id"`
	SubcategoryID        uint    `json:"Subcategory_id"`
	Manufacturer         string  `json:"manufacturer"`
	PrescriptionRequired bool    `json:"prescription_required"`
	AvgRating            float64 `json:"avg_rating"`
	Category             *Category
}

type PharmacyUpdateRequest struct {
	Name                 *string  `json:"name"`
	Description          *string  `json:"description"`
	Price                *float64 `json:"price"`
	InStock              *bool    `json:"in_stock"`
	StockQuantity        *int     `json:"stock_quantity"`
	CategoryID           *uint    `json:"category_id"`
	SubcategoryID        *uint    `json:"Subcategory_id"`
	Manufacturer         *string  `json:"manufacturer"`
	PrescriptionRequired *bool    `json:"prescription_required"`
	AvgRating            *float64 `json:"avg_rating"`
	Category             *Category
}
