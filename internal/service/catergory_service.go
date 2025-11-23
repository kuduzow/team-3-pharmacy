package service

import (
	"pharmacy-team/internal/models"
	"pharmacy-team/internal/repository"
)

type CategoryService interface {
	GetAllCategory() ([]models.Category, error)
}

type categoryService struct {
	category repository.CategoryRepository
}

func NewCategoryService(category repository.CategoryRepository) CategoryService {
	return &categoryService{category: category}
}

func (c *categoryService) GetAllCategory() ([]models.Category, error) {
	var category []models.Category

	if err := c.category.Get(&category); err != nil {
		return nil, err
	}
	return category, nil
}
