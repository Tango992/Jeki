package dto

type UserJoinedData struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password,omitempty"`
	BirthDate string `json:"birth_date"`
	Role      string `json:"role_id"`
	CreatedAt string `json:"created_at"`
	Verified  bool   `json:"verified"`
}
