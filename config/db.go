package config

import (
	"fmt"
	"swai/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(config *Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DBUser, config.DBPassword, config.DBHost, config.DBName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&entity.User{}, &entity.Report{}, &entity.Map{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
