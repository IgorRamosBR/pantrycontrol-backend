package main

import (
	"github.com/labstack/gommon/log"
	"pantrycontrol-backend/internal/application/routes"
	"pantrycontrol-backend/internal/domain/repository"
	"pantrycontrol-backend/internal/domain/services"
	"pantrycontrol-backend/internal/infra/configuration"
)

func main() {
	appConfig := configuration.CreateConfig()
	database := configuration.CreateDatabase(appConfig.DatabaseUrl)
	defer database.Close()

	productRepository := repository.CreateProductRepository(database)
	productService := services.CreateProductService(productRepository)

	router := routes.Route(productService)

	err := router.Start(":" + appConfig.Port)
	if err != nil {
		log.Error("Error to start application.", err)
	}
}

