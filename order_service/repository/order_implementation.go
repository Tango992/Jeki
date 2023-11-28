package repository

import (
	"order-service/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderRepository struct {
	Collection *mongo.Collection
}

func NewOrderRepository(collection *mongo.Collection) OrderRepository {
	return OrderRepository{
		Collection: collection,
	}
}

func (o OrderRepository) Create(data *model.Order) error {
	return nil
}

func (o OrderRepository) FindById(orderId string) (model.Order, error) {
	return model.Order{}, nil
}

func (o OrderRepository) FindWithFilter(filter bson.M) ([]model.Order, error) {
	return nil, nil
}

func (o OrderRepository) Update(orderId string, updateData bson.M) error {
	return nil
}