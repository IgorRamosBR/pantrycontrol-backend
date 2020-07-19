package routes

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"pantrycontrol-backend/internal/application/handlers"
	"pantrycontrol-backend/internal/domain/services"
)

func Route(productService services.ProductService) *echo.Echo {
	productHandler := handlers.CreateProductHandler(productService)

	e := echo.New()
	e.Use(middleware.Logger())


	e.POST("/products", productHandler.SaveProduct)
	e.GET("/products", productHandler.FindProducts)

	return e
}
