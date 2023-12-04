package helpers

import (
	"api-gateway/dto"
	"api-gateway/pb/merchantpb"
)

func AssertRestaurantDetailed(data *merchantpb.RestaurantDetailed) dto.RestaurantDataById {
	restaurantData := dto.RestaurantDataById{
		ID: uint(data.Id),
		Name: data.Name,
		Address: data.Address,
		Latitude: data.Latitude,
		Longitude: data.Longitude,
	}

	var menus []dto.Menu
	for _, menu := range data.Menus {
		menuTmp := dto.Menu{
			ID: uint(menu.Id),
			Name: menu.Name,
			Category: menu.Category,
			Price: menu.Price,
		}
		menus = append(menus, menuTmp)
	}
	restaurantData.Menus = menus

	return restaurantData
}