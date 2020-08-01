package services

import (
	"errors"
	"github.com/go-pg/pg/v10"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"pantrycontrol-backend/internal/application"
	"pantrycontrol-backend/internal/domain/models/dto"
	"pantrycontrol-backend/internal/domain/models/entities"
	mock_repository "pantrycontrol-backend/internal/domain/repository/mocks"
	"testing"
)

func TestProductServiceImpl_FindProducts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := createProduct()
	produductRepository := mock_repository.NewMockProductRepository(ctrl)
	produductRepository.EXPECT().FindAll().Return([]entities.Product{product}, nil)
	productService := CreateProductService(produductRepository)

	products, err := productService.FindProducts()

	assert.Nil(t, err)
	assert.Equal(t, []entities.Product{product}, products)
}

func TestProductServiceImpl_FindProducts_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	produductRepository := mock_repository.NewMockProductRepository(ctrl)
	produductRepository.EXPECT().FindAll().Return(nil, errors.New("Erro interno."))
	productService := CreateProductService(produductRepository)

	products, err := productService.FindProducts()

	assert.NotNil(t, err)
	assert.Nil(t, products)
}

func TestProductServiceImpl_FindProductById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := createProduct()
	productRepository := mock_repository.NewMockProductRepository(ctrl)
	productRepository.EXPECT().FindById(123).Return(product, nil)
	productService := CreateProductService(productRepository)

	productFound, err := productService.FindProductById(123)

	assert.Nil(t, err)
	assert.Equal(t, product, productFound)
}

func TestProductServiceImpl_FindProductById_NotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productRepository := mock_repository.NewMockProductRepository(ctrl)
	productRepository.EXPECT().FindById(123).Return(entities.Product{}, pg.ErrNoRows)
	productService := CreateProductService(productRepository)

	_, err := productService.FindProductById(123)

	assert.NotNil(t, err)
	assert.Equal(t, application.ErrNotFound, err)
}

func TestProductServiceImpl_FindProductById_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productRepository := mock_repository.NewMockProductRepository(ctrl)
	productRepository.EXPECT().FindById(123).Return(entities.Product{}, errors.New("Erro interno."))
	productService := CreateProductService(productRepository)

	_, err := productService.FindProductById(123)

	assert.NotNil(t, err)
}

func TestProductServiceImpl_SaveProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := entities.Product{Name: "Arroz", Unit: "kg", Brand: "Carreteiro", Category: "Alimentos Básicos"}
	productDto := dto.ProductDTO{Name: "Arroz", Unit: "kg", Brand: "Carreteiro", Category: "Alimentos Básicos"}

	productRepository := mock_repository.NewMockProductRepository(ctrl)
	productRepository.EXPECT().Save(product).Return(nil)
	productService := CreateProductService(productRepository)

	err := productService.SaveProduct(productDto)

	assert.Nil(t, err)
}

func TestProductServiceImpl_SaveProduct_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := entities.Product{Name: "Arroz", Unit: "kg", Brand: "Carreteiro", Category: "Alimentos Básicos"}
	productDto := dto.ProductDTO{Name: "Arroz", Unit: "kg", Brand: "Carreteiro", Category: "Alimentos Básicos"}

	productRepository := mock_repository.NewMockProductRepository(ctrl)
	productRepository.EXPECT().Save(product).Return(errors.New("Erro interno."))
	productService := CreateProductService(productRepository)

	err := productService.SaveProduct(productDto)

	assert.NotNil(t, err)
}

func TestProductServiceImpl_UpdateProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := entities.Product{Name: "Arroz", Unit: "kg", Brand: "Carreteiro", Category: "Alimentos Básicos"}
	productDto := dto.ProductDTO{Name: "Arroz", Unit: "kg", Brand: "Carreteiro", Category: "Alimentos Básicos"}

	productRepository := mock_repository.NewMockProductRepository(ctrl)
	productRepository.EXPECT().Update(123, product).Return(nil)
	productService := CreateProductService(productRepository)

	err := productService.UpdateProduct(123, productDto)

	assert.Nil(t, err)
}

func TestProductServiceImpl_UpdateProduct_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := entities.Product{Name: "Arroz", Unit: "kg", Brand: "Carreteiro", Category: "Alimentos Básicos"}
	productDto := dto.ProductDTO{Name: "Arroz", Unit: "kg", Brand: "Carreteiro", Category: "Alimentos Básicos"}

	productRepository := mock_repository.NewMockProductRepository(ctrl)
	productRepository.EXPECT().Update(123, product).Return(errors.New("Erro interno."))
	productService := CreateProductService(productRepository)

	err := productService.UpdateProduct(123, productDto)

	assert.NotNil(t, err)
}

func TestProductServiceImpl_DeleteProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productRepository := mock_repository.NewMockProductRepository(ctrl)
	productRepository.EXPECT().Delete(123).Return(nil)
	productService := CreateProductService(productRepository)

	err := productService.DeleteProduct(123)

	assert.Nil(t, err)
}

func TestProductServiceImpl_DeleteProduct_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productRepository := mock_repository.NewMockProductRepository(ctrl)
	productRepository.EXPECT().Delete(123).Return(errors.New("Erro interno."))
	productService := CreateProductService(productRepository)

	err := productService.DeleteProduct(123)

	assert.NotNil(t, err)
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