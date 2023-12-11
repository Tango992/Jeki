package dto

import (
	"api-gateway/models"
	"api-gateway/pb/merchantpb"
	"api-gateway/pb/orderpb"
)

// Swagger Users
type SwaggerResponseRegister struct {
	Message string      `json:"message" extensions:"x-order=0"`
	Data    models.User `json:"data" extensions:"x-order=1"`
}

// Swagger Merchant
// For Customers

type SwaggerResponseGetAllCategories struct {
	Message string                 `json:"message" extensions:"x-order=0"`
	Data    []*merchantpb.Category `json:"data" extensions:"x-order=1"`
}
type SwaggerResponseGetAllRestaurant struct {
	Message string                                `json:"message" extensions:"x-order=0"`
	Data    *merchantpb.RestaurantCompactRepeated `json:"data" extensions:"x-order=1"`
}

type SwaggerResponseGetRestaurantByID struct {
	Message string                         `json:"message" extensions:"x-order=0"`
	Data    *merchantpb.RestaurantDetailed `json:"data" extensions:"x-order=1"`
}

type SwaggerResponseGetMenuById struct {
	Message string           `json:"message" extensions:"x-order=0"`
	Data    *merchantpb.Menu `json:"data" extensions:"x-order=1"`
}

// For Admin
type SwaggerRequestMenu struct {
	Name       string  `json:"name" extensions:"x-order=0"`
	Price      float32 `json:"price" extensions:"x-order=1"`
	CategoryID uint    `json:"category_id" extensions:"x-order=2"`
}

type SwaggerResponseGetRestaurantByAdminID struct {
	Message string                     `json:"message" extensions:"x-order=0"`
	Data    *merchantpb.RestaurantData `json:"data" extensions:"x-order=1"`
}

type SwaggerResponseCreateRestaurant struct {
	Message string                `json:"message" extensions:"x-order=0"`
	Data    ResponseNewRestaurant `json:"data" extensions:"x-order=1"`
}

type SwaggerResponseUpdateRestaurant struct {
	Message string                   `json:"message" extensions:"x-order=0"`
	Data    ResponseUpdateRestaurant `json:"data" extensions:"x-order=1"`
}

type SwaggerResponseGetMenuByAdminID struct {
	Message string                          `json:"message" extensions:"x-order=0"`
	Data    *merchantpb.MenuCompactRepeated `json:"data" extensions:"x-order=1"`
}

type SwaggerResponseGetMenuIdByAdminID struct {
	Message string                  `json:"message" extensions:"x-order=0"`
	Data    *merchantpb.MenuCompact `json:"data" extensions:"x-order=1"`
}

type SwaggerResponseCreateMenuByAdminID struct {
	Message string      `json:"message" extensions:"x-order=0"`
	Data    NewMenuData `json:"data" extensions:"x-order=1"`
}

type SwaggerResponseUpdateMenuByAdminID struct {
	Message string         `json:"message" extensions:"x-order=0"`
	Data    UpdateMenuData `json:"data" extensions:"x-order=1"`
}

// Swagger Orders
// For Users
type SwaggerResponseOrder struct {
	Message string         `json:"message" extensions:"x-order=0"`
	Data    *orderpb.Order `json:"data" extensions:"x-order=1"`
}

type SwaggerResponese struct {
	Message string          `json:"message" extensions:"x-order=0"`
	Data    *orderpb.Orders `json:"data" extensions:"x-order=1"`
}

// Update for Merchant & Driver
type SwaggerResponseUpdateOrder struct {
	Message string                     `json:"message" extensions:"x-order=0"`
	Data    *orderpb.RequestUpdateData `json:"data" extensions:"x-order=1"`
}

// For Drivers
type SwaggerResponseDriverGetAllOrders struct {
	Message string          `json:"message" extensions:"x-order=0"`
	Data    *orderpb.Orders `json:"data" extensions:"x-order=1"`
}

type SwaggerResponseDriverGetCurrentOrder struct {
	Message string         `json:"message" extensions:"x-order=0"`
	Data    *orderpb.Order `json:"data" extensions:"x-order=1"`
}
