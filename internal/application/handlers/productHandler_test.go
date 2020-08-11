package handlers

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"pantrycontrol-backend/internal/application"
	"pantrycontrol-backend/internal/domain/models/dto"
	"pantrycontrol-backend/internal/domain/models/entities"
	mock_services "pantrycontrol-backend/internal/domain/services/mocks"
	"strings"
	"testing"
)

func TestProductHandler_FindProducts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productService := mock_services.NewMockProductService(ctrl)
	productHandler := CreateProductHandler(productService)
	productService.EXPECT().FindProducts().Return([]entities.Product{createProduct()}, nil)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/products", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, productHandler.FindProducts(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestProductHandler_FindProducts_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productService := mock_services.NewMockProductService(ctrl)
	productHandler := CreateProductHandler(productService)
	productService.EXPECT().FindProducts().Return(nil, errors.New("Erro interno."))

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/products", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, productHandler.FindProducts(c)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	}
}

func TestProductHandler_FindProductById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productService := mock_services.NewMockProductService(ctrl)
	productService.EXPECT().FindProductById(123).Return(createProduct(), nil)
	productHandler := CreateProductHandler(productService)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/products/:id")
	c.SetParamNames("id")
	c.SetParamValues("123")

	if assert.NoError(t, productHandler.FindProductById(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestProductHandler_FindProductById_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productService := mock_services.NewMockProductService(ctrl)
	productService.EXPECT().FindProductById(123).Return(entities.Product{}, errors.New("Erro interno."))
	productHandler := CreateProductHandler(productService)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/products/:id")
	c.SetParamNames("id")
	c.SetParamValues("123")

	if assert.NoError(t, productHandler.FindProductById(c)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	}
}

func TestProductHandler_FindProductById_WrongId(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productService := mock_services.NewMockProductService(ctrl)
	productHandler := CreateProductHandler(productService)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/products/:id")
	c.SetParamNames("id")
	c.SetParamValues("123a")

	if assert.NoError(t, productHandler.FindProductById(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

func TestProductHandler_FindProductById_Id_NotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productService := mock_services.NewMockProductService(ctrl)
	productService.EXPECT().FindProductById(123).Return(entities.Product{}, application.ErrNotFound)
	productHandler := CreateProductHandler(productService)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/products/:id")
	c.SetParamNames("id")
	c.SetParamValues("123")

	if assert.NoError(t, productHandler.FindProductById(c)) {
		assert.Equal(t, http.StatusNotFound, rec.Code)
	}
}

func TestProductHandler_SaveProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productService := mock_services.NewMockProductService(ctrl)
	product := `{"name":"Feijao","unit":"kg","brand":"Maximo","category":"Alimentos Basicos"}`
	productDTO := dto.ProductDTO{
		Name:     "Feijao",
		Unit:     "kg",
		Brand:    "Maximo",
		Category: "Alimentos Basicos",
	}
	productService.EXPECT().SaveProduct(productDTO).Return(nil)
	productHandler := CreateProductHandler(productService)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(product))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/products")

	if assert.NoError(t, productHandler.SaveProduct(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
	}
}

func TestProductHandler_SaveProduct_BadRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productService := mock_services.NewMockProductService(ctrl)
	product := `{"name":1,"unit":"kg","brand":"Maximo","category":"Alimentos Basicos"}`
	productHandler := CreateProductHandler(productService)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(product))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/products")

	if assert.NoError(t, productHandler.SaveProduct(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

func TestProductHandler_SaveProduct_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productService := mock_services.NewMockProductService(ctrl)
	product := `{"name":"Feijao","unit":"kg","brand":"Maximo","category":"Alimentos Basicos"}`
	productDTO := dto.ProductDTO{
		Name:     "Feijao",
		Unit:     "kg",
		Brand:    "Maximo",
		Category: "Alimentos Basicos",
	}
	productService.EXPECT().SaveProduct(productDTO).Return(errors.New("Erro interno."))
	productHandler := CreateProductHandler(productService)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(product))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/products")

	if assert.NoError(t, productHandler.SaveProduct(c)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	}
}

func TestProductHandler_UpdateProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := `{"name":"Feijao","unit":"kg","brand":"Maximo","category":"Alimentos Basicos"}`
	productDTO := dto.ProductDTO{
		Name:     "Feijao",
		Unit:     "kg",
		Brand:    "Maximo",
		Category: "Alimentos Basicos",
	}

	productService := mock_services.NewMockProductService(ctrl)
	productService.EXPECT().UpdateProduct(123, productDTO).Return(nil)
	productHandler := CreateProductHandler(productService)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(product))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/products/:id")
	c.SetParamNames("id")
	c.SetParamValues("123")

	if assert.NoError(t, productHandler.UpdateProduct(c)) {
		assert.Equal(t, http.StatusNoContent, rec.Code)
	}
}

