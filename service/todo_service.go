package service

import (
	"github.com/denichodev/rest-api-mongo/model"
	"github.com/denichodev/rest-api-mongo/repository"
)

type TodoService struct {
	repository repository.TodoRepositoryContract
}

func (s *TodoService) Create(text string) (model.Todo, error) {
	payload := model.Todo{
		Text: text,
	}

	id, err := s.repository.Save(&payload)

	if err != nil {
		return model.Todo{}, err
	}

	payload.ID = id

	return payload, nil
}

func (s *TodoService) Get() ([]model.Todo, error) {
	payload, err := s.repository.Get()

	if err != nil {
		return []model.Todo{}, err
	}

	return payload, nil
}
