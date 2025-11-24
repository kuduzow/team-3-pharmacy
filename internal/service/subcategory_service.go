package service

import (
	"errors"
	"pharmacy-team/internal/models"
	"pharmacy-team/internal/repository"
	"strings"
)

type SubCategoryService interface {
	GetSubCategory(CategoryID uint) ([]models.SubCategory, error)
	CreateSubCategory(categoryID uint, req models.SubCategoryCreateRequest) (*models.SubCategory, error)
}

type subCategoryService struct {
	subcategoryRepo repository.SubCategoryRepository
}

func NewSubCategoryService(subcategoryRepo repository.SubCategoryRepository) SubCategoryService {
	return &subCategoryService{subcategoryRepo: subcategoryRepo}
}

func (s *subCategoryService) GetSubCategory(categoryID uint) ([]models.SubCategory, error) {
	return s.subcategoryRepo.GetAllSubCategory()
}

func (s *subCategoryService) CreateSubCategory(categoryID uint, req models.SubCategoryCreateRequest) (*models.SubCategory, error) {
	if err := s.validateSubCategoryCreate(categoryID, req); err != nil {
		return nil, err
	}
	subcategory := &models.SubCategory{
		ID:         req.ID,
		CategoryID: categoryID,
		Name:       req.Name,
	}

	if err := s.subcategoryRepo.Create(subcategory); err != nil {
		return nil, err
	}
	return subcategory, nil
}
func (c *subCategoryService) validateSubCategoryCreate(categoryID uint, req models.SubCategoryCreateRequest) error {
	if req.ID < 0 {
		return errors.New("айди не может быть отрицательным")
	}

	if categoryID > 0 {
		return errors.New("айди категории не может быть отрицательны")
	}

	if strings.TrimSpace(req.Name) == "" {
		return errors.New("имя не может быть пустым")
	}

	return nil
}
