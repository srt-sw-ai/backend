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
    Gender            string     `gorm:"type:varchar(10);not null" json:"gender"`
    Birthday          string     `gorm:"type:varchar(10);not null" json:"birthday"`
    Phone             string     `gorm:"type:varchar(20);not null" json:"phone"`
    EmergencyPhone    string     `gorm:"type:varchar(20);not null" json:"emergency_phone"`
    Address           string     `gorm:"type:varchar(255);not null" json:"address"`
    Allergys          string     `gorm:"type:varchar(255);not null" json:"allergys"`
    UnderlyingDiseases string     `gorm:"type:varchar(255);not null" json:"underlying_diseases"`
    Medicines         string     `gorm:"type:varchar(255);not null" json:"medicines"`
    BloodType         string     `gorm:"type:varchar(10);not null" json:"blood_type"`
    Weight            string     `gorm:"type:varchar(10);not null" json:"weight"`
    Height            string     `gorm:"type:varchar(10);not null" json:"height"`
}