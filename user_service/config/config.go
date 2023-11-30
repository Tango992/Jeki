package config

import (
	"fmt"
	"user-service/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	dsn := "user=postgres dbname=deploy host=localhost password=secret port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	if err := db.AutoMigrate(&models.Role{}, &models.User{}, &models.Verification{}, &models.DriverStatus{}, &models.Driver{}); err != nil {
		return nil, err
	}
	
	return db, nil
}
