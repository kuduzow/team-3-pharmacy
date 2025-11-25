package transport

import (
	"pharmacy-team/internal/models"
	"pharmacy-team/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PromoHandler struct {
	service service.PromocodeService
}

func NewPromoHandler(service service.PromocodeService) *PromoHandler {
	return &PromoHandler{service: service}
}

func (h *PromoHandler) Create(c *gin.Context) {
	var req models.PromocodeCreate

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"err": "Неверные данные",
		})
		return
	}

	promo, err := h.service.CreatePromocode(req)

	if err != nil {
		c.JSON(400, gin.H{
			"err": err.Error()})
		return
	}

	c.JSON(200, promo)

}

func (h *PromoHandler) GetByPromocode(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"err": "Неверный ID",
		})
		return
	}

	promo, err := h.service.GetByPromocode(uint(id))

	if err != nil {
		c.JSON(400, gin.H{
			"err": "Нет такого промокода",
		})
		return
	}

	c.JSON(200, promo)

}

func (h *PromoHandler) UpdatePromocode(c *gin.Context) {
	idstr := c.Param("id")

	id, err := strconv.ParseUint(idstr, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"err": "Неверный ID",
		})
		return
	}

	var req models.PromocodeUpdate

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"err": "Неверный запрос",
		})
		return
	}

	promo, err := h.service.UpdatePromocode(uint(id), req)

	if err != nil {
		c.JSON(400, gin.H{
			"err": "Ошибка при изменении",
		})
		return
	}

	c.JSON(200, promo)
}

func (h *PromoHandler) DeletePromocode(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 64)

	if err != nil {
		c.JSON(400, gin.H{
			"err": "Неверный ID",
		})
		return
	}

	if err := h.service.DeletePromocode(uint(id)); err != nil {
		c.JSON(400, gin.H{
			"error": "Промокод не найден",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "промокод удален",
	})
}

func (h *PromoHandler) GetAllPromocode(c *gin.Context) {

	result, err := h.service.GetAllPromocode()
	if err != nil {
		c.JSON(400, gin.H{
			"err": "Промокодов нету",
		})
		return
	}

	c.JSON(200, result)
}

func (h *PromoHandler) RegisterRoutes(r *gin.Engine) {
	promocodes := r.Group("/promocodes")
	{
		promocodes.POST("", h.Create)
		promocodes.GET("", h.GetAllPromocode)
		promocodes.GET("/:id", h.GetByPromocode)
		promocodes.PUT("/:id", h.UpdatePromocode)
		promocodes.DELETE("/:id", h.DeletePromocode)
	}
}
