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
	c.JSON(400, gin.H{"error": "Еще не реализовано"})
}

func (h *UserHandler) GetUserCart(c *gin.Context) {
	c.JSON(400, gin.H{"error": "Еще не реализовано"})
}

func (h *UserHandler) RegisterRoutes(r *gin.Engine) {
	users := r.Group("/users")
	{
		users.POST("", h.CreateUser)
		users.GET("/:id", h.GetUserByID)
		users.PATCH("/:id", h.UpdateUser)
		users.DELETE("/:id", h.DeleteUser)
		users.GET("", h.GetAllUsers)

		users.GET("/:id/orders", h.GetUserOrders)
		users.GET("/:id/cart", h.GetUserCart)
	}
}
