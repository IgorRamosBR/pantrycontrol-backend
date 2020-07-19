package handlers

import (
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"net/http"
	"pantrycontrol-backend/internal/domain/models"
	"pantrycontrol-backend/internal/domain/services"
)

type ProductHandler struct {
	ProductService services.ProductService
}

func CreateProductHandler(productService services.ProductService) ProductHandler {
	return ProductHandler{ProductService: productService}
}

func (h *ProductHandler) SaveProduct(c echo.Context) error {
	product := models.Product{}

	err := c.Bind(&product)
	if err != nil {
		c.Logger().Error("Error to bind a product.")
		return c.JSON(http.StatusInternalServerError, err)
	}

	err = h.ProductService.SaveProduct(product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	c.Logger().Info("Product saved with success.")
	return c.NoContent(http.StatusCreated)
}

func (p *ProductHandler) FindProducts(c echo.Context) error {
	products, err := p.ProductService.FindProducts()
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, products)
}


