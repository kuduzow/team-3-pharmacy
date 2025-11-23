package repository

import (
	"pharmacy-team/internal/models"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	Get(*[]models.Category) error
	Create(*models.Category) error
}

type gormCategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &gormCategoryRepository{db: db}
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
