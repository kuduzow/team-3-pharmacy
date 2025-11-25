package service

import (
	"pharmacy-team/internal/models"
	"pharmacy-team/internal/repository"
)

type CartService interface {
	CreateCart(req models.CartCreateRequest) (*models.Cart, error)
	GetCart(userID uint) (*models.Cart, error)
	AddItemToCart(userID uint, item models.CartItemRequest) (*models.Cart, error)
	RemoveItemFromCart(userID uint, medicineID uint) (*models.Cart, error)
}

type cartService struct {
	repo    repository.CartRepository
	medRepo repository.PharmacyRepository
}

// func NewCartService(
// 	repo repository.CartRepository,
// 	medRepo repository.PharmacyRepository,
// ) CartService {
// 	return &cartService{
// 		repo:    repo,
// 		medRepo: medRepo,
// 	}
// }

func (c *cartService) GetCart(userID uint) (*models.Cart, error) {	
	cart,err:=c.repo.GetCart(userID)
	if err!=nil{
		return nil,err
	}
	cart.TotalPrice = CalculateTotalPrice(cart.Items)
	return cart,nil
}

func CalculateTotalPrice(items []models.CartItem) float64 {
	var total float64
	for _, item := range items {
		total += item.Price * float64(item.Quantity)
	}

	return total
}

func (c *cartService) CreateCart(req models.CartCreateRequest) (*models.Cart, error) {
	
	cart, err := c.repo.GetOrCreateCart(req.UserID)
	 if err != nil {
		return nil, err
	}

	cart = &models.Cart{
		UserID: req.UserID,
		Items:  []models.CartItem{},
	}
	return cart, nil
}