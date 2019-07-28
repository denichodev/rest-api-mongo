package repository

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/denichodev/rest-api-mongo/model"
	"go.mongodb.org/mongo-driver/mongo"
)

// TodoRepository is the repository for Todo
type TodoRepository struct {
	db *mongo.Database
}

// Save stores new Todo to db
func (r *TodoRepository) Save(data *model.Todo) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	data.Created = time.Now()

	insertedOneResult, err := r.db.Collection("todo").InsertOne(ctx, data)

	if err != nil {
		return "", err
	}

	return insertedOneResult.InsertedID.(primitive.ObjectID).Hex(), nil
}
