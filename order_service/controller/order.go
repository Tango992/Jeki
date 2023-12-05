package controller

import (
	"context"
	"order-service/helpers"
	"order-service/model"
	"order-service/pb/merchantpb"
	pb "order-service/pb/orderpb"
	"order-service/pb/userpb"
	"order-service/repository"
	"order-service/service"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	grpcMetadata "google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

const (
	statusOnProcess = "process"
	statusDone      = "done"
	statusCancel    = "cancelled"
)

type OrderController struct {
	pb.UnimplementedOrderServiceServer
	Repository      repository.Order
	PaymentService  service.PaymentService
	UserService     userpb.UserClient
	MerchantService merchantpb.MerchantClient
}

func NewOrderController(r repository.Order, p service.PaymentService, us userpb.UserClient, ms merchantpb.MerchantClient) OrderController {
	return OrderController{
		Repository:      r,
		PaymentService:  p,
		UserService:     us,
		MerchantService: ms,
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

	token, err := helpers.SignJwtForGrpc()
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	ctxWithAuth := grpcMetadata.AppendToOutgoingContext(ctx, "authorization", "Bearer "+token)
	if _, err := o.UserService.SetDriverStatusOnline(ctxWithAuth, &userpb.DriverId{Id: data.UserId}); err != nil {
		return nil, err
	}

	switch data.Status {
	case statusCancel:
		if err := o.Repository.CancelOrderStatus(ctx, objectId); err != nil {
			return nil, err
		}

	case statusDone:
		if err := o.Repository.CompleteOrderStatus(ctx, objectId); err != nil {
			return nil, err
		}
	}

	if err := o.Repository.UpdateDriverStatus(ctx, objectId, data.UserId, data.Status); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (o OrderController) UpdatePaymentOrderStatus(ctx context.Context, data *pb.RequestUpdatePayment) (*emptypb.Empty, error) {
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
	token, err := helpers.SignJwtForGrpc()
	if err != nil {
		return nil, err
	}

	objectId, err := primitive.ObjectIDFromHex(data.OrderId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if err := o.Repository.UpdateRestaurantStatus(ctx, objectId, data.UserId, data.Status); err != nil {
		return nil, err
	}

	switch data.Status {
	case statusCancel:
		if err := o.Repository.CancelOrderStatus(ctx, objectId); err != nil {
			return nil, err
		}

		orderData, _ := o.GetOrderById(ctx, &pb.OrderId{Id: data.OrderId})

		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()

		ctxWithAuth := grpcMetadata.AppendToOutgoingContext(ctx, "authorization", "Bearer "+token)
		if _, err := o.UserService.SetDriverStatusOnline(ctxWithAuth, &userpb.DriverId{Id: orderData.Driver.Id}); err != nil {
			return nil, err
		}

	case statusDone:
		if err := o.Repository.CompleteOrderStatus(ctx, objectId); err != nil {
			return nil, err
		}
	}

	return &emptypb.Empty{}, nil
}

func (o OrderController) PostOrder(ctx context.Context, data *pb.RequestOrderData) (*pb.Order, error) {
	token, err := helpers.SignJwtForGrpc()
	if err != nil {
		return nil, err
	}

	newObjectId := primitive.NewObjectID()

	userData := model.User{
		Id:    int(data.UserId),
		Name:  data.Name,
		Email: data.Email,
		Address: model.Address{
			Latitude:  data.Address.Latitude,
			Longitude: data.Address.Longitude,
		},
	}

	var requestMenus []*merchantpb.RequestMenuDetail
	for _, menu := range data.OrderItems {
		menuTmp := &merchantpb.RequestMenuDetail{
			Id:  menu.MenuId,
			Qty: menu.Qty,
		}
		requestMenus = append(requestMenus, menuTmp)
	}

	requestMenuDetails := &merchantpb.RequestMenuDetails{
		RequestMenuDetails: requestMenus,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	ctxWithAuth := grpcMetadata.AppendToOutgoingContext(ctx, "authorization", "Bearer "+token)
	merchantData, err := o.MerchantService.CalculateOrder(ctxWithAuth, requestMenuDetails)
	if err != nil {
		return nil, err
	}

	var menus []model.Menu
	var itemsSubtotal float32
	for _, v := range merchantData.ResponseMenuDetails {
		itemsSubtotal += v.Subtotal
		menu := model.Menu{
			Id:       int(v.Id),
			Name:     v.Name,
			Qty:      int(v.Qty),
			Price:    v.Price,
			Subtotal: v.Subtotal,
		}
		menus = append(menus, menu)
	}

	restaurantData := model.Restaurant{
		Id:      int(merchantData.RestaurantData.Id),
		AdminId: int(merchantData.RestaurantData.AdminId),
		Name:    merchantData.RestaurantData.Name,
		Address: model.Address{
			Latitude:  merchantData.RestaurantData.Latitude,
			Longitude: merchantData.RestaurantData.Longitude,
		},
		Status: statusOnProcess,
	}

	deliveryFee := helpers.CalculateDeliveryFee(restaurantData, userData)
	grandTotal := deliveryFee + itemsSubtotal

	orderDetailData := model.OrderDetail{
		Menus:         menus,
		DeliveryFee:   deliveryFee,
		ItemsSubtotal: itemsSubtotal,
		GrandTotal:    grandTotal,
		Status:        statusOnProcess,
		CreatedAt:     primitive.NewDateTimeFromTime(time.Now()),
	}

	ctx, cancel = context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	ctxWithAuth = grpcMetadata.AppendToOutgoingContext(ctx, "authorization", "Bearer "+token)
	availableDriver, err := o.UserService.GetAvailableDriver(ctxWithAuth, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}

	driverData := model.Driver{
		Id:     int(availableDriver.Id),
		Name:   availableDriver.Name,
		Status: statusOnProcess,
	}

	paymentData, err := o.PaymentService.MakeInvoice(newObjectId, grandTotal)
	if err != nil {
		return nil, err
	}

	orderData := model.Order{
		Id:          newObjectId,
		Restaurant:  restaurantData,
		OrderDetail: orderDetailData,
		User:        userData,
		Driver:      driverData,
		Payment:     paymentData,
	}

	ctx, cancel = context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	if err := o.Repository.Create(ctx, &orderData); err != nil {
		return nil, err
	}

	ctx, cancel = context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	responseOrderData, err := o.GetOrderById(ctx, &pb.OrderId{Id: newObjectId.Hex()})
	return responseOrderData, nil
}
