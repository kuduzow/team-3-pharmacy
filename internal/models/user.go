package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	FullName       string `json:"full_name" gorm:"type:varchar(255);not null"`
	Email          string `json:"email" gorm:"type:varchar(255);uniqueIndex;not null"`
	Phone          string `json:"phone" gorm:"type:varchar(255);not null"`
	DefaultAddress string `json:"default_address" gorm:"type:varchar(255);not null"`

	//Cart    Cart    `json:"-" gorm:"foreignKey:UserID"`
	Orders   []Order   `json:"-" gorm:"foreignKey:UserID"`
	Reviews  []Review  `json:"-" gorm:"foreignKey:UserID"`
	Payments []Payment `json:"-" gorm:"foreignKey:UserID"`
}

type UserCreate struct {
	FullName       string `json:"full_name" binding:"required"`
	Email          string `json:"email" binding:"required,email"`
	Phone          string `json:"phone" binding:"required"`
	DefaultAddress string `json:"default_address" binding:"required"`
}

type UserUpdate struct {
	FullName       *string `json:"full_name"`
	Email          *string `json:"email" binding:"omitempty,email"`
	Phone          *string `json:"phone"`
	DefaultAddress *string `json:"default_address"`
}
