package transport

import (
	"net/http"
	"pharmacy-team/internal/models"
	"pharmacy-team/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SubCategoryHandler struct {
	service service.SubCategoryService
}

func NewSubCategoryHandler(service service.SubCategoryService) *SubCategoryHandler {
	return &SubCategoryHandler{service: service}
}

func (s *SubCategoryHandler) RegisterRoutes(r *gin.Engine) {
	r.GET("/category/:id/subcategory")
	r.POST("category/:id/subcategory")
}

func (s *SubCategoryHandler) Get(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	categoryID := uint(id)

	subcategory, err := s.service.GetSubCategory(categoryID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, subcategory)

}

func (s *SubCategoryHandler) Create(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	CategoryID := uint(id)

	var req models.SubCategoryCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	create, err := s.service.CreateSubCategory(CategoryID, req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, create)
}
