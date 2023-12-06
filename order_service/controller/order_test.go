package controller_test

import (
	"context"
	"order-service/controller"
	m "order-service/mock"
	"order-service/model"
	"order-service/pb/merchantpb"
	pb "order-service/pb/orderpb"
	"order-service/pb/userpb"
	"order-service/repository"
	"order-service/service"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	orderRepository = repository.NewMockOrderRepository()
	paymentService = service.NewMockPaymentService()
	userService = m.NewMockUserService()
	merchantService = m.NewMockMerchantService()
	orderController = controller.NewOrderController(&orderRepository, &paymentService, &userService, &merchantService)
)

var (
	modelOrder = model.Order{
		Id: primitive.NewObjectID(),
		Restaurant: model.Restaurant{},
		OrderDetail: model.OrderDetail{},
		User: model.User{},
		Driver: model.Driver{},
		Payment: model.Payment{},
	}
)

func TestMain(m *testing.M) {
	m.Run()
}

func TestGetDriverAllOrders(t *testing.T) {
	var (
		driverID uint = 1
	)

	orders := []model.Order{
		modelOrder,
	}
	
	orderRepository.Mock.On("FindDriverAllOrders", context.Background(), driverID).Return(orders, nil)

	pbResponse, err := orderController.GetDriverAllOrders(context.Background(), &pb.DriverId{Id: uint32(driverID)})
	assert.Nil(t, err)
	assert.NotEmpty(t, pbResponse)
}

func TestGetDriverCurrentOrder(t *testing.T) {
	var (
		driverID uint = 1
	)
	
	orderRepository.Mock.On("FindDriverCurrentOrder", context.Background(), driverID).Return(modelOrder, nil)

	pbResponse, err := orderController.GetDriverCurrentOrder(context.Background(), &pb.DriverId{Id: uint32(driverID)})
	assert.Nil(t, err)
	assert.NotEmpty(t, pbResponse)
}

func TestGetOrderById(t *testing.T) {
	pbRequest := &pb.OrderId{
		Id: "abcd",
	}

	orderRepository.Mock.On("FindById", context.Background(), "abcd").Return(modelOrder, nil)

	pbResponse, err := orderController.GetOrderById(context.Background(), pbRequest)
	assert.Nil(t, err)
	assert.NotEmpty(t, pbResponse)
}

func TestGetRestaurantAllOrders(t *testing.T) {
	pbRequest := &pb.AdminId{
		Id: 1,
	}

	orders := []model.Order{modelOrder}

	orderRepository.Mock.On("FindRestaurantAllOrders", context.Background(), uint(pbRequest.Id)).Return(orders, nil)

	pbResponse, err := orderController.GetRestaurantAllOrders(context.Background(), pbRequest)
	assert.Nil(t, err)
	assert.NotEmpty(t, pbResponse)
}

func TestGetRestaurantCurrentOrder(t *testing.T) {
	pbRequest := &pb.AdminId{
		Id: 1,
	}

	orders := []model.Order{modelOrder}

	orderRepository.Mock.On("FindRestaurantCurrentOrders", context.Background(), uint(pbRequest.Id)).Return(orders, nil)

	pbResponse, err := orderController.GetRestaurantCurrentOrders(context.Background(), pbRequest)
	assert.Nil(t, err)
	assert.NotEmpty(t, pbResponse)
}

func TestGetUserCurrentOrders(t *testing.T) {
	pbRequest := &pb.UserId{
		Id: 1,
	}

	orders := []model.Order{modelOrder}

	orderRepository.Mock.On("FindUserCurrentOrders", context.Background(), uint(pbRequest.Id)).Return(orders, nil)

	pbResponse, err := orderController.GetUserCurrentOrders(context.Background(), pbRequest)
	assert.Nil(t, err)
	assert.NotEmpty(t, pbResponse)
}

func TestGetUserAllOrders(t *testing.T) {
	pbRequest := &pb.UserId{
		Id: 1,
	}

	orders := []model.Order{modelOrder}

	orderRepository.Mock.On("FindUserAllOrders", context.Background(), uint(pbRequest.Id)).Return(orders, nil)

	pbResponse, err := orderController.GetUserAllOrders(context.Background(), pbRequest)
	assert.Nil(t, err)
	assert.NotEmpty(t, pbResponse)
}

