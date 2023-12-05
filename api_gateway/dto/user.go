package dto

type UserRegister struct {
	FirstName string `json:"first_name" validate:"required" extensions:"x-order=0"`
	LastName  string `json:"last_name" extensions:"x-order=1"`
	Email     string `json:"email" validate:"required,email" extensions:"x-order=2"`
	Password  string `json:"password" validate:"required" extensions:"x-order=3"`
	BirthDate string `json:"birth_date" validate:"required" extensions:"x-order=4"`
}

type UserLogin struct {
	Email    string `json:"email" validate:"required" extensions:"x-order=0"`
	Password string `json:"password" validate:"required" extensions:"x-order=1"`
}
