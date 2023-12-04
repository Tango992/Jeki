package config

import (
	"fmt"
	"os"
	"user-service/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	dsn := os.Getenv("POSTGRES_URI")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	if err := db.AutoMigrate(&models.Role{}, &models.User{}, &models.Verification{}, &models.DriverStatus{}, &models.Driver{}); err != nil {
		return nil, err
	}
	
	return db, nil
}
