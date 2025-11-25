package transport

import (
	"pharmacy-team/internal/models"
	"pharmacy-team/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PaymentHandler struct {
	servicePayment service.PaymentService
}

func NewPaymentHandler(s service.PaymentService) *PaymentHandler {
	return &PaymentHandler{servicePayment: s}
}

func (t *PaymentHandler) RegisterRoutes(r *gin.Engine) {
	r.POST("/orders/:id/payments", t.CreatePayment)
	r.GET("/orders/:id/payments", t.GetPaymentsByOrder)
	r.GET("/payments/:id", t.GetPaymentById)
}

func (h *PaymentHandler) CreatePayment(c *gin.Context) {
	orderID, _ := strconv.Atoi(c.Param("id"))

	var req models.PaymentCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	payment, err := h.servicePayment.Create(uint(orderID), req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, payment)
}

func (h *PaymentHandler) GetPaymentsByOrder(c *gin.Context) {
	orderID, _ := strconv.Atoi(c.Param("id"))

	payments, err := h.servicePayment.ListByOrder(uint(orderID))
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, payments)
}
func (t *PaymentHandler) GetPaymentById(c *gin.Context) {
	id := c.Param("id")

	paymentID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid payment id"})
		return
	}

	payment, err := t.servicePayment.Get(uint(paymentID))
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, payment)
}
