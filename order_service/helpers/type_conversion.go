package helpers

import (
	"order-service/model"
	pb "order-service/pb/orderpb"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AssertOrdersToPb(ordersTmp []model.Order) []*pb.Order {
	orders := []*pb.Order{}
	for _, orderTmp := range ordersTmp {
		order := AssertOrderToPb(orderTmp)
		orders = append(orders, order)
	}
	return orders
}

func AssertOrderToPb(orderTmp model.Order) *pb.Order {
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
			Price: menuTmp.Price,
			Subtotal: menuTmp.Subtotal,
		}

		menus = append(menus, menu)
	}
	
	orderDetailData := &pb.OrderDetail{
		Menu: menus,
		ItemsSubtotal: orderTmp.OrderDetail.ItemsSubtotal,
		DeliveryFee: orderTmp.OrderDetail.DeliveryFee,
		GrandTotal: orderTmp.OrderDetail.GrandTotal,
		Status: orderTmp.OrderDetail.Status,
		CreatedAt: orderTmp.OrderDetail.CreatedAt.Time().String(),
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
	return order
}

func AssertOrderResponseToPb(objectId primitive.ObjectID, restaurant model.Restaurant, orderDetail model.OrderDetail, user model.User, driver model.Driver, payment model.Payment) *pb.Order {
	restaurantPb := &pb.Restaurant{
		Id: uint32(restaurant.Id),
		AdminId: uint32(restaurant.AdminId),
		Name: restaurant.Name,
		Address: &pb.Address{
			Latitude: restaurant.Address.Latitude,
			Longitude: restaurant.Address.Longitude,
		},
		Status: restaurant.Status,
	}

	var menus []*pb.Menu
	for _, menuTmp := range orderDetail.Menus {
		menu := &pb.Menu{
			Id: uint32(menuTmp.Id),
			Name: menuTmp.Name,
			Quantity: uint32(menuTmp.Qty),
			Price: menuTmp.Price,
			Subtotal: menuTmp.Subtotal,
		}
		menus = append(menus, menu)
	}

	orderDetailPb := &pb.OrderDetail{
		Menu: menus,
		ItemsSubtotal: orderDetail.ItemsSubtotal,
		DeliveryFee: orderDetail.DeliveryFee,
		GrandTotal: orderDetail.GrandTotal,
		Status: orderDetail.Status,
		CreatedAt: orderDetail.CreatedAt.Time().Format("2006-01-02 15:04:05"),
	}

	userPb := &pb.User{
		UserId: uint32(user.Id),
		Name: user.Name,
		Email: user.Email,
		Address: &pb.Address{
			Latitude: user.Address.Latitude,
			Longitude: user.Address.Longitude,
		},
	}

	driverPb := &pb.Driver{
		Id: uint32(driver.Id),
		Name: driver.Name,
		Status: driver.Status,
	}
	
	paymentPb := &pb.Payment{
		InvoiceId: payment.InvoiceId,
		InvoiceUrl: payment.InvoiceUrl,
		Total: payment.Total,
		Method: payment.Method,
		Status: payment.Status,
	}
	
	return &pb.Order{
		ObjectId: objectId.Hex(),
		Restaurant: restaurantPb,
		OrderDetail: orderDetailPb,
		User: userPb,
		Driver: driverPb,
		Payment: paymentPb,
	}
}