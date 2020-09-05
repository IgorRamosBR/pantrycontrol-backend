package services

import (
	"github.com/go-pg/pg/v10"
	"github.com/labstack/gommon/log"
	"github.com/pkg/errors"
	"pantrycontrol-backend/internal/application"
	"pantrycontrol-backend/internal/domain/dto"
	"pantrycontrol-backend/internal/domain/entities"
	"pantrycontrol-backend/internal/domain/repository"
)

type ProductService interface {
	FindProducts() ([]entities.Product, error)
	FindProductById(id string) (entities.Product, error)
	SaveProduct(productDTO dto.ProductDTO) error
	UpdateProduct(id string, productDTO dto.ProductDTO) error
	DeleteProduct(id string) error
}

type ProductServiceImpl struct {
	ProductRepository repository.ProductRepository
}

func CreateProductService(productRepository repository.ProductRepository) ProductService {
	return &ProductServiceImpl{ProductRepository: productRepository}
}

func (p *ProductServiceImpl) FindProducts() ([]entities.Product, error) {
	var products []entities.Product

	products, err := p.ProductRepository.FindAll()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return products, nil
}

func (p *ProductServiceImpl) FindProductById(id string) (entities.Product, error) {
	product, err := p.ProductRepository.FindById(id)
	if err != nil {
		if errors.Cause(err) == pg.ErrNoRows {
			err = application.ErrNotFound
			log.Warn("Product not found.", err)
			return entities.Product{}, err
		}

		log.Error("Error to find a product.", err)
		return entities.Product{}, err
	}
	return product, err
}

func (p *ProductServiceImpl) SaveProduct(productDTO dto.ProductDTO) error {
	product := productDTO.ToProduct()
	err := p.ProductRepository.Save(product)
	if err != nil {
		log.Error("Error to save a product.", err)
		return err
	}
	return nil
}

func (p *ProductServiceImpl) UpdateProduct(id string, productDTO dto.ProductDTO) error {
	product := productDTO.ToProduct()
	err := p.ProductRepository.Update(id, product)

	if err != nil {
		log.Error("Error to update a product.", err)
		return err
	}
	return nil
}

func (p *ProductServiceImpl) DeleteProduct(id string) error {
	err := p.ProductRepository.Delete(id)
	if err != nil {
		log.Error("Error to delete a product.", err)
		return err
	}
	return nil
}
