package models

import "time"

type Post struct {
	ID        uint   `gorm:"primaryKey"`
	Author    string `json:"author"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	CreatedAt time.Time
}
