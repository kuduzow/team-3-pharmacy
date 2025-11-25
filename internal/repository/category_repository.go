package repository

import (
	"pharmacy-team/internal/models"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	Get(*[]models.Category) error

	Create(*models.Category) error

	GetSubcategories(categoryID uint) ([]models.SubCategory, error)
}

type gormCategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &gormCategoryRepository{db: db}
}

func (r *gormCategoryRepository) GetSubcategories(categoryID uint) ([]models.SubCategory, error) {
	var req []models.SubCategory
	if err := r.db.Where("category_id = ?", categoryID).Find(&req).Error; err != nil {
		return nil, err
	}
	return req, nil
}

func (r *gormCategoryRepository) Get(category *[]models.Category) error {
	if err := r.db.First(category).Error; err != nil {
		return err
	}
	return nil
}

func (r *gormCategoryRepository) Create(category *models.Category) error {
	if category == nil {
		return nil
	}
	return r.db.Save(category).Error
}
