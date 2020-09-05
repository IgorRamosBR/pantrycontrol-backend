package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type ShoppingList struct {
	ID           primitive.ObjectID  `bson:"_id,omitempty"`
	Products     []Product           `bson:"products"`
	ShoppingDate primitive.Timestamp `bson:"shoppingDate,omitempty"`
	Market       Market              `bson:"market"`
}

type Market struct {
	Name string `bson:"name"`
	Logo string `bson:"logo"`
}
