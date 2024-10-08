package entity

import (
	"time"
)

type User struct {
    ID                uint       `gorm:"primaryKey" json:"id"`
    CreatedAt         time.Time  `json:"created_at"`
    UpdatedAt         time.Time  `json:"updated_at"`
    DeletedAt         *time.Time `gorm:"index" json:"deleted_at,omitempty"`
    Email             string     `gorm:"type:varchar(100);uniqueIndex;not null" json:"email"`
    Password          string     `gorm:"type:varchar(255);not null" json:"password"`
    NickName          string     `gorm:"type:varchar(50);not null" json:"nickname"`
    ImageUri          string     `gorm:"type:varchar(255)" json:"image_uri"`
    HashedRefreshToken string    `gorm:"type:varchar(255)" json:"hashed_refresh_token"`
}