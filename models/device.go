package models

import "time"

type Device struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	LastPing time.Time `json:"lastPing"`
}
