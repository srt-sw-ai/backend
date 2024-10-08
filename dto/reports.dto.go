package dto

type CreateReportDto struct {
	Type      string  `json:"type" validate:"required"`
	Title     string  `json:"title" validate:"required"`
	Content   string  `json:"content" validate:"required"`
	Date      string  `json:"date" validate:"required"`
	Latitude  float64 `json:"latitude" validate:"required"`
	Longitude float64 `json:"longitude" validate:"required"`
}

// type ReportResponseDto struct {
// 	ID        uint      `json:"id"`
// 	Type      string    `json:"type"`
// 	Title     string    `json:"title"`
// 	Content   string    `json:"content"`
// 	Latitude  float64 `json:"latitude"`
// 	Longitude float64 `json:"longitude"`
// 	Date      string    `json:"date"`
// 	UserID    uint      `json:"userId"`
// 	CreatedAt time.Time `json:"createdAt"`
// }
