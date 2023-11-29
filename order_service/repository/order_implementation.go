package repository

import (
	"context"
	"order-service/model"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (o OrderRepository) Create(ctx context.Context,  data *model.Order) error {
	res, err := o.Collection.InsertOne(ctx, data)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	data.Id = res.InsertedID.(primitive.ObjectID)
	return nil
}

func (o OrderRepository) FindById(ctx context.Context, orderId string) (model.Order, error) {
	objectId, err := primitive.ObjectIDFromHex(orderId)
	if err != nil {
		return model.Order{}, status.Error(codes.InvalidArgument, err.Error())
	}
	
	res := o.Collection.FindOne(ctx, bson.M{"_id": objectId})
	if err := res.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return model.Order{}, status.Error(codes.NotFound, err.Error())
		}
		return model.Order{}, status.Error(codes.Internal, err.Error())
	}
	
	var order model.Order
	if err := res.Decode(&order); err != nil {
		return model.Order{}, status.Error(codes.Internal, err.Error())
	}
	return order, nil
}

func (o OrderRepository) FindCurrentUserOrders(ctx context.Context, userId uint) ([]model.Order, error) {
	filter := bson.D{
		{Key: "$and", Value:
			bson.A{
				bson.D{{Key: "user.id", Value: userId}},
				bson.D{{Key: "driver.status", Value: "process"}},
			},
		},
	}

	orders, err := o.FindWithFilter(ctx, filter)
	if err != nil {
		return []model.Order{}, status.Error(codes.Internal, err.Error())
	}
	return orders, nil
}

func (o OrderRepository) FindWithFilter(ctx context.Context, filter bson.D) ([]model.Order, error) {
	res, err := o.Collection.Find(ctx, filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return []model.Order{}, status.Error(codes.NotFound, err.Error())
		}
		return []model.Order{}, status.Error(codes.Internal, err.Error())
	}

	orders := []model.Order{}
	if err := res.All(ctx, &orders); err != nil {
		return []model.Order{}, status.Error(codes.Internal, err.Error())
	}
	return orders, nil
}

func (o OrderRepository) Update(ctx context.Context, orderId string, updateData bson.M) error {
	return nil
}