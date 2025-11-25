package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	ID          int    `json:"id"`
	Name        string `json:"name"`
	SubCategory []SubCategory
}

type CategoryCreateRequest struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	SubCategory []SubCategory
}
