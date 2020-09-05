package services

import (
	"errors"
	"github.com/go-pg/pg/v10"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"pantrycontrol-backend/internal/application"
	"pantrycontrol-backend/internal/domain/dto"
	"pantrycontrol-backend/internal/domain/entities"
	mock_repository "pantrycontrol-backend/internal/domain/repository/mocks"
	"testing"
)


var id = primitive.NewObjectID()

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
	productRepository.EXPECT().FindById(id.String()).Return(product, nil)
	productService := CreateProductService(productRepository)

	productFound, err := productService.FindProductById(id.String())

	assert.Nil(t, err)
	assert.Equal(t, product, productFound)
}

func TestProductServiceImpl_FindProductById_NotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productRepository := mock_repository.NewMockProductRepository(ctrl)
	productRepository.EXPECT().FindById(id.String()).Return(entities.Product{}, pg.ErrNoRows)
	productService := CreateProductService(productRepository)

	_, err := productService.FindProductById(id.String())

	assert.NotNil(t, err)
	assert.Equal(t, application.ErrNotFound, err)
}

func TestProductServiceImpl_FindProductById_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productRepository := mock_repository.NewMockProductRepository(ctrl)
	productRepository.EXPECT().FindById(id.String()).Return(entities.Product{}, errors.New("Erro interno."))
	productService := CreateProductService(productRepository)

	_, err := productService.FindProductById(id.String())

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
	productRepository.EXPECT().Update(id.String(), product).Return(nil)
	productService := CreateProductService(productRepository)

	err := productService.UpdateProduct(id.String(), productDto)

	assert.Nil(t, err)
}

func TestProductServiceImpl_UpdateProduct_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := entities.Product{Name: "Arroz", Unit: "kg", Brand: "Carreteiro", Category: "Alimentos Básicos"}
	productDto := dto.ProductDTO{Name: "Arroz", Unit: "kg", Brand: "Carreteiro", Category: "Alimentos Básicos"}

	productRepository := mock_repository.NewMockProductRepository(ctrl)
	productRepository.EXPECT().Update(id.String(), product).Return(errors.New("Erro interno."))
	productService := CreateProductService(productRepository)

	err := productService.UpdateProduct(id.String(), productDto)

	assert.NotNil(t, err)
}

func TestProductServiceImpl_DeleteProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productRepository := mock_repository.NewMockProductRepository(ctrl)
	productRepository.EXPECT().Delete(id.String()).Return(nil)
	productService := CreateProductService(productRepository)

	err := productService.DeleteProduct(id.String())

	assert.Nil(t, err)
}

func TestProductServiceImpl_DeleteProduct_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productRepository := mock_repository.NewMockProductRepository(ctrl)
	productRepository.EXPECT().Delete(id.String()).Return(errors.New("Erro interno."))
	productService := CreateProductService(productRepository)

	err := productService.DeleteProduct(id.String())

	assert.NotNil(t, err)
}

func createProduct() entities.Product {
	return entities.Product{
		ID:		   id,
		Name:     "Arroz",
		Unit:     "kg",
		Brand:    "Carreteiro",
		Category: "Alimentos Básicos",
	}
}