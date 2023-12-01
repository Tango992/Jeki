package model

type Restaurant struct {
	ID        uint    `gorm:"primaryKey"`
	AdminId   uint    `gorm:"not null;unique"`
	Name      string  `gorm:"not null"`
	Address   string  `gorm:"not null"`
	Latitude  float32 `gorm:"not null"`
	Longitude float32 `gorm:"not null"`
	Menus     []Menu
}
