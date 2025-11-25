package transport

import (
	"pharmacy-team/internal/models"
	"pharmacy-team/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	service service.CartService
}

func NewCartHandler(service service.CartService) *CartHandler {
	return &CartHandler{service: service}
}
func (h *CartHandler) RegisterRoutes(r *gin.Engine) {
	r.POST("/users/:id/cart/items", h.CreateItem)
	r.GET("/users/:id/cart", h.GetCart)
	r.PATCH("/users/:id/cart/items/:item_id", h.UpdateItem)
	r.DELETE("/users/:id/cart/items/:item_id", h.DeleteItem)
	r.DELETE("/users/:id/cart", h.DeleteCart)
}

func (h *CartHandler) GetCart(c *gin.Context) {
	idStr := c.Param("id")

	userID64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid user id"})
		return
	}
	userID := uint(userID64)

	cart, err := h.service.GetCart(userID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, cart)
}

func (h *CartHandler) CreateItem(c *gin.Context) {
	idStr := c.Param("id")

	userID64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid user id"})
		return
	}
	userID := uint(userID64)

	var req models.CartItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	cart, err := h.service.AddItemToCart(userID, req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, cart)
}

func (h *CartHandler) UpdateItem(c *gin.Context) {
	idStr := c.Param("id")
	itemIDStr := c.Param("item_id")
	userID64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid user id"})
		return
	}
	itemID64, err := strconv.ParseUint(itemIDStr, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid item id"})
		return
	}
	userID := uint(userID64)
	itemID := uint(itemID64)
	var req models.CartUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	cart, err := h.service.UpdateItemInCart(userID, itemID, req.Quantity)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, cart)
}

func (h *CartHandler) DeleteItem(c *gin.Context) {
	idStr := c.Param("id")
	itemIDStr := c.Param("item_id")
	userID64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid user id"})
		return
	}
	itemID64, err := strconv.ParseUint(itemIDStr, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid item id"})
		return
	}
	userID := uint(userID64)
	itemID := uint(itemID64)
	cart, err := h.service.RemoveItemFromCart(userID, itemID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, cart)
}

func (h *CartHandler) DeleteCart(c *gin.Context) {
	idStr := c.Param("id")	
	userID64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid user id"})
		return
	}
	userID := uint(userID64)
	err = h.service.DeleteCart(userID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "cart deleted successfully"})
}