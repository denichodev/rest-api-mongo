package service

import "github.com/denichodev/rest-api-mongo/model"

type TodoServiceContract interface {
	Create(text string) (model.Todo, error)
}
