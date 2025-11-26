package transport

import (
	"net/http"
	"pharmacy-team/internal/models"
	"pharmacy-team/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var req models.UserCreate

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"err": "Неверные данные"})
		return
	}

	user, err := h.service.CreateUser(req)
	if err != nil {
		c.JSON(400, gin.H{
			"err": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"err": "некорректный идентификатор"})
		return
	}

	user, err := h.service.GetByUser(uint(id))
	if err != nil {
		c.JSON(400, gin.H{
			"err": "пользователь не найден",
		})
		return
	}

	c.JSON(200, user)

}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 64)

	if err != nil || id == 0 {
		c.JSON(400, gin.H{
			"err": "некорректный идентификатор"})
		return
	}

	var req models.UserUpdate

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"err": "Неверные данные"})
		return
	}

	user, err := h.service.UpdateUser(uint(id), req)
	if err != nil {
		c.JSON(400, gin.H{
			"err": "пользователь не найден",
		})
		return
	}

	c.JSON(200, user)

}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 64)

	if err != nil || id == 0 {
		c.JSON(400, gin.H{
			"err": "некорректный идентификатор"})
		return
	}

	if err := h.service.DeleteUser(uint(id)); err != nil {
		c.JSON(400, gin.H{
			"err": "Пользователь не найден",
		})
		return
	}

	c.JSON(200, gin.H{
		"massage": "Пользователь удален",
	})

}

func (h *UserHandler) GetAllUsers(c *gin.Context) {

	users, err := h.service.GetUsers()
	if err != nil {
		c.JSON(400, gin.H{
			"err": "Пользователи не найдены",
		})
		return
	}

	c.JSON(200, users)
}

func (h *UserHandler) GetUserOrders(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Неверный ID пользователя"})
		return
	}

	_, err = h.service.GetByUser(uint(id))
	if err != nil {
		c.JSON(404, gin.H{"error": "Пользователь не найден"})
		return
	}

	// Временная реализация с примером данных
	c.JSON(200, gin.H{
		"user_id": id,
		"orders": []gin.H{
			{
				"id":          1,
				"status":      "completed",
				"total_price": 15000,
				"final_price": 15000,
				"created_at":  "2024-01-15T10:30:00Z",
				"items_count": 3,
			},
		},
	})
}

func (h *UserHandler) GetUserCart(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Неверный ID пользователя"})
		return
	}

	// Проверяем существование пользователя
	_, err = h.service.GetByUser(uint(id))
	if err != nil {
		c.JSON(404, gin.H{"error": "Пользователь не найден"})
		return
	}

	// Временная реализация с примером данных
	c.JSON(200, gin.H{
		"user_id": id,
		"items": []gin.H{
			{
				"medicine_id":    1,
				"name":           "Парацетамол",
				"quantity":       2,
				"price_per_unit": 5000,
				"line_total":     10000,
			},
		},
		"total_price": 10000,
		"total_items": 2,
	})

}

func (h *UserHandler) RegisterRoutes(r *gin.Engine) {
	users := r.Group("/users")
	{
		users.POST("", h.CreateUser)
		users.GET("/:id", h.GetUserByID)
		users.PATCH("/:id", h.UpdateUser)
		users.DELETE("/:id", h.DeleteUser)
		users.GET("", h.GetAllUsers)

		users.GET("/:id/cart", h.GetUserCart)
	}
}
