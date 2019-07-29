package repository

import (
	"github.com/denichodev/rest-api-mongo/model"
)

// TodoRepositoryContract represents interface
// for TodoRepository
type TodoRepositoryContract interface {
	Save(data *model.Todo) (string, error)
	Get() ([]model.Todo, error)
}
