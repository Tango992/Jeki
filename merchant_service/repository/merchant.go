package repository

import (
	"merchant-service/dto"
	"merchant-service/model"
)

type Merchant interface{
	FindAllRestaurants() ([]model.Restaurant, error)
	FindRestaurantByID(id uint32) (*model.Restaurant, error)
	FindMultipleMenuDetails([]int) ([]dto.MenuTmp, error)
	FindRestaurantIdByAdminId(uint) (uint, error)
	DeleteMenu(uint, uint) error
	CreateMenu(*model.Menu) error
	CreateRestaurant(*model.Restaurant) error
}