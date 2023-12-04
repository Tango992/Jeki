package dto

type UpdateMenuData struct {
	MenuId     uint32  `json:"menu_id,omitempty"`
	Name       string  `json:"name" validate:"required"`
	CategoryId uint32  `json:"category_id" validate:"required"`
	Price      float32 `json:"price" validate:"required"`
}

type NewMenuData struct {
	ID         uint32  `json:"id,omitempty"`
	Name       string  `json:"name" validate:"required"`
	CategoryId uint32  `json:"category_id" validate:"required"`
	Price      float32 `json:"price" validate:"required"`
}

type ResponseNewRestaurant struct {
	ID        uint    `json:"id"`
	Name      string  `json:"name"`
	Address   string  `json:"address"`
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}

type ResponseUpdateRestaurant struct {
	Name      string  `json:"name"`
	Address   string  `json:"address"`
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}

type UpdateRestaurantData struct {
	Name    string `json:"name" validate:"required"`
	Address string `json:"address" validate:"required"`
}

type NewRestaurantData struct {
	Name    string `json:"name" validate:"required"`
	Address string `json:"address" validate:"required"`
}
