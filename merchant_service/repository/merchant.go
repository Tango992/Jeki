package repository

import (
	"merchant-service/dto"
	"merchant-service/model"
	pb "merchant-service/pb/merchantpb"
)

type Merchant interface{
	FindAllRestaurants() ([]model.Restaurant, error)
	FindRestaurantByID(uint32) (dto.RestaurantDataCompact, error)
	FindMultipleMenuDetails([]int) ([]dto.MenuTmp, error)
	FindRestaurantIdByAdminId(uint32) (uint, error)
	FindAdminIdByMenuId(uint32) (uint32, error)
	FindMenuByRestaurantId(uint32) ([]*pb.Menu, error)
	FindMenuById(uint32) (*pb.Menu, error)
	FindMenusByAdminId(uint32) ([]*pb.MenuCompact, error)
	FindOneMenuByAdminId(menuID uint32, adminID uint32) (*pb.MenuCompact, error)
	FindRestaurantByAdminId(adminID uint32) (*pb.RestaurantData, error)
	UpdateMenu(*pb.UpdateMenuData) error
	UpdateRestaurant(restaurantId uint, data *pb.UpdateRestaurantData) error
	DeleteMenu(uint, uint) error
	CreateMenu(*model.Menu) error
	CreateRestaurant(*model.Restaurant) error
}