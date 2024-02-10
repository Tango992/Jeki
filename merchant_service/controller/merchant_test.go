package controller_test

import (
	"context"
	"merchant-service/controller"
	"merchant-service/dto"
	"merchant-service/model"
	pb "merchant-service/pb/merchantpb"
	"merchant-service/repository"
	"merchant-service/service"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	mockRepository = repository.NewMockMerchantRepository()
	mockCachingService = service.NewMockCachingService()
	merchantController = controller.NewMerchantController(&mockRepository, &mockCachingService)
)

var (
	mockMenu = &pb.Menu{
		Id: 1,
		Name: "Mock menu",
		Category: "Main course",
		Price: 10000,
	}

	mockRestaurantCompact = dto.RestaurantDataCompact{
		Id: 1,
		Name: "Sederhana",
		Address: "Jalan Sudirman",
		Latitude: 0.123,
		Longitude: 0.123,
	}
)

func TestMain(m *testing.M) {
	m.Run()
}

func TestCacheRestaurantDetailed(t *testing.T) {
	var (
		restaurantID uint32 = 1
	)

	var mockMenus []*pb.Menu
	mockMenus = append(mockMenus, mockMenu)

	mockRepository.Mock.On("FindRestaurantByID", restaurantID).Return(mockRestaurantCompact, nil)

	mockRepository.Mock.On("FindMenuByRestaurantId", restaurantID).Return(mockMenus, nil)

	mockCachingService.Mock.On("SetRestaurantDetailed", uint(restaurantID), mock.Anything).Return(nil)
	
	_, err := merchantController.CacheRestaurantDetailed(restaurantID)
	assert.Nil(t, err)
}

func TestFindAllRestaurants(t *testing.T) {
	mockRestaurant := model.Restaurant{
		ID: 1,
		Name: "John Doe",
		Address: "Jalan Sudirman",
	}
	mockRestaurants := []model.Restaurant{mockRestaurant}
	
	mockRepository.Mock.On("FindAllRestaurants").Return(mockRestaurants, nil)

	pbResponse, err := merchantController.FindAllRestaurants(context.Background(), &emptypb.Empty{})
	assert.Nil(t, err)
	assert.NotEmpty(t, pbResponse)
}

func TestFindRestaurantById(t *testing.T) {
	var mockRestaurantID uint32 = 1

	mockPbRestaurantDetailed := &pb.RestaurantDetailed{
		Id: mockRestaurantID,
		Name: "Sederhana",
		Latitude: 0.123,
		Longitude: 0.123,
		Menus: []*pb.Menu{
			{
				Id: 1,
				Name: "Mock menu",
				Category: "Main Course",
				Price: 10000,
			},
		},
	}
	
	
	mockRepository.Mock.On("FindRestaurantByID", mockRestaurantID).Return(mockRestaurantCompact, nil)

	var mockMenus []*pb.Menu
	mockMenus = append(mockMenus, mockMenu)
	
	mockRepository.Mock.On("FindMenuByRestaurantId", mockRestaurantID).Return(mockMenus, nil)

	mockCachingService.Mock.On("GetRestaurantDetailed", uint(mockRestaurantID)).Return(mockPbRestaurantDetailed, nil)

	pbResponse, err := merchantController.FindRestaurantById(context.Background(), &pb.IdRestaurant{Id: 1})
	assert.Nil(t, err)
	assert.NotEmpty(t, pbResponse)
}

func TestCreateMenu(t *testing.T) {
	var (
		dummyMenuID uint = 1
		dummyRestaurantID uint = 1
	)
	
	pbRequest := &pb.NewMenuData{
		AdminId: 1,
		Name: "Sate padang",
		CategoryId: 1,
		Price: 10000,
	}

	menuData := model.Menu{
		RestaurantId: dummyRestaurantID,
		Name: pbRequest.Name,
		CategoryId: uint(pbRequest.CategoryId),
		Price: pbRequest.Price,
	}

	var mockMenus []*pb.Menu
	mockMenus = append(mockMenus, mockMenu)

	mockRepository.Mock.On("FindRestaurantIdByAdminId", pbRequest.AdminId).Return(dummyRestaurantID, nil)

	mockRepository.Mock.On("CreateMenu", &menuData).Return(nil).Run(func(args mock.Arguments) {
		menu := args.Get(0).(*model.Menu)
		menu.ID = dummyMenuID
	})

	mockRepository.Mock.On("FindRestaurantByID", uint32(dummyRestaurantID)).Return(mockRestaurantCompact, nil)

	mockRepository.Mock.On("FindMenuByRestaurantId", uint32(dummyRestaurantID)).Return(mockMenus, nil)

	mockCachingService.Mock.On("SetRestaurantDetailed", dummyRestaurantID, mock.Anything).Return(nil)

	pbResponse, err := merchantController.CreateMenu(context.Background(), pbRequest)
	assert.Nil(t, err)
	assert.NotEmpty(t, pbResponse)
	assert.Equal(t, dummyMenuID, uint(pbResponse.Id))
}

