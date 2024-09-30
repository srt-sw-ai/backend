package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email              string `gorm:"type:varchar(100);uniqueIndex;not null"`
	Password           string `gorm:"type:varchar(255);not null"`
	NickName           string `gorm:"type:varchar(50);not null"`
	ImageUri           string `gorm:"type:varchar(255)"`
	HashedRefreshToken string `gorm:"type:varchar(255)"`
}
