package object

import (
	"net/http"

	"github.com/go-chi/render"
)

type TodoResponse struct {
	ID      string `json:"id"`
	Text    string `json:"text"`
	Done    bool   `json:"done"`
	Created string `json:"created"`
}

type TodoCreatedResponse struct {
	Data   TodoResponse `json:"data"`
	Status string       `json:"status"`
}

func (res *TodoCreatedResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, http.StatusCreated)
	return nil
}
