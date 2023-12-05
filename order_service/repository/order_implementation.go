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

const (
	orderStatusProcess = "process"
	orderStatusDone = "done"
	orderStatusCancel = "cancelled"
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

func (o OrderRepository) FindDriverCurrentOrder(ctx context.Context, driverId uint) (model.Order, error) {
	filter := bson.D{
		{Key: "$and", Value:
			bson.A{
				bson.D{{Key: "driver.id", Value: driverId}},
				bson.D{{Key: "driver.status", Value: orderStatusProcess}},
				bson.D{{Key: "order_detail.status", Value: orderStatusProcess}},
				bson.D{{Key: "payment.status", Value: "PAID"}},
			},
		},
	}

	res := o.Collection.FindOne(ctx, filter)
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

func (o OrderRepository) FindDriverAllOrders(ctx context.Context, driverId uint) ([]model.Order, error) {
	filter := bson.D{{
		Key: "driver.id", 
		Value: driverId,
	}}

	orders, err := o.FindWithFilter(ctx, filter)
	if err != nil {
		return []model.Order{}, status.Error(codes.Internal, err.Error())
	}
	return orders, nil
}

func (o OrderRepository) FindRestaurantAllOrders(ctx context.Context, adminId uint) ([]model.Order, error) {
	filter := bson.D{{
		Key: "restaurant.admin_id", 
		Value: adminId,
	}}

	orders, err := o.FindWithFilter(ctx, filter)
	if err != nil {
		return []model.Order{}, status.Error(codes.Internal, err.Error())
	}
	return orders, nil
}

func (o OrderRepository) FindUserAllOrders(ctx context.Context, userId uint) ([]model.Order, error) {
	filter := bson.D{{
		Key: "user.id", 
		Value: userId,
	}}

	orders, err := o.FindWithFilter(ctx, filter)
	if err != nil {
		return []model.Order{}, status.Error(codes.Internal, err.Error())
	}
	return orders, nil
}

func (o OrderRepository) FindRestaurantCurrentOrders(ctx context.Context, adminId uint) ([]model.Order, error) {
	filter := bson.D{
		{Key: "$and", Value:
			bson.A{
				bson.D{{Key: "restaurant.admin_id", Value: adminId}},
				bson.D{{Key: "restaurant.status", Value: orderStatusProcess}},
				bson.D{{Key: "order_detail.status", Value: orderStatusProcess}},
				bson.D{{Key: "payment.status", Value: "PAID"}},
			},
		},
	}

	orders, err := o.FindWithFilter(ctx, filter)
	if err != nil {
		return []model.Order{}, status.Error(codes.Internal, err.Error())
	}
	return orders, nil
}

func (o OrderRepository) FindUserCurrentOrders(ctx context.Context, userId uint) ([]model.Order, error) {
	filter := bson.D{
		{Key: "$and", Value:
			bson.A{
				bson.D{{Key: "user.id", Value: userId}},
				bson.D{{Key: "driver.status", Value: orderStatusProcess}},
				bson.D{{Key: "order_detail.status", Value: orderStatusProcess}},
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

func (o OrderRepository) UpdateRestaurantStatus(ctx context.Context, orderId primitive.ObjectID, userId uint32, status string) error {
	filter := bson.D{
		{Key: "$and", Value:
			bson.A{
				bson.D{{Key: "_id", Value: orderId}},
				bson.D{{Key: "restaurant.admin_id", Value: userId}},
				bson.D{{Key: "order_detail.status", Value: orderStatusProcess}},
			},
		},
	}
	updateData := bson.M{"$set": bson.M{"restaurant.status": status}}

	if err := o.UpdateWithFilter(ctx, filter, updateData); err != nil {
		return err
	}
	return nil
}

func (o OrderRepository) UpdateDriverStatus(ctx context.Context, orderId primitive.ObjectID, userId uint32, status string) error {
	filter := bson.D{
		{Key: "$and", Value:
			bson.A{
				bson.D{{Key: "_id", Value: orderId}},
				bson.D{{Key: "driver.id", Value: userId}},
				bson.D{{Key: "order_detail.status", Value: orderStatusProcess}},
			},
		},
	}
	updateData := bson.M{"$set": bson.M{"driver.status": status}}
	if err := o.UpdateWithFilter(ctx, filter, updateData); err != nil {
		return err
	}
	return nil
}

func (o OrderRepository) UpdatePaymentStatus(ctx context.Context, orderId primitive.ObjectID, status string) error {
	filter := bson.D{{Key: "_id", Value: orderId}}
	updateData := bson.M{"$set": bson.M{"payment.status": status}}

	if err := o.UpdateWithFilter(ctx, filter, updateData); err != nil {
		return err
	}
	return nil
}

func (o OrderRepository) CancelOrderStatus(ctx context.Context, orderId primitive.ObjectID) error {
	filter := bson.D{{Key: "_id", Value: orderId}}
	updateData := bson.M{"$set": bson.M{
		"order_detail.status": orderStatusCancel,
		"restaurant.status": orderStatusCancel,
		"driver.status": orderStatusCancel,
	}}

	if err := o.UpdateWithFilter(ctx, filter, updateData); err != nil {
		return err
	}
	return nil
}

func (o OrderRepository) CompleteOrderStatus(ctx context.Context, orderId primitive.ObjectID) error {
	filter := bson.D{{Key: "_id", Value: orderId}}
	updateData := bson.M{"$set": bson.M{
		"order_detail.status": orderStatusDone,
		"restaurant.status": orderStatusDone,
		"driver.status": orderStatusDone,
	}}

	if err := o.UpdateWithFilter(ctx, filter, updateData); err != nil {
		return err
	}
	return nil
}

func (o OrderRepository) UpdateWithFilter(ctx context.Context, field bson.D, data bson.M) error {
	res := o.Collection.FindOneAndUpdate(ctx, field, data)
	if err := res.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return status.Error(codes.NotFound, err.Error())
		}
		return status.Error(codes.Internal, err.Error())
	}
	return nil
}