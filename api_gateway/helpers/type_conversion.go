package helpers

import (
	"api-gateway/dto"
	"api-gateway/pb/orderpb"
)

func AssertToPbOrderItems(orders dto.NewOrderRequest) []*orderpb.OrderItem {
	var pbOrderItems []*orderpb.OrderItem
	for _, order := range orders.Items{
		pbOrderItem := &orderpb.OrderItem{
			MenuId: order.MenuID,
			Qty: order.Qty,
		}
		pbOrderItems = append(pbOrderItems, pbOrderItem)
	}
	return pbOrderItems
}