package entity

import (
	"time"

	"gorm.io/gorm"
)

type Report struct {
	gorm.Model
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Type      string    `json:"type"`
	Title     string    `json:"title"`
	Content   string    `json:"content" gorm:"type:varchar(100)"`
	Location  string    `json:"location"`
	Date      string    `json:"date"`
	UserID    uint      `json:"userId" gorm:"column:user_id"`
	CreatedAt time.Time `gorm:"column:createdat" json:"createdAt"`
}
