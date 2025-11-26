package models

import (
	"time"

	"gorm.io/gorm"
)

type PayMethod string

const (
	PayCard             PayMethod = "card"
	PaymentCash         PayMethod = "cash"
	PaymentBankTransfer PayMethod = "bank_transfer"
)

type PayStatus string

const (
	PayPending PayStatus = "pending"
	PaySuccess PayStatus = "success"
	PayFailed  PayStatus = "failed"
)

type Payment struct {
    gorm.Model
    OrderID uint      `gorm:"not null;index"`
    Amount  int       `gorm:"not null"`
    Status  PayStatus `gorm:"type:varchar(20);not null"`
    Method  PayMethod `gorm:"type:varchar(20);not null"`
    PaidAt  time.Time
}

type PaymentCreate struct {
	OrderID uint      `json:"order_id"`
	Amount  int       `json:"amount"`
	Status  PayStatus `json:"status"`
	Method  PayMethod `json:"method"`
	PaidAt  time.Time `json:"paid_at"`
}

type PaymentUpdate struct {
	OrderID uint       `json:"order_id"`
	Amount  *int       `json:"amount"`
	Status  *PayStatus `json:"status"`
	Method  *PayMethod `json:"method"`
	PaidAt  *time.Time `json:"paid_at"`
}
