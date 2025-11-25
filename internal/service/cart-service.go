package service

import (
	"errors"
	"pharmacy-team/internal/models"
	"pharmacy-team/internal/repository"
)

type CartService interface {
	CreateCart(req models.CartCreateRequest) (*models.Cart, error)
	GetCart(userID uint) (*models.Cart, error)
	AddItemToCart(userID uint, item models.CartItemRequest) (*models.Cart, error)
	RemoveItemFromCart(userID uint, medicineID uint) (*models.Cart, error)
	UpdateItemInCart(userID uint, itemID uint, quantity int) (*models.Cart, error)
	DeleteCart(userID uint) error
}

type cartService struct {
	repo    repository.CartRepository
	medRepo repository.PharmacyRepository
}

func NewCartService(
	repo repository.CartRepository,
	medRepo repository.PharmacyRepository,
) CartService {
	return &cartService{
		repo:    repo,
		medRepo: medRepo,
	}
}

func (c *cartService) GetCart(userID uint) (*models.Cart, error) {
	cart, err := c.repo.GetCart(userID)
	if err != nil {
		return nil, err
	}
	cart.TotalPrice = CalculateTotalPrice(cart.Items)
	return cart, nil
}

func CalculateTotalPrice(items []models.CartItem) float64 {
	var total float64
	for _, item := range items {
		total += item.PricePerUnit * float64(item.Quantity)
	}

	return total
}

func (s *cartService) CreateCart(req models.CartCreateRequest) (*models.Cart, error) {
	return s.repo.GetOrCreateCart(req.UserID)
}

func (c *cartService) UpdateItemInCart(userID uint, itemID uint, quantity int) (*models.Cart, error) {
	cart, err := c.repo.GetOrCreateCart(userID)
	if err != nil {
		return nil, err
	}

	item, err := c.repo.GetItemByID(itemID)
	if err != nil {
		return nil, err
	}

	if item.CartID != cart.ID {
		return nil, errors.New("item does not belong to user’s cart")
	}

	item.Quantity = quantity

	if err := c.repo.UpdateItem(item); err != nil {
		return nil, err
	}

	return c.repo.GetCart(userID)
}

func (c *cartService) DeleteCart(userID uint) error {
	return c.repo.DeleteCart(userID)
}
func (c *cartService) RemoveItemFromCart(userID uint, medicineID uint) (*models.Cart, error) {
	cart, err := c.repo.GetOrCreateCart(userID)
	if err != nil {
		return nil, err
	}
	item, err := c.repo.GetCartItem(cart.ID, medicineID)
	if err != nil {
		return nil, err
	}
	if err := c.repo.DeleteItem(item.ID); err != nil {
		return nil, err
	}	
	return c.repo.GetCart(userID)
}
func (c *cartService) AddItemToCart(userID uint, itemReq models.CartItemRequest) (*models.Cart, error) {
	cart, err := c.repo.GetOrCreateCart(userID)
	if err != nil {
		return nil, err
	}
	medicine, err := c.medRepo.GetByID(itemReq.MedicineID)
	if err != nil {
		return nil, err
	}
	cartItem := &models.CartItem{
		CartID:       cart.ID,
		MedicineID:   medicine.ID,
		Name:         medicine.Name,
		Quantity:     itemReq.Quantity,
		PricePerUnit: medicine.Price,
	}
	if err := c.repo.AddItem(cartItem); err != nil {
		return nil, err
	}
	return c.repo.GetCart(userID)
}
