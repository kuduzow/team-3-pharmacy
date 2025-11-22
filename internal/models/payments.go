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
	OrderID uint   `json:"order_id" gorm:"not null;index"` 
	Amount  int              `json:"amount" gorm:"column:amount;not null"`
	Status  PayStatus        `json:"status" gorm:"column:status;type:varchar(20);not null"`
	Method  PayMethod        `json:"method" gorm:"column:method;type:varchar(20);not null"`
	PaidAt  time.Time        `json:"paid_at" gorm:"column:paid_at"`
}

type PaymentCreate struct {
  	OrderID uint   `json:"order_id" gorm:"not null;index"` 
	Amount int       `json:"amount"`
	Status PayStatus `json:"status"`
	Method PayMethod `json:"method"`
	PaidAt time.Time `json:"paid_at"`
}

type PaymentUpdate struct {
    	OrderID uint   `json:"order_id" gorm:"not null;index"` 
	Amount *int       `json:"amount"`
	Status *PayStatus `json:"status"`
	Method *PayMethod `json:"method"`
	PaidAt *time.Time `json:"paid_at"`
}
