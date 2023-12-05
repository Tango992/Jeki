package controller

import (
	"context"
	"merchant-service/model"
	pb "merchant-service/pb/merchantpb"
	"merchant-service/repository"
	"merchant-service/service"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
	pb.UnimplementedMerchantServer
	Repository repository.Merchant
	CachingService service.CachingService
}

func NewMerchantController(r repository.Merchant, cs service.CachingService) Server {
	return Server{
		Repository: r,
		CachingService: cs,
	}
}

func (s Server) CacheRestaurantDetailed(restaurantID uint32) error {
	restaurant, err := s.Repository.FindRestaurantByID(restaurantID)
	if err != nil {
		return err
	}

	menus, err := s.Repository.FindMenuByRestaurantId(restaurantID)
	if err != nil {
		return err
	}

	pbRestaurantData := &pb.RestaurantDetailed{
		Id:        restaurant.Id,
		Name:      restaurant.Name,
		Address:   restaurant.Address,
		Latitude:  restaurant.Latitude,
		Longitude: restaurant.Longitude,
		Menus:     menus,
	}

	if err := s.CachingService.SetRestaurantDetailed(uint(restaurantID), pbRestaurantData); err != nil {
		return err
	}
	return nil
}

func (s Server) FindAllRestaurants(ctx context.Context, empty *emptypb.Empty) (*pb.RestaurantCompactRepeated, error) {
	restaurants, err := s.Repository.FindAllRestaurants()
	if err != nil {
		return nil, err
	}

	var pbRestaurants []*pb.RestaurantCompact
	for _, r := range restaurants {
		pbRestaurant := &pb.RestaurantCompact{
			Id:      uint32(r.ID),
			Name:    r.Name,
			Address: r.Address,
		}
		pbRestaurants = append(pbRestaurants, pbRestaurant)
	}

	return &pb.RestaurantCompactRepeated{
		Restaurants: pbRestaurants,
	}, nil
}

func (s Server) FindRestaurantById(ctx context.Context, idReq *pb.IdRestaurant) (*pb.RestaurantDetailed, error) {
	restaurantID := idReq.GetId()

	result, err := s.CachingService.GetRestaurantDetailed(uint(restaurantID))
	if err != nil {
		return nil, err
	}

	if result != nil {
		return result, nil
	}
	
	restaurant, err := s.Repository.FindRestaurantByID(restaurantID)
	if err != nil {
		return nil, err
	}

	menus, err := s.Repository.FindMenuByRestaurantId(restaurantID)
	if err != nil {
		return nil, err
	}

	pbRestaurantData := &pb.RestaurantDetailed{
		Id:        restaurant.Id,
		Name:      restaurant.Name,
		Address:   restaurant.Address,
		Latitude:  restaurant.Latitude,
		Longitude: restaurant.Longitude,
		Menus:     menus,
	}

	if err := s.CachingService.SetRestaurantDetailed(uint(idReq.Id), pbRestaurantData); err != nil {
		return nil, err
	}
	return pbRestaurantData, nil
}

func (s Server) CreateMenu(ctx context.Context, data *pb.NewMenuData) (*pb.MenuId, error) {
	restaurantId, err := s.Repository.FindRestaurantIdByAdminId(data.AdminId)
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				return nil, status.Error(codes.FailedPrecondition, "please cerate a restaurant first before posting a menu")
			}
		}
		return nil, err
	}

	menuData := model.Menu{
		RestaurantId: restaurantId,
		Name:         data.Name,
		CategoryId:   uint(data.CategoryId),
		Price:        data.Price,
	}
	
	if err := s.Repository.CreateMenu(&menuData); err != nil {
		return nil, err
	}

	if err := s.CacheRestaurantDetailed(uint32(restaurantId)); err != nil {
		return nil, err
	}
	return &pb.MenuId{Id: uint32(menuData.ID)}, nil
}

