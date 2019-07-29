package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/render"

	"github.com/denichodev/rest-api-mongo/db"
	"github.com/denichodev/rest-api-mongo/repository"

	"github.com/denichodev/rest-api-mongo/service"
	"github.com/go-chi/chi"
)

func CreateRouter() chi.Router {
	r := chi.NewRouter()

	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Get("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		payload := map[string]string{
			"name": "Test Mongo API",
		}

		res, err := json.Marshal(payload)

		if err != nil {
			fmt.Println(err)
		}

		w.Write(res)
	}))

	dbConn := db.Get()

	todoService := service.NewTodoService(repository.NewTodoRepository(dbConn.Database("todo-go")))
	todoHandler := NewTodoHandler(todoService)

	r.Mount("/todo", todoHandler.GetRoutes())

	return r
}
