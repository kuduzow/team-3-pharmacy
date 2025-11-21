package models

import "gorm.io/gorm"

type OrderStatus string

const (
	OrderStatusDraft          OrderStatus = "draft"
	OrderStatusPendingPayment OrderStatus = "pending_payment"
	OrderStatusPaid           OrderStatus = "paid"
	OrderStatusCanceled       OrderStatus = "canceled"
	OrderStatusShipped        OrderStatus = "shipped"
	OrderStatusCompleted      OrderStatus = "completed"
)

type OrderPaymentStatus string

const (
	OrderPaymentNotPaid       OrderPaymentStatus = "not_paid"
	OrderPaymentPartiallyPaid OrderPaymentStatus = "partially_paid"
	OrderPaymentPaid          OrderPaymentStatus = "paid"
)

type Order struct {
	gorm.Model

	UserID          uint               `gorm:"not null;index"`
	Status          OrderStatus        `gorm:"type:varchar(30);default:'pending_payment'"`
	PaymentStatus   OrderPaymentStatus `gorm:"type:varchar(30);default:'not_paid'"`
	TotalPrice      int                `gorm:"not null"`
	DiscountTotal   int                `gorm:"default:0"`
	FinalPrice      int                `gorm:"not null"`
	PaidAmount      int                `gorm:"default:0"`
	DeliveryAddress string             `gorm:"type:text;not null"`
	Comment         string             `gorm:"type:text"`
	PromoCode       *string            `gorm:"type:varchar(50)"`

	Items    []OrderItem `gorm:"constraint:OnDelete:CASCADE"`
	Payments []Payment   `gorm:"constraint:OnDelete:SET NULL"`
}

type OrderCreateRequest struct {
	DeliveryAddress string  `json:"delivery_address" binding:"required"`
	Comment         string  `json:"comment,omitempty"`
	PromoCode       *string `json:"promocode,omitempty"`
}

type OrderUpdateRequest struct {
	Status          *OrderStatus        `json:"status,omitempty"`
	PaymentStatus   *OrderPaymentStatus `json:"payment_status,omitempty"`
	DeliveryAddress *string             `json:"delivery_address,omitempty"`
	Comment         *string             `json:"comment,omitempty"`
}