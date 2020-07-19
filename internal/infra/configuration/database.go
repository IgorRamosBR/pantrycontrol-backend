package configuration

import (
	"context"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type DatabaseConfiguration struct {
	ctx    context.Context
	client *mongo.Client
}

func CreateDatabase() DatabaseConfiguration {
	ctx, _  := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Panic("Error to connect MongoDb")
	}

	return DatabaseConfiguration{ctx: ctx, client: client}
}

func (d *DatabaseConfiguration) Connect(name string) *mongo.Database {
	database := d.client.Database(name)
	log.Info("Connected to MongoDB!")
	return database
}

func (d *DatabaseConfiguration) Disconnect()  {
	d.Disconnect()
}