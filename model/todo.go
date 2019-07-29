package model

import (
	"time"
)

type Todo struct {
	ID      string
	Text    string    `bson:"text"`
	Done    bool      `bson:"done"`
	Created time.Time `bson:"created"`
	Deleted time.Time `bson:"deleted"`
}
