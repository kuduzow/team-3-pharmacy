package repository

import (
	"pharmacy-team/internal/models"

	"gorm.io/gorm"
)

type PharmacyRepository interface {
	Create(*models.Pharmacy) error

	Update(*models.Pharmacy) error

	Delete(id uint) error
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
