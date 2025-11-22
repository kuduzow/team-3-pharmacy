package service

import (
	"errors"
	"pharmacy-team/internal/models"
	"pharmacy-team/internal/repository"
	"strings"
)

type PharmacyService interface {
	CreatePharmacy(req models.PharmacyCreateRequest) (*models.Pharmacy, error)

	UpdatePharmacy(id uint, req models.PharmacyUpdateRequest) (*models.Pharmacy, error)

	DeletePharmacy(id uint) error

	GetPharmacy() ([]models.Pharmacy, error)

	GetPharmacyByID(id uint) (*models.Pharmacy, error)
}

type pharmacyService struct {
	pharmacy repository.PharmacyRepository
}

func NewPharmacyService(pharmacy repository.PharmacyRepository) PharmacyService {
	return &pharmacyService{pharmacy: pharmacy}
}

func (s *pharmacyService) GetPharmacyByID(id uint) (*models.Pharmacy, error) {
	pharmacy, err := s.pharmacy.GetByID(id)
	if err != nil {
		return nil, errors.New("лекарство не найдено")
	}

	return pharmacy, nil
}

func (s *pharmacyService) CreatePharmacy(req models.PharmacyCreateRequest) (*models.Pharmacy, error) {
	if err := s.validatePharmacyCreate(req); err != nil {
		return nil, err
	}

	pharmacy := &models.Pharmacy{
		Name:                 req.Name,
		Description:          req.Description,
		Price:                req.Price,
		InStock:              req.InStock,
		StockQuantity:        req.StockQuantity,
		CategoryID:           req.CategoryID,
		SubcategoryID:        req.SubcategoryID,
		Manufacturer:         req.Manufacturer,
		PrescriptionRequired: req.PrescriptionRequired,
		AvgRating:            req.AvgRating,
	}

	if err := s.pharmacy.Create(pharmacy); err != nil {
		return nil, err
	}

	return pharmacy, nil
}

func (s *pharmacyService) UpdatePharmacy(id uint, req models.PharmacyUpdateRequest) (*models.Pharmacy, error) {
	pharmacy, err := s.pharmacy.GetByID(id)
	if err != nil {
		return nil, errors.New("не должно быть пустым")
	}
	s.applyPharmacyUpdate(pharmacy, req)
	if err := s.pharmacy.Update(pharmacy); err != nil {
		return nil, err
	}

	return pharmacy, nil

}

func (s *pharmacyService) DeletePharmacy(id uint) error {
	if _, err := s.pharmacy.GetByID(id); err != nil {
		return errors.New("error")
	}
	return s.pharmacy.Delete(id)
}

func (s *pharmacyService) GetPharmacy() ([]models.Pharmacy, error) {
	pharmacy, err := s.pharmacy.GetAll()
	if err != nil {
		return nil, err
	}
	return pharmacy, nil
}

func (s *pharmacyService) validatePharmacyCreate(req models.PharmacyCreateRequest) error {
	if strings.TrimSpace(req.Name) == "" {
		return errors.New("название не должно быть пустым")
	}

	if strings.TrimSpace(req.Description) == "" {
		return errors.New("описание не должно быть пустым")
	}

	if req.Price < 0 {
		return errors.New("цена не может быть отрицательной")
	}

	if req.StockQuantity < 0 {
		return errors.New("на складе не может быть отрицательное количество товаров")
	}

	if !req.InStock && req.StockQuantity > 0 {
		return errors.New("если товара нет на складе его не может быть в наличии")
	}
	return nil
}
func (s *pharmacyService) applyPharmacyUpdate(pharmacy *models.Pharmacy, req models.PharmacyUpdateRequest) error {

	if req.Name != nil {
		trimmed := strings.TrimSpace(*req.Name)
		if trimmed == "" {
			return errors.New("поле name не должно быть пустым")
		}
		pharmacy.Name = trimmed
	}

	if req.Description != nil {
		pharmacy.Description = strings.TrimSpace(*req.Description)
	}

	if req.Price != nil {
		pharmacy.Price = *req.Price
	}

	if req.InStock != nil {
		pharmacy.InStock = *req.InStock
	}

	if !pharmacy.InStock && pharmacy.StockQuantity > 0 {
		return errors.New("если товара нет на складе его не может быть в наличии")
	}

	if req.PrescriptionRequired != nil {
		pharmacy.PrescriptionRequired = *req.PrescriptionRequired
	}
	return nil

}