func (s Server) DeleteMenu(ctx context.Context, data *pb.AdminIdMenuId) (*emptypb.Empty, error) {
	restaurantId, err := s.Repository.FindRestaurantIdByAdminId(data.AdminId)
	if err != nil {
		return nil, err
	}

	if err := s.Repository.DeleteMenu(restaurantId, uint(data.MenuId)); err != nil {
		return nil, err
	}

	if err := s.CacheRestaurantDetailed(uint32(restaurantId)); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s Server) CreateRestaurant(ctx context.Context, data *pb.NewRestaurantData) (*pb.IdRestaurant, error) {
	restaurantData := model.Restaurant{
		AdminId:   uint(data.AdminId),
		Name:      data.Name,
		Address:   data.Address,
		Latitude:  data.Latitude,
		Longitude: data.Longitude,
	}

	if err := s.Repository.CreateRestaurant(&restaurantData); err != nil {
		return nil, err
	}

	if err := s.CacheRestaurantDetailed(uint32(restaurantData.ID)); err != nil {
		return nil, err
	}
	return &pb.IdRestaurant{Id: uint32(restaurantData.ID)}, nil
}

func (s Server) FindMenuById(ctx context.Context, data *pb.MenuId) (*pb.Menu, error) {
	menu, err := s.Repository.FindMenuById(data.Id)
	if err != nil {
		return nil, err
	}
	return menu, nil
}

func (s Server) UpdateMenu(ctx context.Context, data *pb.UpdateMenuData) (*emptypb.Empty, error) {	
	restaurantAdminId, err := s.Repository.FindAdminIdByMenuId(data.MenuId)
	if err != nil {
		return nil, err
	}

	if restaurantAdminId != data.AdminId {
		return nil, status.Error(codes.PermissionDenied, "invalid admin access to edit menu data")
	}

	if err := s.Repository.UpdateMenu(data); err != nil {
		return nil, err
	}

	restaurantId, err := s.Repository.FindRestaurantIdByAdminId(data.AdminId)
	if err != nil {
		return nil, err
	}

	if err := s.CacheRestaurantDetailed(uint32(restaurantId)); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s Server) UpdateRestaurant(ctx context.Context, data *pb.UpdateRestaurantData) (*emptypb.Empty, error) {
	restaurantId, err := s.Repository.FindRestaurantIdByAdminId(data.AdminId)
	if err != nil {
		return nil, err
	}

	if err := s.Repository.UpdateRestaurant(restaurantId, data); err != nil {
		return nil, err
	}

	if err := s.CacheRestaurantDetailed(uint32(restaurantId)); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s Server) CalculateOrder(ctx context.Context, data *pb.RequestMenuDetails) (*pb.CalculateOrderResponse, error) {
	var (
		menuIds       []int
		menuIdWithQty = map[uint32]uint32{}
	)

	for _, val := range data.RequestMenuDetails {
		menuIdWithQty[val.Id] = val.Qty
		menuIds = append(menuIds, int(val.Id))
	}

	restaurantData, err := s.Repository.FindRestaurantMetadataByMenuIds(menuIds)
	if err != nil {
		return nil, err
	}

	menuDatas, err := s.Repository.FindMultipleMenuDetails(menuIds)
	if err != nil {
		return nil, err
	}

	pbMenuDetails := []*pb.ResponseMenuDetail{}
	for _, menu := range menuDatas {
		quantity := menuIdWithQty[menu.ID]
		subtotal := menu.Price * float32(quantity)

		menuData := &pb.ResponseMenuDetail{
			Id:       uint32(menu.ID),
			Name:     menu.Name,
			Qty:      uint32(quantity),
			Price:    menu.Price,
			Subtotal: subtotal,
		}
		pbMenuDetails = append(pbMenuDetails, menuData)
	}

	return &pb.CalculateOrderResponse{
		RestaurantData: restaurantData,
		ResponseMenuDetails: pbMenuDetails,
	}, nil
}

func (s Server) FindMenusByAdminId(ctx context.Context, data *pb.AdminId) (*pb.MenuCompactRepeated, error) {
	menus, err := s.Repository.FindMenusByAdminId(data.Id)
	if err != nil {
		return nil, err
	}

	return &pb.MenuCompactRepeated{
		Menus: menus,
	}, nil
}

func (s Server) FindOneMenuByAdminId(ctx context.Context, data *pb.AdminIdMenuId) (*pb.MenuCompact, error) {
	menu, err := s.Repository.FindOneMenuByAdminId(data.MenuId, data.AdminId)
	if err != nil {
		return nil, err
	}
	return menu, nil
}

func (s Server) FindRestaurantByAdminId(ctx context.Context, data *pb.AdminId) (*pb.RestaurantData, error) {
	restaurant, err := s.Repository.FindRestaurantByAdminId(data.Id)
	if err != nil {
		return nil, err
	}
	return restaurant, nil
}
