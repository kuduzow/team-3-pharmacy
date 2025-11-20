package repository

import (
	"pharmacy-team/internal/models"

	"gorm.io/gorm"
)

type PaymentRepository interface {
	Create(payment *models.Payment) error
	GetById(id uint) (*models.Payment, error)
	Delete(id uint) error
	//ListByUserID(userID uint) ([]models.Payment, error)
}

type gormPaymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &gormPaymentRepository{db: db}
}

func (r *gormPaymentRepository) Create(payment *models.Payment) error {
	if payment == nil {
		return nil
	}
	return r.db.Create(payment).Error
}

func (r *gormPaymentRepository) GetById(id uint) (*models.Payment, error) {
	var Payment models.Payment
	if err := r.db.First(&Payment, id).Error; err != nil  {
		return nil,err
	}
	return &Payment,nil
}

func (r *gormPaymentRepository) Delete(id uint) error {
		return r.db.Delete(&models.Payment{}, id).Error
}
