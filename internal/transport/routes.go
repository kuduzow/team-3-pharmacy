package transport

import (
	"pharmacy-team/internal/service"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	router *gin.Engine,
	categoryServ service.CategoryService,
	orderServ service.OrderService,
	paymentServ service.PaymentService,
	pharmacyServ service.PharmacyService,
	promocodeServ service.PromocodeService,
	rewievsServ service.ReviewService,
	subcategoryServ service.SubCategoryService,
	userServ service.UserService,
) {
	categoryHandler := NewCategoryHandler(categoryServ)
	orderHandler := NewOrderHandler(orderServ)
	paymentHandler := NewPaymentHandler(paymentServ)
	pharmacyHandler := NewPharmacyHandler(pharmacyServ)
	promocodeHandler := NewPromoHandler(promocodeServ)
	reviewsHandler := NewReviewsHandler(rewievsServ)
	subcategoryHandler := NewSubCategoryHandler(subcategoryServ)
	userHandler := NewUserHandler(userServ)

	categoryHandler.RegisterRoutes(router)
	orderHandler.Register(router)
	paymentHandler.RegisterRoutes(router)
	pharmacyHandler.RegisterRoutes(router)
	promocodeHandler.RegisterRoutes(router)
	reviewsHandler.Routes(router)
	subcategoryHandler.RegisterRoutes(router)
	userHandler.RegisterRoutes(router)
}
