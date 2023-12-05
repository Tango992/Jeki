package dto

type UpdateMenuData struct {
	MenuId     uint32  `json:"menu_id,omitempty" extensions:"x-order=0"`
	Name       string  `json:"name" validate:"required" extensions:"x-order=1"`
	CategoryId uint32  `json:"category_id" validate:"required" extensions:"x-order=2"`
	Price      float32 `json:"price" validate:"required" extensions:"x-order=3"`
}

type NewMenuData struct {
	ID         uint32  `json:"id,omitempty" extensions:"x-order=0"`
	Name       string  `json:"name" validate:"required" extensions:"x-order=1"`
	CategoryId uint32  `json:"category_id" validate:"required" extensions:"x-order=2"`
	Price      float32 `json:"price" validate:"required" extensions:"x-order=3"`
}

type ResponseNewRestaurant struct {
	ID        uint    `json:"id" extensions:"x-order=0"`
	Name      string  `json:"name" extensions:"x-order=1"`
	Address   string  `json:"address" extensions:"x-order=2"`
	Latitude  float32 `json:"latitude" extensions:"x-order=3"`
	Longitude float32 `json:"longitude" extensions:"x-order=4"`
}

type ResponseUpdateRestaurant struct {
	Name      string  `json:"name" extensions:"x-order=1"`
	Address   string  `json:"address" extensions:"x-order=2"`
	Latitude  float32 `json:"latitude" extensions:"x-order=3"`
	Longitude float32 `json:"longitude" extensions:"x-order=4"`
}

type UpdateRestaurantData struct {
	Name    string `json:"name" validate:"required" extensions:"x-order=0"`
	Address string `json:"address" validate:"required" extensions:"x-order=1"`
}

type NewRestaurantData struct {
	Name    string `json:"name" validate:"required" extensions:"x-order=0"`
	Address string `json:"address" validate:"required" extensions:"x-order=1"`
}
