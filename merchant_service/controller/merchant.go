package controller

import (
	"context"
	"math"
	"merchant-service/dto"
	"merchant-service/model"
	pb "merchant-service/pb/merchantpb"
	"merchant-service/repository"
	"merchant-service/service"

	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
	pb.UnimplementedMerchantServer
	Repository     repository.Merchant
	CachingService service.CachingService
}

func NewMerchantController(r repository.Merchant, cs service.CachingService) Server {
	return Server{
		Repository:     r,
		CachingService: cs,
	}
}

// Gets restaurant data from database and cache service. Returns the restaurant data if needed and an error.
func (s Server) CacheRestaurantDetailed(restaurantID uint32) (*pb.RestaurantDetailed, error) {
	restaurantChan := make(chan dto.RestaurantDataCompact, 1)
	menusChan := make(chan []*pb.Menu, 1)
	g, _ := errgroup.WithContext(context.Background())

	g.Go(func() error {
		restaurant, err := s.Repository.FindRestaurantByID(restaurantID)
		if err != nil {
			return err
		}
		restaurantChan <- restaurant
		return nil
	})

	g.Go(func() error {
		menus, err := s.Repository.FindMenuByRestaurantId(restaurantID)
		if err != nil {
			return err
		}
		menusChan <- menus
		return nil
	})

	if err := g.Wait(); err != nil {
		return nil, err
	}

	close(restaurantChan)
	close(menusChan)

	restaurant := <-restaurantChan
	menus := <-menusChan

	pbRestaurantData := &pb.RestaurantDetailed{
		Id:        restaurant.Id,
		Name:      restaurant.Name,
		Address:   restaurant.Address,
		Latitude:  restaurant.Latitude,
		Longitude: restaurant.Longitude,
		Menus:     menus,
	}

	if err := s.CachingService.SetRestaurantDetailed(uint(restaurantID), pbRestaurantData); err != nil {
		return nil, err
	}
	return pbRestaurantData, nil
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

	restaurantData, err := s.CacheRestaurantDetailed(restaurantID)
	if err != nil {
		return nil, err
	}

	return restaurantData, nil
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

	if _, err := s.CacheRestaurantDetailed(uint32(restaurantId)); err != nil {
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

	if _, err := s.CacheRestaurantDetailed(uint32(restaurantId)); err != nil {
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

	if _, err := s.CacheRestaurantDetailed(uint32(restaurantData.ID)); err != nil {
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
	restaurantAdminIdChan := make(chan uint32, 1)
	restaurantIdChan := make(chan uint, 1)
	g, _ := errgroup.WithContext(ctx)

	g.Go(func() error {
		restaurantAdminId, err := s.Repository.FindAdminIdByMenuId(data.MenuId)
		if err != nil {
			return err
		}
		restaurantAdminIdChan <- restaurantAdminId
		return nil
	})

	g.Go(func() error {
		restaurantId, err := s.Repository.FindRestaurantIdByAdminId(data.AdminId)
		if err != nil {
			return err
		}
		restaurantIdChan <- restaurantId
		return nil
	})

	if err := g.Wait(); err != nil {
		return nil, err
	}

	close(restaurantAdminIdChan)
	close(restaurantIdChan)

	if <-restaurantAdminIdChan != data.AdminId {
		return nil, status.Error(codes.PermissionDenied, "invalid admin access to edit menu data")
	}

	if err := s.Repository.UpdateMenu(data); err != nil {
		return nil, err
	}

	if _, err := s.CacheRestaurantDetailed(uint32(<-restaurantIdChan)); err != nil {
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

	if _, err := s.CacheRestaurantDetailed(uint32(restaurantId)); err != nil {
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

	g, _ := errgroup.WithContext(ctx)
	restaurantDataChan := make(chan *pb.RestaurantMetadata, 1)
	menuDatasChan := make(chan []dto.MenuTmp, 1)

	g.Go(func() error {
		restaurantData, err := s.Repository.FindRestaurantMetadataByMenuIds(menuIds)
		if err != nil {
			return err
		}
		restaurantDataChan <- restaurantData
		return nil
	})

	g.Go(func() error {
		menuDatas, err := s.Repository.FindMultipleMenuDetails(menuIds)
		if err != nil {
			return err
		}
		menuDatasChan <- menuDatas
		return nil
	})

	if err := g.Wait(); err != nil {
		return nil, err
	}

	close(restaurantDataChan)
	close(menuDatasChan)

	pbMenuDetails := []*pb.ResponseMenuDetail{}
	for _, menu := range <-menuDatasChan {
		quantity := menuIdWithQty[menu.ID]
		subtotal := math.Round((float64(menu.Price) * float64(quantity)))

		menuData := &pb.ResponseMenuDetail{
			Id:       uint32(menu.ID),
			Name:     menu.Name,
			Qty:      uint32(quantity),
			Price:    menu.Price,
			Subtotal: float32(subtotal),
		}
		pbMenuDetails = append(pbMenuDetails, menuData)
	}

	return &pb.CalculateOrderResponse{
		RestaurantData:      <-restaurantDataChan,
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

func (s Server) FindAllCategories(ctx context.Context, data *emptypb.Empty) (*pb.Categories, error) {
	categories, err := s.Repository.FindAllCategories()
	if err != nil {
		return nil, err
	}
	return categories, nil
}
