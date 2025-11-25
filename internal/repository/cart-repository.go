package repository

import (
	"pharmacy-team/internal/models"

	"gorm.io/gorm"
)

type CartRepository interface {
	GetCart(userID uint) (*models.Cart, error)
	GetOrCreateCart(userID uint) (*models.Cart, error)
	GetCartItem(cartID uint, medicineID uint) (*models.CartItem, error)
	AddItem(item *models.CartItem) error
	UpdateItem(item *models.CartItem) error
	DeleteItem(itemID uint) error
}
type gormCartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) CartRepository {
	return &gormCartRepository{db: db}
}

func (r *gormCartRepository) GetCart(userID uint) (*models.Cart, error) {
	var cart models.Cart

	err := r.db.Preload("Items").Where("user_id = ?", userID).First(&cart).Error
	if err != nil {
		return nil, err
	}

	return &cart, nil
}

func (r *gormCartRepository) GetOrCreateCart(userID uint) (*models.Cart, error) {
	var cart models.Cart

	err := r.db.Preload("Items").Where("user_id = ?", userID).First(&cart).Error

	if err == nil {
		return &cart, nil
	}

	cart = models.Cart{UserID: userID}

	if err := r.db.Create(&cart).Error; err != nil {
		return nil, err
	}

	return &cart, nil
}

func (r *gormCartRepository) GetCartItem(cartID uint, medicineID uint) (*models.CartItem, error) {
    var item models.CartItem
    err := r.db.Where("cart_id = ? AND medicine_id = ?", cartID, medicineID).
        First(&item).Error

    if err != nil {
        return nil, err
    }

    return &item, nil
}

func (r *gormCartRepository) AddItem(item *models.CartItem) error {
    return r.db.Create(item).Error
}

func (r *gormCartRepository) UpdateItem(item *models.CartItem) error {
    return r.db.Save(item).Error
}

func (r *gormCartRepository) DeleteItem(itemID uint) error {
    return r.db.Delete(&models.CartItem{}, itemID).Error
}

