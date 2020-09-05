package repository

import (
	"context"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"pantrycontrol-backend/internal/domain/entities"
)

type ProductRepository interface {
	FindAll() ([]entities.Product, error)
	FindById(string) (entities.Product, error)
	Save(entities.Product) error
	Update(string, entities.Product) error
	Delete(string) error
}

type ProductRepositoryImpl struct {
	db *mongo.Collection
}

func CreateProductRepository(db *mongo.Collection) ProductRepository {
	return &ProductRepositoryImpl{db: db}
}

func (r *ProductRepositoryImpl) FindAll() ([]entities.Product, error) {
	var products []entities.Product
	ctx := context.TODO()

	cursor, err := r.db.Find(ctx, options.Find())
	if err != nil {
		log.Error(err)
		return nil, err
	}

	for cursor.Next(ctx) {
		var product entities.Product
		err := cursor.Decode(&product)
		if err != nil {
			log.Fatal(err)
		}

		products = append(products, product)
	}

	if err := cursor.Err(); err != nil {
		log.Error(err)
	}

	_ = cursor.Close(ctx)
	return products, nil
}

func (r *ProductRepositoryImpl) FindById(id string) (entities.Product, error) {
	product := entities.Product{}
	err := r.db.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&product)
	if err != nil {
		log.Error("Error to find a product.", err)
		return entities.Product{}, err
	}
	return product, err
}

func (r *ProductRepositoryImpl) Save(product entities.Product) error {
	_, err := r.db.InsertOne(context.TODO(), product)
	if err != nil {
		log.Error("Error to save a product.", err)
		return err
	}
	return nil
}

func (r *ProductRepositoryImpl) Update(id string, product entities.Product) error {
	_, err := r.db.UpdateOne(context.TODO(), bson.M{"_id": id}, product)

	if err != nil {
		log.Error("Error to update a product.", err)
		return err
	}
	return nil
}

func (r *ProductRepositoryImpl) Delete(id string) error {
	_, err := r.db.DeleteOne(context.TODO(), bson.M{"_id": id})

	if err != nil {
		log.Error("Error to delete a product.", err)
		return err
	}
	return nil
}