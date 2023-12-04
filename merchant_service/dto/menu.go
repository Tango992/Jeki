package dto

type MenuTmp struct {
	ID    uint32
	Name  string
	Price float32
}

type UpdateMenu struct {
	ID         uint    `gorm:"primaryKey"`
	Name       string  `gorm:"not null"`
	CategoryId uint    `gorm:"not null"`
	Price      float32 `gorm:"not null"`
}