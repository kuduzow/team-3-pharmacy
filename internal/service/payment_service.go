package service

import (
	"errors"
	"pharmacy-team/internal/models"
	"pharmacy-team/internal/repository"
	"time"
)

type PaymentService interface {
	Create(orderID uint, req models.PaymentCreate) (*models.Payment, error)
	Get(id uint) (*models.Payment, error)
	Delete(id uint) error
	ListByOrder(orderID uint) ([]models.Payment, error)
}

type paymentService struct {
	payRepo   repository.PaymentRepository
	orderRepo repository.OrderRepository
}

func NewPaymentService(payRepo repository.PaymentRepository, orderRepo repository.OrderRepository) PaymentService {
	return &paymentService{payRepo: payRepo, orderRepo: orderRepo}
}

func (s *paymentService) Create(orderID uint, req models.PaymentCreate) (*models.Payment, error) {
	order, err := s.orderRepo.GetById(orderID)
	if err != nil {
		return nil, errors.New("не нашлось")
	}

	if order.Status != models.OrderStatusPendingPayment && order.Status != models.OrderStatusPaid {
		return nil, errors.New("нет такого закаса")
	}

	remaining := order.FinalPrice - order.PaidAmount
	if req.Amount <= 0 || req.Amount > remaining {
		return nil, errors.New("неправельная сумма")
	}

	if req.Status == models.PaySuccess && req.PaidAt.IsZero() {
		req.PaidAt = time.Now()
	}

	payment := &models.Payment{
		OrderID: orderID, // <- это 
		Amount: req.Amount,
		Status: req.Status,
		Method: req.Method,
		PaidAt: req.PaidAt,
	}

	payment.OrderID = orderID

	if err = s.payRepo.Create(payment); err != nil {
		return nil, err
	}

	if payment.Status == models.PaySuccess {
		order.PaidAmount += payment.Amount
		if order.PaidAmount >= order.FinalPrice {
			order.PaymentStatus = models.OrderPaymentPaid
			order.Status = models.OrderStatusPaid
		} else if order.PaidAmount > 0 {
			order.PaymentStatus = models.OrderPaymentPartiallyPaid
		}
		s.orderRepo.Update(order)
	}
	return payment, nil
}

func (s *paymentService) Get(id uint) (*models.Payment, error) {
	return s.payRepo.GetById(id)
}

func (s *paymentService) Delete(id uint) error {
	pay, err := s.payRepo.GetById(id)
	if err != nil {
		return err
	}
	if pay.Status == models.PaySuccess {
		order, _ := s.orderRepo.GetById(pay.OrderID)
		order.PaidAmount -= pay.Amount
		if order.PaidAmount <= 0 {
			order.PaidAmount = 0
			order.PaymentStatus = models.OrderPaymentNotPaid
			order.Status = models.OrderStatusPendingPayment
		} else if order.PaidAmount < order.FinalPrice {
			order.PaymentStatus = models.OrderPaymentPartiallyPaid
		}
		s.orderRepo.Update(order)
	}
	return s.payRepo.Delete(id)
}

func (s *paymentService) ListByOrder(orderID uint) ([]models.Payment, error) {
	return s.payRepo.ListByOrderID(orderID)
}
