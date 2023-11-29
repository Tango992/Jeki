package controller

import (
	"context"
	"order-service/model"
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

func (o OrderController) PostOrder(ctx context.Context, data *pb.RequestOrderData) (*pb.PostOrderResponse, error){ 
	userData := model.User{
		Id: uint(data.UserId),
		Name: "John Doe", 							// Temporary
		Address: model.Address{
			Latitude: data.Address.Latitude,
			Longitude: data.Address.Longitude,
		},
	}

	menus := []model.Menu{}
	for _, v := range data.OrderItems {
		menu := model.Menu{
			Id: uint(v.MenuId),
			Name: "Menu",							// Temporary
			Qty: uint(v.Qty),
			Subtotal: 100000,						// Temporary
		}
		menus = append(menus, menu)
	}

	totalTemporary := float32(100000)
	
	orderDetailData := model.OrderDetail{
		Menus: menus,
		Total: totalTemporary,
	}

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

	driverData := model.Driver{
		Id: 1,										// Temporary
		Name: "Foo Bar",							// Temporary
		Status: "process",
	}

	paymentData := model.Payment{
		InvoiceId: "someinvoiceid",					// Temporary
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

	if err := o.Repository.Create(&orderData); err != nil {
		return nil, err
	}

	response := &pb.PostOrderResponse{
		OrderId: orderData.Id.Hex(),
	}
	return response, nil 
}