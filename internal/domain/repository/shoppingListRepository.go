package repository

import (
	"context"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"pantrycontrol-backend/internal/domain/entities"
)

type ShoppingListRepository interface {
	FindAll() ([]entities.ShoppingList, error)
	FindById(string) (entities.ShoppingList, error)
	Save(entities.ShoppingList) error
	Update(string, entities.ShoppingList) error
	Delete(string) error
}

type ShoppingListRepositoryImpl struct {
	db *mongo.Collection
}

func CreateProductRepository(db *mongo.Collection) ShoppingListRepository {
	return &ShoppingListRepositoryImpl{db: db}
}

func (r *ShoppingListRepositoryImpl) FindAll() ([]entities.ShoppingList, error) {
	var shoppingLists []entities.ShoppingList
	ctx := context.TODO()

	cursor, err := r.db.Find(ctx, options.Find())
	if err != nil {
		log.Error(err)
		return nil, err
	}

	for cursor.Next(ctx) {
		var shoppingList entities.ShoppingList
		err := cursor.Decode(&shoppingList)
		if err != nil {
			log.Fatal(err)
		}

		shoppingLists = append(shoppingLists, shoppingList)
	}

	if err := cursor.Err(); err != nil {
		log.Error(err)
	}

	_ = cursor.Close(ctx)
	return shoppingLists, nil
}

func (r *ShoppingListRepositoryImpl) FindById(id string) (entities.ShoppingList, error) {
	shoppingList := entities.ShoppingList{}
	err := r.db.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&shoppingList)
	if err != nil {
		log.Error("Error to find a shoppingList.", err)
		return entities.ShoppingList{}, err
	}
	return shoppingList, err
}

func (r *ShoppingListRepositoryImpl) Save(shoppingList entities.ShoppingList) error {
	_, err := r.db.InsertOne(context.TODO(), shoppingList)
	if err != nil {
		log.Error("Error to save a shoppingList.", err)
		return err
	}
	return nil
}

func (r *ShoppingListRepositoryImpl) Update(id string, shoppingList entities.ShoppingList) error {
	_, err := r.db.UpdateOne(context.TODO(), bson.M{"_id": id}, shoppingList)

	if err != nil {
		log.Error("Error to update a shoppingList.", err)
		return err
	}
	return nil
}

func (r *ShoppingListRepositoryImpl) Delete(id string) error {
	_, err := r.db.DeleteOne(context.TODO(), bson.M{"_id": id})

	if err != nil {
		log.Error("Error to delete a shoppingList.", err)
		return err
	}
	return nil
}
