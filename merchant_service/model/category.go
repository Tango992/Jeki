package model

type Category struct {
	Id    uint   `gorm:"primaryKey"`
	Name  string `gorm:"not null"`
	Menus []Menu
}
