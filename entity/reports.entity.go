package entity

import (
	"gorm.io/gorm"
)

type Report struct {
	gorm.Model
	Type      string    `json:"type"`
	Title     string    `json:"title"`
	Content   string    `json:"content" gorm:"type:varchar(100)"`
	UserID    uint      `json:"userId" gorm:"column:user_id"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
}