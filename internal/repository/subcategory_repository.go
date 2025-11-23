package repository

import (
	"pharmacy-team/internal/models"

	"gorm.io/gorm"
)

type SubCategoryRepository interface {
	GetCategoryByID(CategoryID int) ([]models.SubCategory, error)
	GetAllSubCategory() ([]models.SubCategory, error)
}

type gormSubCategoryRepository struct {
	db *gorm.DB
}

func NewSubCategoryRepository(db *gorm.DB) SubCategoryRepository {
	return &gormSubCategoryRepository{db: db}
}

func (r *gormSubCategoryRepository) GetAllSubCategory() ([]models.SubCategory, error) {
	var subcategory []models.SubCategory
	err := r.db.Find(&subcategory).Error
	return subcategory, err
}

func (r *gormSubCategoryRepository) GetCategoryByID(CategoryID int) ([]models.SubCategory, error) {
	var subcategories []models.SubCategory
	err := r.db.Where("category_id = ?", CategoryID).Find(&subcategories).Error
	return subcategories, err
}
