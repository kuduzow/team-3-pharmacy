package service

import (
	"errors"
	"pharmacy-team/internal/models"
	"pharmacy-team/internal/repository"
)

type PromocodeService interface {
	CreatePromocode(req models.PromocodeCreate) (*models.Promocode, error)
	GetByPromocode(id uint) (*models.Promocode, error)
	UpdatePromocode(id uint, req models.PromocodeUpdate) (*models.Promocode, error)
	DeletePromocode(id uint) error
	GetAllPromocode() ([]models.Promocode, error)
}

type promocodeService struct {
	repository repository.PromocodeRepository
}

func NewPromocodeService(reposit repository.PromocodeRepository) PromocodeService {
	return &promocodeService{repository: reposit}
}

func validatePromocode(req models.PromocodeCreate) error {

	if req.Code == "" {
		return errors.New("код промокода не должен быть пустым")
	}

	if len(req.Code) < 3 {
		return errors.New("код промокода должен быть не менее 3 символов")
	}

	if req.Description == "" {
		return errors.New("описание промокода не должно быть пустым")
	}

	if req.DiscountType == "" {
		return errors.New("тип скидки не должен быть пустым")
	}

	if req.DiscountType != models.DiscountTypePercent &&
		req.DiscountType != models.DiscountTypeFixed {
		return errors.New("тип скидки должен быть 'percent' или 'fixed'")
	}

	if req.DiscountValue <= 0 {
		return errors.New("размер скидки должен быть больше 0")
	}

	if req.DiscountType == models.DiscountTypePercent && req.DiscountValue > 100 {
		return errors.New("процентная скидка не может превышать 100%")
	}

	return nil

}

func (s *promocodeService) CreatePromocode(req models.PromocodeCreate) (*models.Promocode, error) {

	if err := validatePromocode(req); err != nil {
		return nil, err
	}

	promocode := &models.Promocode{
		Code:          req.Code,
		Description:   req.Description,
		DiscountType:  req.DiscountType,
		DiscountValue: req.DiscountValue,
		ValidFrom:     req.ValidFrom,
		ValidTo:       req.ValidTo,

		IsActive:    true,
		CurrentUses: 0,
	}

	if err := s.repository.Create(promocode); err != nil {
		return nil, err
	}

	return promocode, nil
}

func (s *promocodeService) GetByPromocode(id uint) (*models.Promocode, error) {
	promo, err := s.repository.GetByID(id)
	if err != nil {
		return nil, errors.New("промокод не найден")
	}

	return promo, err
}

func (s promocodeService) UpdatePromocode(id uint, req models.PromocodeUpdate) (*models.Promocode, error) {
	promo, err := s.repository.GetByID(id)
	if err != nil {
		return nil, errors.New("промокод не найден")
	}

	if req.Code != nil {
		promo.Code = *req.Code
	}
	if req.Description != nil {
		promo.Description = *req.Description
	}
	if req.DiscountType != nil {
		promo.DiscountType = *req.DiscountType
	}
	if req.DiscountValue != nil {
		promo.DiscountValue = *req.DiscountValue
	}
	if req.MinOrderAmount != nil {
		promo.MinOrderAmount = *req.MinOrderAmount
	}
	if req.ValidFrom != nil {
		promo.ValidFrom = *req.ValidFrom
	}
	if req.ValidTo != nil {
		promo.ValidTo = *req.ValidTo
	}
	if req.IsActive != nil {
		promo.IsActive = *req.IsActive
	}

	if err = s.repository.Update(promo); err != nil {
		return nil, err
	}

	return promo, nil

}

func (s promocodeService) DeletePromocode(id uint) error {
	_, err := s.repository.GetByID(id)
	if err != nil {
		return errors.New("промокод не найден")
	}

	if result := s.repository.Delete(id); result != nil {
		return errors.New("ошибка при удалении пользователя")
	}

	return nil
}

func (s promocodeService) GetAllPromocode() ([]models.Promocode, error) {
	result, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return result, nil

}
