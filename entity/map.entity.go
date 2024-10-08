package entity

import "gorm.io/gorm"

type Map struct {
	gorm.Model
	Type      string  `json:"type"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	ReportID  int     `json:"reportId"`
	UserID    uint    `json:"userId"`
}