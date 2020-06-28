package routes

import (
	"github.com/labstack/echo"
	"pantrycontrol-backend/internal/application/handlers"
)

func Route() *echo.Echo {
	e := echo.New()

	productHandler := handlers.ProductHandler{}

	e.POST("/product", productHandler.SaveProduct)
	e.GET("/product", productHandler.FindProducts)

	return e
}
