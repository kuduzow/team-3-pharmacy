package models

import "gorm.io/gorm"

type SubCategory struct {
	gorm.Model
	ID         int    `json:"id"`
	CategoryID int    `json:"category_id"`
	Name       string `json:"name"`
}

type SubCategoryCreateRequest struct {
	ID         int    `json:"id"`
	CategoryID int    `json:"category_id"`
	Name       string `json:"name"`
}
