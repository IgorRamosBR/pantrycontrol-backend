package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"pantrycontrol-backend/internal/application/handlers"
	"pantrycontrol-backend/internal/domain/services"
)

func Route(productService services.ProductService) *echo.Echo {
	productHandler := handlers.CreateProductHandler(productService)
	listHandler := handlers.CreateListHandler()

	e := echo.New()
	e.Use(middleware.Logger())

	e.POST("/v1/products", productHandler.SaveProduct)
	e.GET("/v1/products", productHandler.FindProducts)
	e.GET("/v1/products/:id", productHandler.FindProductById)
	e.PUT("/v1/products/:id", productHandler.UpdateProduct)
	e.POST("/v1/shoppingLists", listHandler.SaveList)
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	return e
}
