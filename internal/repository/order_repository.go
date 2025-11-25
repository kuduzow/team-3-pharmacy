package repository

import (
	"fmt"
	"pharmacy-team/internal/models"

	"gorm.io/gorm"
)

type OrderRepository interface {
	Create(order *models.Order) error
	GetByID(id uint) (*models.Order, error)
	GetByUserID(userID uint) ([]models.Order, error)
	Update(order *models.Order) error
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db: db}
}

func (r *orderRepository) Create(order *models.Order) error {
	return r.db.Create(order).Error
}

func (r *orderRepository) GetByID(id uint) (*models.Order, error) {
	var order models.Order
	err := r.db.Preload("Товар").Preload("Оплата").First(&order, id).Error
	return &order, err
}

func (r *orderRepository) GetByUserID(userID uint) ([]models.Order, error) {
	var orders []models.Order
	err := r.db.Where("user_id = ?", userID).Find(&orders).Error
	if err != nil {
		fmt.Errorf("error",err)
	}
	return orders, err
}

func (r *orderRepository) Update(order *models.Order) error {
	return r.db.Save(order).Error
}
