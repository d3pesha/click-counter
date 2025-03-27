package api

import "time"

var GetStats struct {
	From time.Time `json:"from"`
	To   time.Time `json:"to"`
}
