package model

type Menu struct {
	ID           uint    `gorm:"primaryKey"`
	RestaurantId uint    `gorm:"not null"`
	Name         string  `gorm:"not null"`
	CategoryId   uint    `gorm:"not null"`
	Price        float32 `gorm:"not null"`
}
