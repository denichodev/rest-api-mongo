package repository

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
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

	fmt.Println("inserted", insertedOneResult.InsertedID.(primitive.ObjectID).Hex())

	return insertedOneResult.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (r *TodoRepository) Get() ([]model.Todo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	todoResult := []model.Todo{}

	cur, err := r.db.Collection("todo").Find(ctx, bson.D{})
	if err != nil {
		fmt.Println("error", err)
		return nil, err
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			return nil, err
		}

		resultDateTime := result["created"].(primitive.DateTime)
		resultTime := time.Unix(int64(resultDateTime)/1000, int64(resultDateTime)%1000*1000000)

		if err != nil {
			return nil, err
		}

		// convert map to json
		todo := model.Todo{
			ID:      result["_id"].(primitive.ObjectID).Hex(),
			Text:    result["text"].(string),
			Done:    result["done"].(bool),
			Created: resultTime,
		}

		todoResult = append(todoResult, todo)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return todoResult, nil
}
