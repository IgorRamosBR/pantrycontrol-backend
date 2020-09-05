package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Name     string             `bson:"name,omitempty"`
	Unit  	 string             `bson:"unit,omitempty"`
	Brand    string             `bson:"brand,omitempty"`
	Category string             `bson:"category,omitempty"`
}
