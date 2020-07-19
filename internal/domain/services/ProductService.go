package services

import (
	"context"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"pantrycontrol-backend/internal/domain/models"
)

type ProductService struct {
	ProductRepository *mongo.Collection
}

func CreateProductService(productRepository *mongo.Collection) ProductService {
	return ProductService{ProductRepository: productRepository}
}

func (p *ProductService) SaveProduct(product models.Product) error {
	_, err := p.ProductRepository.InsertOne(context.TODO(), product)
	if err != nil {
		log.Error("Error to save a product.")
		return err
	}
	return nil
}

func (p *ProductService) FindProducts() ([]*models.Product, error) {
	var products []*models.Product

	cursor, err := p.ProductRepository.Find(context.TODO(), options.Find())
	if err != nil {
		log.Error(err)
		return nil, err
	}

	for cursor.Next(context.TODO()) {
		var product models.Product
		err := cursor.Decode(&product)
		if err != nil {
			log.Fatal(err)
		}

		products = append(products, &product)
	}


	if err := cursor.Err(); err != nil {
		log.Error(err)
	}

	_ = cursor.Close(context.TODO())
	return products, nil
}