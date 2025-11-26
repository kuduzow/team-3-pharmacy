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
	orderRepo repository.OrderRepository
	cartRepo  repository.CartRepository
}

func NewOrderService(orderRepo repository.OrderRepository, cartRepo repository.CartRepository) OrderService {
	return &orderService{
		orderRepo: orderRepo,
		cartRepo:  cartRepo,
	}
}

func (s *orderService) CreateOrder(userID uint, req models.OrderCreate) (*models.Order, error) {
	cart, err := s.cartRepo.GetCart(userID)
	if err != nil {
		return nil, errors.New("корзина не найдена")
	}

	if len(cart.Items) == 0 {
		return nil, errors.New("корзина пуста")
	}

	orderItems := make([]models.OrderItem, len(cart.Items))
	total := 0
	for i, item := range cart.Items {
		lineTotal := int(item.PricePerUnit * float64(item.Quantity))
		orderItems[i] = models.OrderItem{
			PharmacyID:   0, 
			PharmacyName: item.Name,
			Quantity:     item.Quantity,
			PricePerUnit: int(item.PricePerUnit),
			LineTotal:    lineTotal,
		}
		total += lineTotal
	}

	order := &models.Order{
		UserID:          userID,
		CartID:          &cart.ID,
		Status:          models.OrderStatusPendingPayment,
		PaymentStatus:   models.OrderPaymentNotPaid,
		TotalPrice:      total,
		FinalPrice:      total,
		DiscountTotal:   0,
		DeliveryAddress: req.DeliveryAddress,
		Comment:         req.Comment,
		Items:           orderItems,
	}

	if err := s.orderRepo.Create(order); err != nil {
		return nil, err
	}

	_ = s.cartRepo.DeleteCart(userID)

	return order, nil
}

func (s *orderService) GetOrder(id uint) (*models.Order, error) {
	return s.orderRepo.GetByID(id)
}

func (s *orderService) GetUserOrders(userID uint) ([]models.Order, error) {
	return s.orderRepo.GetByUserID(userID)
}

func (s *orderService) UpdateStatus(id uint, status models.OrderStatus) (*models.Order, error) {
	order, err := s.orderRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if order.Status == models.OrderStatusPaid && status == models.OrderStatusPendingPayment {
		return nil, errors.New("невозможно изменить")
	}

	order.Status = status
	if err := s.orderRepo.Update(order); err != nil {
		return nil, err
	}

	return order, nil
}
