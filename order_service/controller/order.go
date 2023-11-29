package controller

import (
	"context"
	"order-service/model"
	"order-service/pb"
	"order-service/repository"
	"time"
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

func (o OrderController) GetUserCurrentOrders(ctx context.Context, userData *pb.UserId) (*pb.Orders, error) {
	ordersTmp, err := o.Repository.FindCurrentUserOrders(ctx, uint(userData.Id))
	if err != nil {
		return nil, err
	}

	orders := []*pb.Order{}
	for _, orderTmp := range ordersTmp {
		restaurantData := &pb.Restaurant{
			Id: uint32(orderTmp.Restaurant.Id),
			AdminId: uint32(orderTmp.Restaurant.AdminId),
			Name: orderTmp.Restaurant.Name,
			Address: &pb.Address{
				Latitude: orderTmp.Restaurant.Address.Latitude,
				Longitude: orderTmp.Restaurant.Address.Longitude,
			},
			Status: orderTmp.Restaurant.Status,
		}

		var menus []*pb.Menu
		for _, menuTmp := range orderTmp.OrderDetail.Menus {
			menu := &pb.Menu{
				Id: uint32(menuTmp.Id),
				Name: menuTmp.Name,
				Quantity: uint32(menuTmp.Qty),
				Subtotal: menuTmp.Subtotal,
			}

			menus = append(menus, menu)
		}
		
		orderDetailData := &pb.OrderDetail{
			Menu: menus,
			Total: orderTmp.OrderDetail.Total,
		}

		userData := &pb.User{
			UserId: uint32(orderTmp.User.Id),
			Name: orderTmp.User.Name,
			Email: orderTmp.User.Email,
			Address: &pb.Address{
				Latitude: orderTmp.User.Address.Latitude,
				Longitude: orderTmp.User.Address.Longitude,
			},
		}

		driverData := &pb.Driver{
			Id: uint32(orderTmp.Driver.Id),
			Name: orderTmp.Driver.Name,
			Status: orderTmp.Driver.Status,
		}
		
		paymentData := &pb.Payment{
			InvoiceId: orderTmp.Payment.InvoiceId,
			InvoiceUrl: orderTmp.Payment.InvoiceUrl,
			Total: orderTmp.Payment.Total,
			Method: orderTmp.Payment.Method,
			Status: orderTmp.Payment.Status,
		}

		order := &pb.Order{
			ObjectId: orderTmp.Id.Hex(),
			Restaurant: restaurantData,
			OrderDetail: orderDetailData,
			User: userData,
			Driver: driverData,
			Payment: paymentData,
		}
		orders = append(orders, order)
	}
	
	return &pb.Orders{Orders: orders}, nil
}

func (o OrderController) GetUserAllOrders(ctx context.Context, userData *pb.UserId) (*pb.Orders, error) {
	return nil, nil
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
		menu := model.Menu{
			Id: uint(v.MenuId),
			Name: "Menu",							// Temporary
			Qty: uint(v.Qty),
			Subtotal: 100000,						// Temporary
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