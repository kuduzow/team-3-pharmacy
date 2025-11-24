package models

import (
	"time"

	"gorm.io/gorm"
)

type DiscountType string

const (
	DiscountTypePercent DiscountType = "percent"
	DiscountTypeFixed   DiscountType = "fixed"
)

type Promocode struct {
	gorm.Model
	Code           string       `json:"code" gorm:"type:varchar(50);uniqueIndex;not null"`
	Description    string       `json:"description" gorm:"type:text"`
	DiscountType   DiscountType `json:"discount_type" gorm:"type:varchar(20);not null"`
	DiscountValue  float64      `json:"discount_value" gorm:"type:decimal(10,2);not null"`
	MinOrderAmount float64      `json:"min_order_amount" gorm:"type:decimal(10,2);default:0"`
	ValidFrom      time.Time    `json:"valid_from"`
	ValidTo        time.Time    `json:"valid_to"`
	MaxUses        *int         `json:"max_uses"`
	CurrentUses    int          `json:"current_uses" gorm:"default:0"`
	MaxUsesPerUser *int         `json:"max_uses_per_user"`
	IsActive       bool         `json:"is_active" gorm:"default:true"`
}

type PromocodeCreate struct {
	Code           string       `json:"code" `
	Description    string       `json:"description" `
	DiscountType   DiscountType `json:"discount_type"`
	DiscountValue  float64      `json:"discount_value" `
	MinOrderAmount float64      `json:"min_order_amount" `
	ValidFrom      time.Time    `json:"valid_from" `
	ValidTo        time.Time    `json:"valid_to" `
	MaxUses        *int         `json:"max_uses" `
	MaxUsesPerUser *int         `json:"max_uses_per_user" `
}

type PromocodeUpdateRequest struct {
	Code           *string       `json:"code"`
	Description    *string       `json:"description"`
	DiscountType   *DiscountType `json:"discount_type"`
	DiscountValue  *float64      `json:"discount_value"`
	ValidFrom      *time.Time    `json:"valid_from"`
	ValidTo        *time.Time    `json:"valid_to"`
	MinOrderAmount *float64      `json:"min_order_amount"`
	MaxUses        *int          `json:"max_uses"`
	MaxUsesPerUser *int          `json:"max_uses_per_user"`
	IsActive       *bool         `json:"is_active"`
}
