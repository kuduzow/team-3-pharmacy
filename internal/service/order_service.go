package service

import (
	"pharmacy-team/internal/models"
	"pharmacy-team/internal/repository"
	"time"
)

type OrderService struct {
	orderRepo   repository.OrderRepository
	paymentRepo repository.PaymentRepository
}

func NewOrderService(orderRepo repository.OrderRepository, paymentRepo repository.PaymentRepository) *OrderService {
	return &OrderService{
		orderRepo:   orderRepo,
		paymentRepo: paymentRepo,
	}
}

func (s *OrderService) CreateOrder(order *models.Order) (*models.Order, error) {
	var total int = 0
	for i := range order.Items {
	order.Items[i].LineTotal = order.Items[i].Quantity * order.Items[i].PricePerUnit
	total += order.Items[i].LineTotal
}

	order.TotalPrice = total
	order.FinalPrice = total - order.DiscountTotal
	order.Status = models.OrderStatusPendingPayment
	order.CreatedAt = time.Now()
	order.UpdatedAt = time.Now()

	if err := s.orderRepo.Create(order); err != nil {
		return nil, err
	}

	return order, nil
}

func (s *OrderService) GetOrderByID(id uint) (*models.Order, error) {
	order, err := s.orderRepo.GetById(id)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (s *OrderService) UpdateOrderStatus(order *models.Order, newStatus models.OrderStatus) error {
	order.Status = newStatus
	order.UpdatedAt = time.Now()
	return s.orderRepo.Update(order)
}
