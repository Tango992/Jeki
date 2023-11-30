package helpers

import (
	"order-service/model"
	pb "order-service/pb/orderpb"
)

func AssertOrdersToPb(ordersTmp []model.Order) []*pb.Order {
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
	return order
}