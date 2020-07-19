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

func (h *ProductHandler) FindProducts(c echo.Context) error {
	products, err := h.ProductService.FindProducts()
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, products)
}

func (h *ProductHandler) FindProductById(c echo.Context) error {
	id := c.Param("id")
	product, err := h.ProductService.FindProductById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, product)
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

func (h *ProductHandler) UpdateProduct(c echo.Context) error {
	id := c.Param("id")
	product := models.Product{}

	err := c.Bind(&product)
	if err != nil {
		c.Logger().Error("Error to bind a product.")
		return c.JSON(http.StatusInternalServerError, err)
	}

	err = h.ProductService.UpdateProduct(id, product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	c.Logger().Info("Product updated with success.")
	return c.NoContent(http.StatusNoContent)
}

func (h *ProductHandler) DeleteProduct(c echo.Context) error {
	id := c.Param("id")

	err := h.ProductService.DeleteProduct(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	c.Logger().Info("Product deleted with success.")
	return c.NoContent(http.StatusNoContent)
}

