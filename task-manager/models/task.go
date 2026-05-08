package models


import (
	"time"
)



type Task struct {
	ID int `json:"id" gorm:"primaryKey"`
	Title string `json:"title"`
	Description string `json:"description"`
	Done bool `json:"done" gorm:"default:false"`
	CreatedAt time.Time `json:"created_at"`
}

