package main

import (
	"github.com/labstack/gommon/log"
	"pantrycontrol-backend/internal/application/routes"
	"pantrycontrol-backend/internal/domain/repository"
	"pantrycontrol-backend/internal/domain/services"
	"pantrycontrol-backend/internal/infra/configuration"
)

// @title Pantry Control Backend
// @version 1.0
// @description Documentation from pantry-control-backend.
// @termsOfService http://swagger.io/terms/

// @contact.name Igor Pestana
// @contact.email igorir7@gmail.com
// @BasePath /v1

func main() {
	databaseConfig := configuration.CreateDatabase()
	database := databaseConfig.Connect("pantry")
	defer databaseConfig.Disconnect()

	productCollection := database.Collection("products")

	productRepository := repository.CreateProductRepository(productCollection)
	productService := services.CreateProductService(productRepository)

	router := routes.Route(productService)

	err := router.Start(":8080")
	if err != nil {
		log.Error("Error to start application.", err)
	}
}

