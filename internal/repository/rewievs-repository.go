package repository

import (
	"errors"
	"pharmacy-team/internal/models"

	"gorm.io/gorm"
)

type ReviewRepository interface {
	Create(review *models.CreateReviewRequest) error
	Update(review *models.UpdateReviewRequest) error
	Delete(id uint) error
	GetReviewsByPharmacyID(pharmacyID uint) ([]models.Review, error)
}

type gormReviewRepository struct {
	db *gorm.DB
}

func NewReviewRepository(db *gorm.DB) ReviewRepository {
	return &gormReviewRepository{db: db}
}

func (r *gormReviewRepository) Create(review *models.CreateReviewRequest) error {
	if review == nil {
		return errors.New("review is nil")
	}
	
	return r.db.Create(review).Error
}

func (r *gormReviewRepository) Update(review *models.UpdateReviewRequest) error {
	if review == nil {
		return errors.New("review is nil")
	}
	
	return r.db.Save(review).Error
}

func (r *gormReviewRepository) Delete(id uint) error {
	return r.db.Delete(&models.Review{}, id).Error
}

func (r *gormReviewRepository) GetReviewsByPharmacyID(pharmacyID uint) ([]models.Review, error) {
	var reviews []models.Review
	err := r.db.Where("pharmacy_id = ?", pharmacyID).Find(&reviews).Error
	return reviews, err		
}