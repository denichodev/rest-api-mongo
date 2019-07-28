package handler

import "github.com/denichodev/rest-api-mongo/service"

func NewTodoHandler(
	todoService service.TodoServiceContract,
) *TodoHandler {
	return &TodoHandler{
		todoService: todoService,
	}
}
