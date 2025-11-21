package repository

import (
	"pharmacy-team/internal/models"

	"gorm.io/gorm"
)

type PharmacyRepository interface {
	Create(*models.Pharmacy) error

	Update(*models.Pharmacy) error

	Delete(id uint) error

	Get(*models.Pharmacy) error

	GetByID(id uint) (*models.Pharmacy, error)

	GetAll() ([]models.Pharmacy, error)
}

type gormPharmacyRepository struct {
	db *gorm.DB
}

func NewPharmacyRepository(db *gorm.DB) PharmacyRepository {
	return &gormPharmacyRepository{db: db}
}

func (r *gormPharmacyRepository) Create(pharmacy *models.Pharmacy) error {
	if pharmacy == nil {
		return nil
	}
	return r.db.Create(pharmacy).Error
}

func (r *gormPharmacyRepository) Update(pharmacy *models.Pharmacy) error {
	if pharmacy == nil {
		return nil
	}
	return r.db.Save(pharmacy).Error
}

func (r *gormPharmacyRepository) Delete(id uint) error {
	return r.db.Delete(&models.Pharmacy{}, id).Error
}

func (r *gormPharmacyRepository) Get(pharmacy *models.Pharmacy) error {
	if err := r.db.First(pharmacy).Error; err != nil {
		return err
	}
	return nil
}
func (r *gormPharmacyRepository) GetByID(id uint) (*models.Pharmacy, error) {
	var pharmacy models.Pharmacy

	if err := r.db.First(&pharmacy, id).Error; err != nil {
		return nil, err
	}

	return &pharmacy, nil
}
func (r *gormPharmacyRepository) GetAll() ([]models.Pharmacy, error) {
	var pharmacy []models.Pharmacy
	if err := r.db.Find(&pharmacy).Error; err != nil {
		return nil, err
	}
	return pharmacy, nil
}
