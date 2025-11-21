package repository

import (
	"pharmacy-team/internal/models"

	"gorm.io/gorm"
)

type OrderRepository interface {
	Create(order *models.Order) error
	GetByLookId(id uint) (*models.Order, error)
	GetById(id uint) (*models.Order, error)
	GetUserOrders(userID uint) ([]models.Order, error)
	Update(order *models.Order) error
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db: db}
}

func (r orderRepository) Create(order *models.Order) error {
	if order != nil {
		return nil
	}
	return r.db.Create(order).Error
}

func (r orderRepository) GetByLookId(id uint) (*models.Order, error) {
	var order models.Order

	if err := r.db.First(&order, id).Error; err != nil {
		return nil, err
	}

	return &order, nil
}

func (r *orderRepository) GetById(id uint) (*models.Order, error) {
	var order models.Order
	if err := r.db.First(&order, id).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *orderRepository) GetUserOrders(userID uint) ([]models.Order, error) {
	var orders []models.Order
	if err := r.db.Where("user_id = ?", userID).Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *orderRepository) Update(order *models.Order) error {
	return r.db.Save(order).Error
}