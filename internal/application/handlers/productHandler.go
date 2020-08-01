package handlers

import (
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/pkg/errors"
	"net/http"
	"pantrycontrol-backend/internal/application"
	"pantrycontrol-backend/internal/domain/models/dto"
	"pantrycontrol-backend/internal/domain/services"
	"strconv"
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
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "id must be a number.")
	}
	product, err := h.ProductService.FindProductById(id)
	if err != nil {
		if errors.Cause(err) == application.ErrNotFound {
			return c.JSON(http.StatusBadRequest, "Product not found.")
		}
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) SaveProduct(c echo.Context) error {
	productDTO := dto.ProductDTO{}

	err := c.Bind(&productDTO)
	if err != nil {
		c.Logger().Error("Error to bind a productDTO.")
		return c.JSON(http.StatusBadRequest, err)
	}

	err = h.ProductService.SaveProduct(productDTO)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	c.Logger().Info("Product saved with success.")
	return c.NoContent(http.StatusCreated)
}

func (h *ProductHandler) UpdateProduct(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "id must be a number.")
	}
	productDTO := dto.ProductDTO{}

	err = c.Bind(&productDTO)
	if err != nil {
		c.Logger().Error("Error to bind a productDTO.")
		return c.JSON(http.StatusBadRequest, err)
	}

	err = h.ProductService.UpdateProduct(id, productDTO)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	c.Logger().Info("Product updated with success.")
	return c.NoContent(http.StatusNoContent)
}

func (h *ProductHandler) DeleteProduct(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "id must be a number.")
	}

	err = h.ProductService.DeleteProduct(id)
	if err != nil {
		if errors.Cause(err) == application.ErrNotFound {
			return c.JSON(http.StatusBadRequest, "Product not found.")
		}
		return c.JSON(http.StatusInternalServerError, err)
	}

	c.Logger().Info("Product deleted with success.")
	return c.NoContent(http.StatusNoContent)
}