func TestProductHandler_UpdateProduct_WrongId(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := `{"name":"Feijao","unit":"kg","brand":"Maximo","category":"Alimentos Basicos"}`

	productService := mock_services.NewMockProductService(ctrl)
	productHandler := CreateProductHandler(productService)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(product))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/products/:id")
	c.SetParamNames("id")
	c.SetParamValues("123b")

	if assert.NoError(t, productHandler.UpdateProduct(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

func TestProductHandler_UpdateProduct_WrongBody(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := `{"name":1,"unit":"kg","brand":"Maximo","category":"Alimentos Basicos"}`

	productService := mock_services.NewMockProductService(ctrl)
	productHandler := CreateProductHandler(productService)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(product))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/products/:id")
	c.SetParamNames("id")
	c.SetParamValues("123")

	if assert.NoError(t, productHandler.UpdateProduct(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

func TestProductHandler_UpdateProduct_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := `{"name":"Feijao","unit":"kg","brand":"Maximo","category":"Alimentos Basicos"}`
	productDTO := dto.ProductDTO{
		Name:     "Feijao",
		Unit:     "kg",
		Brand:    "Maximo",
		Category: "Alimentos Basicos",
	}

	productService := mock_services.NewMockProductService(ctrl)
	productService.EXPECT().UpdateProduct(123, productDTO).Return(errors.New("Erro interno."))
	productHandler := CreateProductHandler(productService)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(product))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/products/:id")
	c.SetParamNames("id")
	c.SetParamValues("123")

	if assert.NoError(t, productHandler.UpdateProduct(c)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	}
}

func TestProductHandler_DeleteProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productService := mock_services.NewMockProductService(ctrl)
	productService.EXPECT().DeleteProduct(123).Return(nil)
	productHandler := CreateProductHandler(productService)

	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/products/:id")
	c.SetParamNames("id")
	c.SetParamValues("123")

	if assert.NoError(t, productHandler.DeleteProduct(c)) {
		assert.Equal(t, http.StatusNoContent, rec.Code)
	}
}

func TestProductHandler_DeleteProduct_WrongId(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productService := mock_services.NewMockProductService(ctrl)
	productHandler := CreateProductHandler(productService)

	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/products/:id")
	c.SetParamNames("id")
	c.SetParamValues("123b")

	if assert.NoError(t, productHandler.DeleteProduct(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

func TestProductHandler_DeleteProduct_NotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productService := mock_services.NewMockProductService(ctrl)
	productService.EXPECT().DeleteProduct(123).Return(application.ErrNotFound)
	productHandler := CreateProductHandler(productService)

	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/products/:id")
	c.SetParamNames("id")
	c.SetParamValues("123")

	if assert.NoError(t, productHandler.DeleteProduct(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

func TestProductHandler_DeleteProduct_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productService := mock_services.NewMockProductService(ctrl)
	productService.EXPECT().DeleteProduct(123).Return(errors.New("Erro interno."))
	productHandler := CreateProductHandler(productService)

	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/products/:id")
	c.SetParamNames("id")
	c.SetParamValues("123")

	if assert.NoError(t, productHandler.DeleteProduct(c)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	}
}

func createProduct() entities.Product {
	return entities.Product{
		Id:		   123,
		Name:     "Arroz",
		Unit:     "kg",
		Brand:    "Carreteiro",
		Category: "Alimentos Básicos",
	}
}
