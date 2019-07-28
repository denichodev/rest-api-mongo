package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func CreateRouter() chi.Router {
	r := chi.NewRouter()

	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		})
	})

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

	return r
}
