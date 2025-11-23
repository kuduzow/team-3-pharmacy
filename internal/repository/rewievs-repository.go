package repository

import (
	"errors"
	"pharmacy-team/internal/models"

	"gorm.io/gorm"
)

type ReviewRepository interface {
	Create( *models.Review) error
	Update(review *models.Review) error
	Delete(id uint) error
	GetReviewsByPharmacyID(pharmacyID uint) ([]models.Review, error)
	GetByID(id uint)(*models.Review,error)
	CanUserReviewMedicine(userID, medicineID uint) error
}

type gormReviewRepository struct {
	db *gorm.DB
}

func NewReviewRepository(db *gorm.DB) ReviewRepository {
	return &gormReviewRepository{db: db}
}

func (r *gormReviewRepository) Create(review *models.Review) error {

	if review == nil {
		return errors.New("review is nil")
	}
	
	return r.db.Create(review).Error
}

func (r *gormReviewRepository) Update(review *models.Review) error {
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

func (r *gormReviewRepository) GetByID(id uint)(*models.Review, error){
	var review models.Review
	if err := r.db.First(&review,id).Error;err != nil{
		return nil, err 
	}
	return &review, nil
}

func (r *gormReviewRepository) CanUserReviewMedicine(userID, medicineID uint) error {
	var dummy struct{}

	err := r.db.Table("order_items").
		Select("1").
		Joins("JOIN orders o ON o.id = order_items.order_id").
		Where("o.user_id = ? AND order_items.medicine_id = ?", userID, medicineID).
		First(&dummy).Error

	if err != nil {
		return errors.New("пользователь не может оставить отзыв на это лекарство")
	}

	return nil
}