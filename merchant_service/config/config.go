package config

import (
	"fmt"
	"merchant-service/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	dsn := "user=postgres dbname=jeki_merchants host=localhost password=secret port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	if err := db.AutoMigrate(&model.Restaurant{}, &model.Category{}, &model.Menu{}); err != nil {
		return nil, err
	}
	
	return db, nil
}
