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

func TestMain(m *testing.M) {
	m.Run()
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
	
	mockRestaurant := dto.RestaurantDataCompact{
		Id: mockRestaurantID,
		Name: "John Doe",
		Address: "Jalan Sudirman",
		Latitude: -6.2881637,
		Longitude: 107.04284,
	}
	
	mockRepository.Mock.On("FindRestaurantByID", mockRestaurantID).Return(mockRestaurant, nil)

	mockMenu := &pb.Menu{
		Id: 1,
		Name: "Mock menu",
		Category: "Main course",
		Price: 10000,
	}
	var mockMenus []*pb.Menu
	mockMenus = append(mockMenus, mockMenu)
	
	mockRepository.Mock.On("FindMenuByRestaurantId", mockRestaurantID).Return(mockMenus, nil)

	pbResponse, err := merchantController.FindRestaurantById(context.Background(), &pb.IdRestaurant{Id: 1})
	assert.Nil(t, err)
	assert.NotEmpty(t, pbResponse)
}

func TestCreateMenu(t *testing.T) {
	var (
		dummyMenuID uint = 1
		dummyRestaurantID uint = 2
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

	mockRepository.Mock.On("FindRestaurantIdByAdminId", pbRequest.AdminId).Return(dummyRestaurantID, nil)

	mockRepository.Mock.On("CreateMenu", &menuData).Return(nil).Run(func(args mock.Arguments) {
		menu := args.Get(0).(*model.Menu)
		menu.ID = dummyMenuID
	})
	
	pbResponse, err := merchantController.CreateMenu(context.Background(), pbRequest)
	assert.Nil(t, err)
	assert.NotEmpty(t, pbResponse)
	assert.Equal(t, dummyMenuID, uint(pbResponse.Id))
}