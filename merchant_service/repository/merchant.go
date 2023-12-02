package repository

import (
	"merchant-service/dto"
	"merchant-service/model"
)

type Merchant interface{
	FindMultipleMenuDetails([]int) ([]dto.MenuTmp, error)
	FindRestaurantIdByAdminId(uint) (uint, error)
	DeleteMenu(uint, uint) error
	CreateMenu(*model.Menu) error
	CreateRestaurant(*model.Restaurant) error
}