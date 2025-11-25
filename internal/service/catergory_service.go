package service

import (
	"errors"
	"pharmacy-team/internal/models"
	"pharmacy-team/internal/repository"
	"strings"
)

type CategoryService interface {
	GetAllCategory() ([]models.Category, error)
	CreateCategory(req models.CategoryCreateRequest) (*models.Category, error)
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

func (c *categoryService) CreateCategory(req models.CategoryCreateRequest) (*models.Category, error) {
	if err := c.validateCategoryCreate(req); err != nil {
		return nil, err
	}
	category := &models.Category{
		ID:   req.ID,
		Name: req.Name,
	}

	if err := c.category.Create(category); err != nil {
		return nil, err
	}
	return category, nil
}

func (c *categoryService) validateCategoryCreate(req models.CategoryCreateRequest) error {
	if req.ID < 0 {
		return errors.New("айди не может быть отрицательным")
	}

	if strings.TrimSpace(req.Name) == "" {
		return errors.New("имя не может быть пустым")
	}

	return nil
}
