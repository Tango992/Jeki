package controller

import (
	"context"
	"order-service/pb"
	"order-service/repository"
)

type OrderController struct {
	pb.UnimplementedOrderServiceServer
	Repository repository.Order
}

func NewOrderController(r repository.Order) OrderController {
	return OrderController{
		Repository: r,
	}
}

func (o OrderController) GetDriverOrder(ctx context.Context, driverId *pb.DriverId) (*pb.Order, error) {
	return nil, nil
}

func (o OrderController) GetOrderById(ctx context.Context, orderId *pb.OrderId) (*pb.Order, error) {
	return nil, nil
}

func (o OrderController) GetRestaurantOrders(ctx context.Context, restaurantId *pb.RestaurantId) (*pb.Orders, error) {
	return nil, nil
}

func (o OrderController) GetUserOrders(ctx context.Context, userData *pb.UserId) (*pb.Orders, error) {
	return nil, nil
}

func (o OrderController) PostOrder(ctx context.Context, data *pb.RequestOrderData) (*pb.Order, error){ 
	return nil, nil
}