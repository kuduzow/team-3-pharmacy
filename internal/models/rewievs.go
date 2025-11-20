package models

import "gorm.io/gorm"

type Review struct {
	*gorm.Model
	UserID     uint   `gorm:"not null" json:"user_id"`
	MedicineID uint   `gorm:"not null" json:"medicine_id"`
	Rating     int    `gorm:"not null;check:rating>=0 AND rating<=5" json:"rating"`
	Text       string `gorm:"type:text" json:"text"`
}

type CreateReviewRequest struct {
    UserID     uint   `json:"user_id" binding:"required"`
    MedicineID uint   `json:"medicine_id" binding:"required"`
    Rating     int    `json:"rating" binding:"required,min=0,max=5"`
    Text       string `json:"text" binding:"required"`
}

type UpdateReviewRequest struct {
    Rating *int   `json:"rating" binding:"omitempty,min=0,max=5"`
    Text   *string `json:"text" binding:"omitempty"`
}
