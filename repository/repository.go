package repository

import "go.mongodb.org/mongo-driver/mongo"

// NewTodoRepository creates new NewsRepository instance.
func NewTodoRepository(db *mongo.Database) *TodoRepository {
	return &TodoRepository{
		db: db,
	}
}
