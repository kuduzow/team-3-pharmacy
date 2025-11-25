package service

import (
	"errors"
	"pharmacy-team/internal/models"
	"pharmacy-team/internal/repository"
)

type OrderService interface {
	CreateOrder(userID uint, req models.OrderCreate) (*models.Order, error)
	GetOrder(id uint) (*models.Order, error)
	GetUserOrders(userID uint) ([]models.Order, error)
	UpdateStatus(id uint, status models.OrderStatus) (*models.Order, error)
}

type orderService struct {
	repo       repository.OrderRepository
	// cart       CartService
	// promocodes PromocodeService
}

func NewOrderService(repo repository.OrderRepository, /*cart CartService, promo PromocodeService*/) OrderService {
	return &orderService{repo: repo, /*cart: cart, promocodes: promo*/}
}

func (s *orderService) CreateOrder(userID uint, req models.OrderCreate) (*models.Order, error) {
	order := &models.Order{
		UserID:          userID,
		Status:          models.OrderStatusPendingPayment,
		PaymentStatus:   models.OrderPaymentNotPaid,
		DeliveryAddress: req.DeliveryAddress,
		Comment:         req.Comment,
		TotalPrice:      0,  
		DiscountTotal:   0,
		FinalPrice:      0,
		Items:           []models.OrderItem{},
	}

	if err := s.repo.Create(order); err != nil {
		return nil, err
	}

	return order, nil
}


func (s *orderService) GetOrder(id uint) (*models.Order, error) {
	return s.repo.GetByID(id)
}

func (s *orderService) GetUserOrders(userID uint) ([]models.Order, error) {
	return s.repo.GetByUserID(userID)
}

func (s *orderService) UpdateStatus(id uint, status models.OrderStatus) (*models.Order, error) {
	order, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if order.Status == models.OrderStatusPaid && status == models.OrderStatusPendingPayment {
    return nil, errors.New("так нельзя")
}

	order.Status = status

	if err := s.repo.Update(order); err != nil {
		return nil, err
	}

	return order, nil
}
