package repository

import (
	"context"
	"order-service/model"

	"go.mongodb.org/mongo-driver/bson"
)

type Order interface {
	Create(context.Context, *model.Order) error
	FindById(context.Context, string) (model.Order, error)
	FindRestaurantAllOrders(context.Context, uint) ([]model.Order, error)
	FindRestaurantCurrentOrders(context.Context, uint) ([]model.Order, error)
	FindUserAllOrders(context.Context, uint) ([]model.Order, error)
	FindUserCurrentOrders(context.Context, uint) ([]model.Order, error)
	FindDriverAllOrders(context.Context, uint) ([]model.Order, error)
	FindDriverCurrentOrder(context.Context, uint) (model.Order, error)
	FindWithFilter(context.Context, bson.D) ([]model.Order, error)
	Update(context.Context, string, bson.M) error
}
