package repository

import (
	"order-service/model"

	"go.mongodb.org/mongo-driver/bson"
)

type Order interface {
	Create(*model.Order) error
	FindById(string) (model.Order, error)
	FindWithFilter(bson.M) ([]model.Order, error)
	Update(string, bson.M) error
}
