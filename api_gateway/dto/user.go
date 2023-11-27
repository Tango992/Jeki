package dto

type UserRegister struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
	BirthDate string `json:"birth_date" validate:"required"`
}

type UserLogin struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
