package controller

import "api-gateway/pb/orderpb"

type OrderController struct {
	Client orderpb.OrderServiceClient
}

func NewOrderController(client orderpb.OrderServiceClient) OrderController {
	return OrderController{
		Client: client,
	}
}