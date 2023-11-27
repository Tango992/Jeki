package models

type User struct {
	ID        uint32 `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	BirthDate string `json:"birth_date"`
	Role      string `json:"role"`
	CreatedAt string `json:"created_at"`
}
