package repository

import (
	"pharmacy-team/internal/models"
	"time"

	"gorm.io/gorm"
)

type PromocodeRepository interface {
	Create(promocode *models.Promocode) error
	GetByID(id uint) (*models.Promocode, error)
	GetAll() ([]models.Promocode, error)
	Update(*models.Promocode) error
	Delete(id uint) error
}

type gormPromocodeRepository struct {
	db *gorm.DB // подключение к базе данных через GORM
}

func NewPromocodeRepository(db *gorm.DB) PromocodeRepository {
	return &gormPromocodeRepository{db: db}
}

func (r *gormPromocodeRepository) Create(promocode *models.Promocode) error {
	if promocode == nil {
		return nil
	}

	return r.db.Create(promocode).Error
}

func (r *gormPromocodeRepository) GetByID(id uint) (*models.Promocode, error) {
	var promocode models.Promocode

	if err := r.db.First(&promocode, id).Error; err != nil {
		return nil, err
	}

	return &promocode, nil
}

func (r *gormPromocodeRepository) GetAll() ([]models.Promocode, error) {
	var promocode []models.Promocode

	if err := r.db.Find(&promocode).Error; err != nil {
		return nil, err
	}

	return promocode, nil
}

func (r *gormPromocodeRepository) Update(promocode *models.Promocode) error {
	if promocode == nil {
		return nil
	}
	return r.db.Save(promocode).Error
}

func (r *gormPromocodeRepository) Delete(id uint) error {
	if err := r.db.Delete(&models.Promocode{}, id).Error; err != nil {
		return err
	}

	return nil
}

func (r *gormPromocodeRepository) GetByCode(code string) (*models.Promocode, error) {
	var promocode models.Promocode
	if err := r.db.Where("code = ?", code).First(&promocode).Error; err != nil {
		return nil, err
	}
	return &promocode, nil
}

func (r *gormPromocodeRepository) GetActive() ([]models.Promocode, error) {
	var promocodes []models.Promocode
	now := time.Now()
	err := r.db.Where("is_active = ? AND valid_from <= ? AND valid_to >= ?", true, now, now).Find(&promocodes).Error
	return promocodes, err
}
