package object

import (
	"errors"
	"net/http"
)

type StoreTodoRequest struct {
	Text string `json:"text"`
}

func (req *StoreTodoRequest) Bind(r *http.Request) error {
	if req.Text == "" {
		return errors.New("`text` is required")
	}

	return nil
}
