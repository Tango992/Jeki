package controller

import (
	"context"
	"merchant-service/model"
	pb "merchant-service/pb/merchantpb"
	"merchant-service/repository"

	"google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
	pb.UnimplementedMerchantServer
	Repository repository.Merchant
}

func NewUserController(r repository.Merchant) Server {
	return Server{
		Repository: r,
	}
}

func (s Server) FindAllRestaurants(ctx context.Context, empty *emptypb.Empty) (*pb.RestaurantCompactRepeated, error) {
    restaurants, err := s.Repository.FindAllRestaurants()
    if err != nil {
        return nil, err
    }

    var pbRestaurantCompactRepeated []*pb.RestaurantCompact
    for _, r := range restaurants {
        pbRestaurant := &pb.RestaurantCompact{
            Id: uint32(r.ID),
			Name: r.Name,
			Address: r.Address,

        }
        pbRestaurantCompactRepeated = append(pbRestaurantCompactRepeated, pbRestaurant)
    }

    return &pb.RestaurantCompactRepeated{
        Restaurants: pbRestaurantCompactRepeated,
    }, nil
}

func (s Server) FindRestaurantById(ctx context.Context, idReq *pb.IdRestaurant) (*pb.RestaurantDetailed, error) {
    restaurantID := idReq.GetId()

    restaurant, err := s.Repository.FindRestaurantByID(restaurantID)
    if err != nil {
        return nil, err
    }

    pbRestaurantDetailed := convertToRestaurantDetailedProtoBuf(*restaurant)

    return pbRestaurantDetailed, nil
}

func convertToRestaurantDetailedProtoBuf(restaurant model.Restaurant) *pb.RestaurantDetailed {
    var pbMenus []*pb.Menu
    for _, menu := range restaurant.Menus {
        pbMenu := &pb.Menu{
            Id:           uint32(menu.ID),
            Name:         menu.Name,
            CategoryId:   uint32(menu.CategoryId),
            Price:        float32(menu.Price),
        }
        pbMenus = append(pbMenus, pbMenu)
    }

    pbRestaurantDetailed := &pb.RestaurantDetailed{
        Id:        uint32(restaurant.ID),
        Name:      restaurant.Name,
        Address:   restaurant.Address,
        Latitude:  float32(restaurant.Latitude),
        Longitude: float32(restaurant.Longitude),
        Menus:     pbMenus,
    }

    return pbRestaurantDetailed
}


func (s Server) CreateMenu(ctx context.Context, data *pb.NewMenuData) (*pb.MenuId, error) {
	restaurantId, err := s.Repository.FindRestaurantIdByAdminId(uint(data.AdminId))
	if err != nil {
		return nil, err
	}
	
	menuData := model.Menu{
		RestaurantId: restaurantId,
		Name: data.Name,
		CategoryId: uint(data.CategoryId),
		Price: data.Price,
	}

	if err := s.Repository.CreateMenu(&menuData); err != nil {
		return nil, err
	}
	return &pb.MenuId{Id: uint32(menuData.ID)}, nil
}

func (s Server) DeleteMenu(ctx context.Context, data *pb.AdminIdMenuId) (*emptypb.Empty, error) {
	restaurantId, err := s.Repository.FindRestaurantIdByAdminId(uint(data.AdminId))
	if err != nil {
		return nil, err
	}

	if err := s.Repository.DeleteMenu(restaurantId, uint(data.MenuId)); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s Server) CreateRestaurant(ctx context.Context, data *pb.NewRestaurantData) (*pb.IdRestaurant, error) {
	restaurantData := model.Restaurant{
		AdminId: uint(data.AdminId),
		Name: data.Name,
		Address: data.Address,
		Latitude: data.Latitude,
		Longitude: data.Longitude,
	}

	if err := s.Repository.CreateRestaurant(&restaurantData); err != nil {
		return nil, err
	}
	return &pb.IdRestaurant{Id: uint32(restaurantData.ID)}, nil
}



func (s Server) FindMenuById(ctx context.Context, data *pb.MenuId) (*pb.Menu, error) {
	return nil, nil
}

func (s Server) UpdateMenu(context.Context, *pb.UpdateMenuData) (*emptypb.Empty, error) {
	return nil, nil
}

func (s Server) UpdateRestaurant(ctx context.Context, data *pb.UpdateRestaurantData) (*emptypb.Empty, error) {
	return nil, nil
}

func (s Server) FindMenuDetailsWithSubtotal(ctx context.Context, data *pb.RequestMenuDetails) (*pb.ResponseMenuDetails, error) {
	var menuIds []int
	for _, val := range data.RequestMenuDetails {
		menuIds = append(menuIds, int(val.Id))
	}
	
	menuDatas, err := s.Repository.FindMultipleMenuDetails(menuIds)
	if err != nil {
		return nil, err
	}

	var responseDatas []*pb.ResponseMenuDetail
	for i, menu := range menuDatas {
		quantity := data.RequestMenuDetails[i].Qty
		subtotal := menu.Price * float32(quantity)

		menuData := &pb.ResponseMenuDetail{
			Id: uint32(menu.ID),
			Name: menu.Name,
			Qty: quantity,
			Subtotal: subtotal,
		}
		responseDatas = append(responseDatas, menuData)
	}

	pbResponseData := &pb.ResponseMenuDetails{
		ResponseMenuDetails: responseDatas,
	}
	
	return pbResponseData, nil
}

func (s Server) FindMenusByAdminId(ctx context.Context, data *pb.AdminId) (*pb.MenuCompactRepeated, error) {
	return nil, nil
}

func (s Server) FindOneMenuByAdminId(ctx context.Context, data *pb.AdminIdMenuId) (*pb.MenuCompact, error) {
	return nil, nil
}

func (s Server) FindRestaurantByAdminId(ctx context.Context, data *pb.AdminId) (*pb.RestaurantData, error) {
	return nil, nil
}
