package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var instance *mongo.Client

func initConnection() (*mongo.Client, error) {
	c, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	err = c.Connect(ctx)

	return c, err
}

// Get will return singleton of mongodb connection
func Get() *mongo.Client {
	if instance == nil {
		client, err := initConnection()

		if err != nil {
			panic(err.Error())
		}

		instance = client
	}

	return instance
}
