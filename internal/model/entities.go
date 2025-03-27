package model

import "time"

type Banner struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
type BannerClick struct {
	BannerID   int       `json:"-"`
	Timestamp  time.Time `json:"ts"`
	ClickCount int       `json:"v"`
}
