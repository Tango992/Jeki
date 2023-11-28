package controller

import "order-service/pb"

type OrderController struct {
	pb.UnimplementedOrderServiceServer
}

func NewOrderController() OrderController {
	return OrderController{
		
	}
}