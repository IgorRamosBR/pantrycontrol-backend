package main

import (
	"github.com/labstack/gommon/log"
	"pantrycontrol-backend/internal/application/routes"
	"pantrycontrol-backend/internal/domain/services"
	"pantrycontrol-backend/internal/infra/configuration"
)

func main() {
	databaseConfig := configuration.CreateDatabase()
	database := databaseConfig.Connect("pantry")
	defer databaseConfig.Disconnect()

	productRepository := database.Collection("products")
	productService := services.CreateProductService(productRepository)

	router := routes.Route(productService)
	err := router.Start(":8080")
	if err != nil {
		log.Info("Error to start application.")
	}
}

