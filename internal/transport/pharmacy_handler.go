package transport

import (
	"net/http"
	"pharmacy-team/internal/models"
	"pharmacy-team/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PharmacyHandler struct {
	service service.PharmacyService
}

func NewPharmacyHandler(service service.PharmacyService) *PharmacyHandler {
	return &PharmacyHandler{service: service}
}
func (p *PharmacyHandler) RegisterRoutes(r *gin.Engine) {
	r.POST("/pharmacy", p.Create)
	r.GET("/pharmacy", p.Get)
	r.GET("/pharmacy/:id", p.GetByID)
	r.PATCH("/pharmacy/:id", p.Update)
	r.DELETE("/pharmacy/:id", p.Delete)
}

func (h *PharmacyHandler) Create(c *gin.Context) {
	var req models.PharmacyCreateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pharmacy, err := h.service.CreatePharmacy(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, pharmacy)
}

func (p *PharmacyHandler) Get(c *gin.Context) {
	pharmacy, err := p.service.GetPharmacy()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, pharmacy)
}
func (p *PharmacyHandler) GetByID(c *gin.Context) {

	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "неправильный id"})
		return
	}

	pharmacy, err := p.service.GetPharmacyByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, pharmacy)
}

func (p *PharmacyHandler) Update(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var req models.PharmacyUpdateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pharmacy, err := p.service.UpdatePharmacy(uint(id), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, pharmacy)

}

func (p *PharmacyHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if p.service.DeletePharmacy(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}
