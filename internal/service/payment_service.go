package service

import (
	"pharmacy-team/internal/repository"
)


type PaymentService struct {
	paymentRepo *repository.PaymentRepository
	// orderRepo   *repository.OrderRepository
}