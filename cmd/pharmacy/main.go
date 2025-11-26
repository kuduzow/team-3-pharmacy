package main

import (
	"log"
	"pharmacy-team/internal/config"
	"pharmacy-team/internal/models"
	"pharmacy-team/internal/repository"
	"pharmacy-team/internal/service"
	"pharmacy-team/internal/transport"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.SetUpDatabaseConnection()

	// Выполняем миграции моделей
	if err := db.AutoMigrate(
		&models.Category{},
		&models.SubCategory{},
		&models.User{},
		&models.Pharmacy{},
		&models.Review{},
		&models.Order{},
		&models.OrderItem{},
		&models.Payment{},
		&models.Cart{},
	); err != nil {
		log.Fatalf("не удалось выполнить миграции: %v", err)
	}

	CategoryRepo := repository.NewCategoryRepository(db)
	OrderRepo := repository.NewOrderRepository(db)
	PaymentRepo := repository.NewPaymentRepository(db)
	PharmacyRepo := repository.NewPharmacyRepository(db)
	PromocodeRepo := repository.NewPromocodeRepository(db)
	RewievsRepo := repository.NewReviewRepository(db)
	SubCategoryRepo := repository.NewSubCategoryRepository(db)
	UserRepo := repository.NewUserRepository(db)
	CartRepo := repository.NewCartRepository(db)

	CategoryServ := service.NewCategoryService(CategoryRepo)
	OrderServ := service.NewOrderService(OrderRepo,CartRepo)
	PaymentServ := service.NewPaymentService(PaymentRepo, OrderRepo)
	PharmacyServ := service.NewPharmacyService(PharmacyRepo)
	PromocodeServ := service.NewPromocodeService(PromocodeRepo)
	RewievsServ := service.NewReviewService(RewievsRepo, PharmacyRepo)
	SubCategoryServ := service.NewSubCategoryService(SubCategoryRepo)
	UserServ := service.NewUserService(UserRepo)
	CartServ := service.NewCartService(CartRepo, PharmacyRepo)

	router := gin.Default()

	transport.RegisterRoutes(
		router,
		CategoryServ,
		OrderServ,
		PaymentServ,
		PharmacyServ,
		PromocodeServ,
		RewievsServ,
		SubCategoryServ,
		UserServ,
		CartServ,
	)
	router.Run(":8888")

}
