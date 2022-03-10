package models

import "time"

type Record struct {
	ID        string    `json:"id"`
	Lat       float32   `json:"lat"`
	Lng       float32   `json:"lng"`
	Timestamp time.Time `json:"timestamp"`
}
