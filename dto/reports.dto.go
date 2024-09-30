package dto

import "time"

type CreateReportDto struct {
	Type     string `json:"type" validate:"required"`
	Title    string `json:"title" validate:"required"`
	Content  string `json:"content" validate:"required"`
	Location string `json:"location" validate:"required"`
	Date     string `json:"date" validate:"required"`
}

type ReportResponseDto struct {
	ID        uint      `json:"id"`
	Type      string    `json:"type"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Location  string    `json:"location"`
	Date      string    `json:"date"`
	UserID    uint      `json:"userId"`
	CreatedAt time.Time `json:"createdAt"`
}
