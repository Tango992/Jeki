package repository

import (
	"merchant-service/dto"
	"merchant-service/model"
)

type Merchant interface{
	FindAllRestaurants() ([]model.Restaurant, error)
	FindRestaurantByID(uint32) (model.Restaurant, error)
	FindMenuById(id uint32) (model.Menu, error)
	FindMultipleMenuDetails([]int) ([]dto.MenuTmp, error)
	FindRestaurantIdByAdminId(uint) (uint, error)
	DeleteMenu(uint, uint) error
	CreateMenu(*model.Menu) error
	UpdateMenu(menu *dto.UpdateMenu) error
	CreateRestaurant(*model.Restaurant) error
	UpdateRestaurant(restaurant *model.Restaurant) (*model.Restaurant, error)
}