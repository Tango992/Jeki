package models

type Driver struct {
	UserID         uint `gorm:"primaryKey"`
	DriverStatusID uint `gorm:"not null"`
}

type DriverStatus struct {
	ID      uint   `gorm:"primaryKey"`
	Status  string `gorm:"not null"`
	Drivers []Driver
}