func TestUpdateDriverOrderStatus(t *testing.T) {
	objectID := primitive.NewObjectID()

	pbRequest := &pb.RequestUpdateData{
		UserId: 1,
		OrderId: objectID.Hex(),
		Status: "cancelled",
	}
	
	userService.Mock.On("SetDriverStatusOnline", mock.Anything, &userpb.DriverId{Id: 1}).Return(&emptypb.Empty{}, nil)
	orderRepository.Mock.On("CancelOrderStatus", context.Background(), objectID).Return(nil)

	pbResponse, err := orderController.UpdateDriverOrderStatus(context.Background(), pbRequest)
	assert.Nil(t, err)
	assert.Empty(t, pbResponse)
}

func TestUpdatePaymentOrderStatus(t *testing.T) {
	objectID := primitive.NewObjectID()

	pbRequest := &pb.RequestUpdatePayment{
		OrderId: objectID.Hex(),
		Status: "PAID",
		InvoiceId: "abcd",
		Method: "CREDIT_CARD",
		CompletedAt: "2023:12:06 23:53",
	}

	orderRepository.Mock.On("UpdatePaymentStatus", context.Background(), objectID, pbRequest.Status, pbRequest.Method, pbRequest.CompletedAt).Return(nil)

	pbResponse, err := orderController.UpdatePaymentOrderStatus(context.Background(), pbRequest)
	assert.Nil(t, err)
	assert.Empty(t, pbResponse)
}

func TestUpdateRestaurantStatus(t *testing.T) {
	objectID := primitive.NewObjectID()

	pbRequest := &pb.RequestUpdateData{
		UserId: 1,
		OrderId: objectID.Hex(),
		Status: "done",
	}

	orderRepository.Mock.On("UpdateRestaurantStatus", context.Background(), objectID, pbRequest.UserId, pbRequest.Status).Return(nil)

	pbResponse, err := orderController.UpdateRestaurantOrderStatus(context.Background(), pbRequest)
	assert.Nil(t, err)
	assert.Empty(t, pbResponse)
}

func TestPostOrder(t *testing.T) {
	pbRequest := &pb.RequestOrderData{
		UserId: 1,
		Name: "John Doe",
		Email: "john@example.com",
		Address: &pb.Address{
			Latitude: 0.123,
			Longitude: 4.321,
		},
		OrderItems: []*pb.OrderItem{
			{
				MenuId: 1,
				Qty: 1,
			},
		},
	}

	requestCalculateOrder := &merchantpb.RequestMenuDetails{
		RequestMenuDetails: []*merchantpb.RequestMenuDetail{
			{Id: 1, Qty: 1},
		},
	}

	responseCalculateOrder := &merchantpb.CalculateOrderResponse{
		RestaurantData: &merchantpb.RestaurantMetadata{},
		ResponseMenuDetails: []*merchantpb.ResponseMenuDetail{
			{Id: 1, Name: "Mock menu", Qty: 1, Price: 10, Subtotal: 10},
		},
	}

	availableDriver := &userpb.DriverData{
		Id: 1,
		Name: "Foo Bar",
	}

	paymentData := model.Payment{
		InvoiceId: "abcd",
		InvoiceUrl: "efgh",
		Total: 10000,
		Method: "CREDIT_CARD",
		Status: "PAID",
	}
	
	merchantService.Mock.On("CalculateOrder", mock.Anything, requestCalculateOrder).Return(responseCalculateOrder, nil)

	userService.Mock.On("GetAvailableDriver", mock.Anything, &emptypb.Empty{}).Return(availableDriver, nil)

	paymentService.Mock.On("MakeInvoice", mock.Anything, mock.Anything).Return(paymentData, nil)

	orderRepository.Mock.On("Create", mock.Anything, mock.Anything).Return(nil)

	orderRepository.Mock.On("FindById", mock.Anything, mock.Anything).Return(modelOrder, nil)

	pbResponse, err := orderController.PostOrder(context.Background(), pbRequest)
	assert.Nil(t, err)
	assert.NotEmpty(t, pbResponse)
}