func TestDeleteMenu(t *testing.T) {
	var (
		dummyMenuID uint = 1
		dummyAdminID uint32 = 1
		dummyRestaurantID uint = 1
	)

	var mockMenus []*pb.Menu
	mockMenus = append(mockMenus, mockMenu)

	mockRepository.Mock.On("FindRestaurantIdByAdminId", dummyAdminID).Return(dummyRestaurantID, nil)

	mockRepository.Mock.On("DeleteMenu", dummyRestaurantID, dummyMenuID).Return(nil)

	mockRepository.Mock.On("FindRestaurantByID", uint32(dummyRestaurantID)).Return(mockRestaurantCompact, nil)

	mockRepository.Mock.On("FindMenuByRestaurantId", uint32(dummyRestaurantID)).Return(mockMenus, nil)

	mockCachingService.Mock.On("SetRestaurantDetailed", dummyRestaurantID, mock.Anything).Return(nil)

	pbResponse, err := merchantController.DeleteMenu(context.Background(), &pb.AdminIdMenuId{AdminId: dummyAdminID, MenuId: uint32(dummyMenuID)})
	assert.Nil(t, err)
	assert.Empty(t, pbResponse)
}

func TestCreateRestaurant(t *testing.T) {
	var (
		adminID uint32 = 1
		restaurantID uint32 = 1
	)
	
	pbRequest := &pb.NewRestaurantData{
		AdminId: adminID,
		Name: "Sederhana",
		Address: "Jalan Sudirman",
		Latitude: 0.123,
		Longitude: 0.123,
	}


	mockModelRestaurant := &model.Restaurant{
		AdminId: uint(adminID),
		Name: mockRestaurantCompact.Name,
		Address: mockRestaurantCompact.Address,
		Latitude: mockRestaurantCompact.Latitude,
		Longitude: mockRestaurantCompact.Longitude,
	}

	var mockMenus []*pb.Menu
	mockMenus = append(mockMenus, mockMenu)

	mockRepository.Mock.On("CreateRestaurant", mockModelRestaurant).Return(nil).Run(func(args mock.Arguments) {
		restaurant := args.Get(0).(*model.Restaurant)
		restaurant.ID = uint(restaurantID)
	})

	mockRepository.Mock.On("FindRestaurantByID", restaurantID).Return(mockRestaurantCompact, nil)

	mockRepository.Mock.On("FindMenuByRestaurantId", restaurantID).Return(mockMenus, nil)

	mockCachingService.Mock.On("SetRestaurantDetailed", uint(restaurantID), mockMenus).Return(nil)

	pbResponse, err := merchantController.CreateRestaurant(context.Background(), pbRequest)
	assert.Nil(t, err)
	assert.Equal(t, restaurantID, pbResponse.Id)
}

func TestFindMenuById(t *testing.T) {
	pbRequest := &pb.MenuId{Id: 1}

	mockRepository.Mock.On("FindMenuById", pbRequest.Id).Return(mockMenu, nil)

	pbResponse, err := merchantController.FindMenuById(context.Background(), pbRequest)
	assert.Nil(t, err)
	assert.NotEmpty(t, pbResponse)
}

func TestUpdateMenu(t *testing.T) {
	var (
		restaurantAdminID uint32 = 1
		restaurantID uint32 = 1
	)
	
	pbRequest := &pb.UpdateMenuData{
		MenuId: 1,
		AdminId: 1,
		Name: "Mock menu",
		CategoryId: 1,
		Price: 10000,
	}

	var mockMenus []*pb.Menu
	mockMenus = append(mockMenus, mockMenu)

	mockRepository.Mock.On("FindAdminIdByMenuId", pbRequest.MenuId).Return(restaurantAdminID, nil)

	mockRepository.Mock.On("FindRestaurantIdByAdminId", pbRequest.AdminId).Return(uint(restaurantID), nil)

	mockRepository.Mock.On("UpdateMenu", pbRequest).Return(nil)

	mockRepository.Mock.On("FindRestaurantByID", restaurantID).Return(mockRestaurantCompact, nil)

	mockRepository.Mock.On("FindMenuByRestaurantId", restaurantID).Return(mockMenus, nil)

	mockCachingService.Mock.On("SetRestaurantDetailed", uint(restaurantID), mockMenus).Return(nil)

	pbResponse, err := merchantController.UpdateMenu(context.Background(), pbRequest)
	assert.Nil(t, err)
	assert.Empty(t, pbResponse)
}

