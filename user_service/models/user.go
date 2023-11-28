package models

type User struct {
	ID           uint   `json:"id" gorm:"primaryKey"`
	FirstName    string `json:"first_name" gorm:"not null"`
	LastName     string `json:"last_name"  gorm:"not null"`
	Email        string `json:"email"  gorm:"not null; unique"`
	Password     string `json:"password,omitempty" gorm:"not null"`
	BirthDate    string `json:"birth_date" gorm:"not null"`
	RoleID       uint32 `json:"role_id" gorm:"not null"`
	CreatedAt    string `json:"created_at" gorm:"not null"`
	Verification Verification
	Roles        []Role
}

type Verification struct {
	UserID   uint   `json:"id" gorm:"primaryKey"`
	Token    string `json:"token" gorm:"not null"`
	Validate bool   `json:"validate" gorm:"not null"`
}

type Role struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"not null"`
}
