package controller

import (
	"context"
	"order-service/helpers"
	"order-service/model"
	pb "order-service/pb/orderpb"
	"order-service/repository"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
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

func (o OrderController) GetDriverAllOrders(ctx context.Context, driverId *pb.DriverId) (*pb.Orders, error) {
	ordersTmp, err := o.Repository.FindDriverAllOrders(ctx, uint(driverId.Id))
	if err != nil {
		return nil, err
	}
	orders := helpers.AssertOrdersToPb(ordersTmp)
	return &pb.Orders{Orders: orders}, nil
}

func (o OrderController) GetDriverCurrentOrder(ctx context.Context, driverId *pb.DriverId) (*pb.Order, error) {
	orderTmp, err := o.Repository.FindDriverCurrentOrder(ctx, uint(driverId.Id))
	if err != nil {
		return nil, err
	}
	order := helpers.AssertOrderToPb(orderTmp)
	return order, nil
}

func (o OrderController) GetOrderById(ctx context.Context, orderId *pb.OrderId) (*pb.Order, error) {
	orderTmp, err := o.Repository.FindById(ctx, orderId.Id)
	if err != nil {
		return nil, err
	}
	order := helpers.AssertOrderToPb(orderTmp)
	return order, nil
}

func (o OrderController) GetRestaurantAllOrders(ctx context.Context, adminId *pb.AdminId) (*pb.Orders, error) {	
	ordersTmp, err := o.Repository.FindRestaurantAllOrders(ctx, uint(adminId.Id))
	if err != nil {
		return nil, err
	}
	orders := helpers.AssertOrdersToPb(ordersTmp)
	return &pb.Orders{Orders: orders}, nil
}

func (o OrderController) GetRestaurantCurrentOrders(ctx context.Context, adminId *pb.AdminId) (*pb.Orders, error) {	
	ordersTmp, err := o.Repository.FindRestaurantCurrentOrders(ctx, uint(adminId.Id))
	if err != nil {
		return nil, err
	}
	orders := helpers.AssertOrdersToPb(ordersTmp)
	return &pb.Orders{Orders: orders}, nil
}

func (o OrderController) GetUserCurrentOrders(ctx context.Context, userData *pb.UserId) (*pb.Orders, error) {
	ordersTmp, err := o.Repository.FindUserCurrentOrders(ctx, uint(userData.Id))
	if err != nil {
		return nil, err
	}
	orders := helpers.AssertOrdersToPb(ordersTmp)
	return &pb.Orders{Orders: orders}, nil
}

func (o OrderController) GetUserAllOrders(ctx context.Context, userData *pb.UserId) (*pb.Orders, error) {
	ordersTmp, err := o.Repository.FindUserAllOrders(ctx, uint(userData.Id))
	if err != nil {
		return nil, err
	}
	orders := helpers.AssertOrdersToPb(ordersTmp)
	return &pb.Orders{Orders: orders}, nil
}

func (o OrderController) UpdateDriverOrderStatus(ctx context.Context, data *pb.RequestUpdateData) (*emptypb.Empty, error) {
	objectId, err := primitive.ObjectIDFromHex(data.OrderId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	
	if err := o.Repository.UpdateDriverStatus(ctx, objectId, data.Status); err != nil {
		return nil, err
	}
	
	return &emptypb.Empty{}, nil
}

func (o OrderController) UpdatePaymentOrderStatus(ctx context.Context, data *pb.RequestUpdateData) (*emptypb.Empty, error) {
	objectId, err := primitive.ObjectIDFromHex(data.OrderId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	
	if err := o.Repository.UpdatePaymentStatus(ctx, objectId, data.Status); err != nil {
		return nil, err
	}
	
	return &emptypb.Empty{}, nil
}

func (o OrderController) UpdateRestaurantOrderStatus(ctx context.Context, data *pb.RequestUpdateData) (*emptypb.Empty, error) {
	objectId, err := primitive.ObjectIDFromHex(data.OrderId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	
	if err := o.Repository.UpdateRestaurantStatus(ctx, objectId, data.Status); err != nil {
		return nil, err
	}
	
	return &emptypb.Empty{}, nil
}

func (o OrderController) PostOrder(ctx context.Context, data *pb.RequestOrderData) (*pb.PostOrderResponse, error){ 
	userData := model.User{
		Id: uint(data.UserId),
		Name: data.Name,
		Email: data.Email,
		Address: model.Address{
			Latitude: data.Address.Latitude,
			Longitude: data.Address.Longitude,
		},
	}

	/*
		Get data from Merchant Service
	*/
	menus := []model.Menu{}
	for _, v := range data.OrderItems {
		// Call merchant service from this block (?)
		menu := model.Menu{
			Id: uint(v.MenuId),
			Name: "Menu",							// Temporary
			Qty: uint(v.Qty),
			Subtotal: 100000,						// Temporary - Calculate subtotal from singular price
		}
		menus = append(menus, menu)
	}

	totalTemporary := float32(100000)				// Implement distance calculator
	
	orderDetailData := model.OrderDetail{
		Menus: menus,
		Total: totalTemporary,
	}

	/*
		Get data from Merchant Service
	*/
	restaurantData := model.Restaurant{
		Id: 1,										// Temporary
		AdminId: 1,									// Temporary
		Name: "Payakumbuah",						// Temporary
		Address: model.Address{
			Latitude: 0.123,  						// Temporary
			Longitude: 0.321, 						// Temporary
		},
		Status: "process",
	}
	
	/*
		Get data from User Service
	*/
	driverData := model.Driver{
		Id: 1,										// Temporary
		Name: "Foo Bar",							// Temporary
		Status: "process",
	}

	/*
		Get data from Xendit
	*/
	paymentData := model.Payment{
		InvoiceId: "someinvoiceid",					// Temporary
		InvoiceUrl: "https://www.google.com",		// Temporary
		Total: totalTemporary,						// Temporary
		Method: "ovo",								// Temporary
		Status: "pending",							// Temporary
	}

	orderData := model.Order{
		Restaurant: restaurantData,
		OrderDetail: orderDetailData,
		User: userData,
		Driver: driverData,
		Payment: paymentData,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
	defer cancel()
	
	if err := o.Repository.Create(ctx, &orderData); err != nil {
		return nil, err
	}

	response := &pb.PostOrderResponse{
		OrderId: orderData.Id.Hex(),
	}
	return response, nil 
}