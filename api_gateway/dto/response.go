package dto

type Response struct {
	Message string `json:"message" extensions:"x-order=0"`
	Data    any    `json:"data" extensions:"x-order=1"`
}
