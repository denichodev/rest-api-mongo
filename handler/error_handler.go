package handler

import (
	"net/http"

	"github.com/go-chi/render"
)

type errorResponse struct {
	HTTPStatusCode int      `json:"-"`
	Status         string   `json:"status"`
	Errors         []string `json:"errors"`
}

func (err *errorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, err.HTTPStatusCode)
	return nil
}

func createBadRequestResponse(message string) *errorResponse {
	if message == "" {
		message = "Bad request"
	}

	return &errorResponse{
		Status:         "error",
		HTTPStatusCode: http.StatusBadRequest,
		Errors: []string{
			message,
		},
	}
}

func createUnprocessableEntityResponse(message string) *errorResponse {
	if message == "" {
		message = "Unprocessable entity"
	}

	return &errorResponse{
		Status:         "error",
		HTTPStatusCode: http.StatusUnprocessableEntity,
		Errors: []string{
			message,
		},
	}
}

func createInternalServerErrorResponse(message string) *errorResponse {
	if message == "" {
		message = "Internal server errro"
	}

	return &errorResponse{
		Status:         "error",
		HTTPStatusCode: http.StatusInternalServerError,
		Errors: []string{
			message,
		},
	}
}
