package handler

import (
	"net/http"

	"github.com/denichodev/rest-api-mongo/object"
	"github.com/denichodev/rest-api-mongo/service"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type TodoHandler struct {
	todoService service.TodoServiceContract
}

func (h *TodoHandler) GetRoutes() chi.Router {
	r := chi.NewRouter()

	r.Post("/", h.Store)

	return r
}

func (h *TodoHandler) Store(w http.ResponseWriter, r *http.Request) {
	req := object.StoreTodoRequest{}

	if err := render.Bind(r, &req); err != nil {
		if err.Error() == "EOF" {
			render.Render(w, r, createUnprocessableEntityResponse("`text` field is required"))
		} else {
			render.Render(w, r, createUnprocessableEntityResponse(err.Error()))
		}
		return
	}

	todo, err := h.todoService.Create(req.Text)

	if err != nil {
		render.Render(w, r, createInternalServerErrorResponse(err.Error()))
		return
	}

	render.Render(w, r, &object.TodoCreatedResponse{
		Data: object.TodoResponse{
			ID:      todo.ID,
			Text:    todo.Text,
			Done:    todo.Done,
			Created: todo.Created.Format("2006-01-02 15:04:05"),
		},
		Status: "success",
	})
}