func TestUpdateRestaurant(t *testing.T) {
	var (
		restaurantID uint32 = 1
		adminID uint32 = 1
	)

	pbRequest := &pb.UpdateRestaurantData{
		AdminId: adminID,
		Name: "Sederhana",
		Address: "Jalan Sudirman",
		Latitude: 0.123,
		Longitude: 0.123,
	}

	var mockMenus []*pb.Menu
	mockMenus = append(mockMenus, mockMenu)

	mockRepository.Mock.On("FindRestaurantIdByAdminId", adminID).Return(uint(restaurantID), nil)

	mockRepository.Mock.On("UpdateRestaurant", uint(restaurantID), pbRequest).Return(nil)

	mockRepository.Mock.On("FindRestaurantByID", restaurantID).Return(mockRestaurantCompact, nil)

	mockRepository.Mock.On("FindMenuByRestaurantId", restaurantID).Return(mockMenus, nil)

	mockCachingService.Mock.On("SetRestaurantDetailed", uint(restaurantID), mockMenus).Return(nil)

	pbResponse, err := merchantController.UpdateRestaurant(context.Background(), pbRequest)
	assert.Nil(t, err)
	assert.Empty(t, pbResponse)
}

func TestCalculateOrder(t *testing.T) {
	var (
		restaurantID uint32 = 1
		adminID uint32 = 1
		menuIds []int = []int{1}
	)
	
	restaurantMetadata := &pb.RestaurantMetadata{
		Id: restaurantID,
		AdminId: adminID,
		Name: "Sederhana",
		Latitude: 0.123,
		Longitude: 0.123,
	}

	pbRequest := &pb.RequestMenuDetails{
		RequestMenuDetails: []*pb.RequestMenuDetail{
			{
				Id: 1,
				Qty: 2,
			},
		},
	}

	menuDatas := []dto.MenuTmp{
		{
			ID: 1,
			Name: "Mock menu",
			Price: 10000,
		},
	}

	mockRepository.Mock.On("FindRestaurantMetadataByMenuIds", menuIds).Return(restaurantMetadata, nil)

	mockRepository.Mock.On("FindMultipleMenuDetails", menuIds).Return(menuDatas, nil)

	pbResponse, err := merchantController.CalculateOrder(context.Background(), pbRequest)
	assert.Nil(t, err)
	assert.NotEmpty(t, pbResponse)
	assert.Equal(t, float32(20000), pbResponse.ResponseMenuDetails[0].Subtotal)
}

func TestFindMenusByAdminId(t *testing.T) {
	pbRequest := &pb.AdminId{
		Id: 1,
	}

	menuCompacts := []*pb.MenuCompact{
		{
			Name: "Mock menu",
			Category: "Main course",
			Price: 10000,
		},
	}

	mockRepository.Mock.On("FindMenusByAdminId", pbRequest.Id).Return(menuCompacts, nil)

	pbResponse, err := merchantController.FindMenusByAdminId(context.Background(), pbRequest)
	assert.Nil(t, err)
	assert.NotEmpty(t, pbResponse)
	assert.Equal(t, menuCompacts, pbResponse.Menus)
}

func TestFindOneMenuByAdminId(t *testing.T) {
	pbRequest := &pb.AdminIdMenuId{
		AdminId: 1,
		MenuId: 1,
	}

	menuCompact := &pb.MenuCompact{
		Name: "Mock menu",
		Category: "Main course",
		Price: 10000,
	}

	mockRepository.Mock.On("FindOneMenuByAdminId", pbRequest.MenuId, pbRequest.AdminId).Return(menuCompact, nil)

	pbResponse, err := merchantController.FindOneMenuByAdminId(context.Background(), pbRequest)
	assert.Nil(t, err)
	assert.NotEmpty(t, pbResponse)
}

func TestFindRestaurantByAdminId(t *testing.T) {
	pbRequest := &pb.AdminId{
		Id: 1,
	}

	mockRestaurant := &pb.RestaurantData{
		Id: 1,
		AdminId: 1,
		Name: "Sederhana",
		Address: "Jalan Sudirman",
		Latitude: 0.123,
		Longitude: 0.123,
	}

	mockRepository.Mock.On("FindRestaurantByAdminId", pbRequest.Id).Return(mockRestaurant, nil)

	pbResponse, err := merchantController.FindRestaurantByAdminId(context.Background(), pbRequest)
	assert.Nil(t, err)
	assert.NotEmpty(t, pbResponse)
	assert.Equal(t, mockRestaurant, pbResponse)
}