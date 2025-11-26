package transport

import (
	"pharmacy-team/internal/models"
	"pharmacy-team/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PaymentHandler struct {
	service service.PaymentService
}

func NewPaymentHandler(s service.PaymentService) *PaymentHandler {
	return &PaymentHandler{service: s}
}

func (h *PaymentHandler) RegisterRoutes(r *gin.Engine) {
	r.POST("/orders/:id/payments", h.CreatePayment)
	r.GET("/orders/:id/payments", h.GetPaymentsByOrder)
	r.GET("/payments/:id", h.GetPaymentById)
}

func (h *PaymentHandler) CreatePayment(c *gin.Context) {
	orderID, _ := strconv.Atoi(c.Param("id"))

	var req models.PaymentCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	payment, err := h.service.Create(uint(orderID), req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, payment)
}

func (h *PaymentHandler) GetPaymentsByOrder(c *gin.Context) {
	orderID, _ := strconv.Atoi(c.Param("id"))

	payments, err := h.service.ListByOrder(uint(orderID))
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, payments)
}

func (h *PaymentHandler) GetPaymentById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	payment, err := h.service.Get(uint(id))
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, payment)
}
