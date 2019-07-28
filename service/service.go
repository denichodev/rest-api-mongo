package service

import (
	"github.com/denichodev/rest-api-mongo/repository"
)

// NewTodoService returns new instance of TodoService
func NewTodoService(repository repository.TodoRepositoryContract) *TodoService {
	return &TodoService{
		repository: repository,
	}
}
