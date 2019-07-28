package model

import "time"

type Todo struct {
	ID      string
	Text    string
	Done    bool
	Created time.Time
	Deleted time.Time
}
