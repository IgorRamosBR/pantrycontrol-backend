package handlers

import (
	"encoding/json"
	_ "pantrycontrol-backend/docs"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/pkg/errors"
	"net/http"
	"pantrycontrol-backend/internal/application"
	"pantrycontrol-backend/internal/domain/dto"
	"pantrycontrol-backend/internal/domain/services"
)

type ProductHandler struct {
	ProductService services.ProductService
}

func CreateProductHandler(productService services.ProductService) ProductHandler {
	return ProductHandler{ProductService: productService}
}

// FindProducts godoc
// @Summary Find products
// @Produce json
// @Success 200 {array} entities.Product
// @Failure 400 {object} dto.Error
// @Failure 500 {object} dto.Error "When a internal error occur."
// @Router /products [get]
// @Tags products
func (h *ProductHandler) FindProducts(c echo.Context) error {
	products, err := h.ProductService.FindProducts()
	if err != nil {
		log.Error(err)
		return createErrorResponse(c, http.StatusInternalServerError, "Internal server error.")
	}
	return c.JSON(http.StatusOK, products)
}

// FindProducts godoc
// @Summary Find product by id
// @Produce json
// @Param id path string true "Product id"
// @Success 200 {object} entities.Product
// @Failure 400 {object} dto.Error "When request bad formatted"
// @Failure 404 {object} dto.Error "When not found a product"
// @Router /products/{id} [get]
// @Tags products
func (h *ProductHandler) FindProductById(c echo.Context) error {
	id := c.Param("id")
	product, err := h.ProductService.FindProductById(id)
	if err != nil {
		if errors.Cause(err) == application.ErrNotFound {
			return createErrorResponse(c, http.StatusNotFound, "Product not found.")
		}
		return createErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, product)
}

// FindProducts godoc
// @Summary Save Product
// @Produce json
// @Param product body dto.ProductDTO true "Create product"
// @Success 201
// @Failure 400 {object} dto.Error "When request bad formatted"
// @Failure 500 {object} dto.Error "When a internal error occur."
// @Router /products [post]
// @Tags products
func (h *ProductHandler) SaveProduct(c echo.Context) error {
	productDTO := dto.ProductDTO{}

	err := c.Bind(&productDTO)
	if err != nil {
		c.Logger().Error("Error to bind a productDTO.")
		return createErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	err = h.ProductService.SaveProduct(productDTO)
	if err != nil {
		return createErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.Logger().Info("Product saved with success.")
	return c.NoContent(http.StatusCreated)
}

// FindProducts godoc
// @Summary Update Product
// @Produce json
// @Param id path string true "Product id"
// @Success 204
// @Failure 400 {object} dto.Error "When request bad formatted"
// @Failure 500 {object} dto.Error "When a internal error occur."
// @Router /products/{id} [put]
// @Tags products
func (h *ProductHandler) UpdateProduct(c echo.Context) error {
	id := c.Param("id")
	productDTO := dto.ProductDTO{}

	err := c.Bind(&productDTO)
	if err != nil {
		c.Logger().Error("Error to bind a productDTO.")
		return createErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	err = h.ProductService.UpdateProduct(id, productDTO)
	if err != nil {
		return createErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.Logger().Info("Product updated with success.")
	return c.NoContent(http.StatusNoContent)
}

// FindProducts godoc
// @Summary Delete product
// @Produce json
// @Param id path string true "Product id"
// @Success 204
// @Failure 400 {object} dto.Error "When request bad formatted"
// @Failure 500 {object} dto.Error "When a internal error occur."
// @Router /products/{id} [delete]
// @Tags products
func (h *ProductHandler) DeleteProduct(c echo.Context) error {
	id := c.Param("id")

	err := h.ProductService.DeleteProduct(id)
	if err != nil {
		if errors.Cause(err) == application.ErrNotFound {
			return createErrorResponse(c, http.StatusBadRequest, "Product not found.")
		}
		return createErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.Logger().Info("Product deleted with success.")
	return c.NoContent(http.StatusNoContent)
}

func createErrorResponse(c echo.Context, code int, message string) error {
	response, _ := json.Marshal(dto.Error{Message: message})
	return c.JSON(code, response)
}