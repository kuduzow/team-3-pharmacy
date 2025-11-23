package service

import (
	"pharmacy-team/internal/models"
	"pharmacy-team/internal/repository"
)

type SubCategoryService interface {
	GetSubCategory(CategoryID uint) ([]models.SubCategory, error)
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
