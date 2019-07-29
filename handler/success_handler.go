package handler

import (
	"net/http"

	"github.com/go-chi/render"
)

type SuccessResponse struct {
	HTTPStatusCode int         `json:"-"`
	Data           interface{} `json:"data"`
	Status         string      `json:"status"`
}

type SuccessListResponse struct {
	HTTPStatusCode int           `json:"-"`
	Data           []interface{} `json:"data"`
	Status         string        `json:"status"`
}

func (res *SuccessResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, res.HTTPStatusCode)
	res.Status = "success"
	return nil
}

func (res *SuccessListResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, res.HTTPStatusCode)
	res.Status = "success"
	return nil
}
