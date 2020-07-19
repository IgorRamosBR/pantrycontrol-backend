package services

import (
	"context"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
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

func (p *ProductService) FindProductById(id string) (models.Product, error) {
	product := models.Product{}
	err := p.ProductRepository.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&product)
	if err != nil {
		log.Error("Error to find a product.", err)
		return models.Product{}, err
	}
	return product, err
}

func (p *ProductService) SaveProduct(product models.Product) error {
	_, err := p.ProductRepository.InsertOne(context.TODO(), product)
	if err != nil {
		log.Error("Error to save a product.", err)
		return err
	}
	return nil
}

func (p *ProductService) UpdateProduct(id string, product models.Product) error {
	_, err := p.ProductRepository.UpdateOne(context.TODO(), bson.M{"_id": id}, product)

	if err != nil {
		log.Error("Error to update a product.", err)
		return err
	}
	return nil
}

func (p *ProductService) DeleteProduct(id string) error {
	_, err := p.ProductRepository.DeleteOne(context.TODO(), bson.M{"_id": id})

	if err != nil {
		log.Error("Error to delete a product.", err)
		return err
	}
	return nil
}
