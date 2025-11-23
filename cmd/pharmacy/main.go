package main

import (
	"log"
	"pharmacy-team/internal/config"
	"pharmacy-team/internal/models"
)

func main() {
	db := config.SetUpDatabaseConnection()

	// Выполняем миграции моделей
	if err := db.AutoMigrate(
		&models.Payment{},
		&models.Pharmacy{},
		&models.Review{},
		&models.User{},
		&models.Category{},
		&models.SubCategory{},
		&models.Order{},
		&models.OrderItem{},
	); err != nil {
		log.Fatalf("не удалось выполнить миграции: %v", err)
	}
}
