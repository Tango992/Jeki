package service

import (
	"context"
	"encoding/json"
	"fmt"
	"merchant-service/pb/merchantpb"
	"time"

	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CachingService interface {
	SetRestaurantDetailed(restaurantId uint, data *merchantpb.RestaurantDetailed) error
	GetRestaurantDetailed(restaurantId uint) (*merchantpb.RestaurantDetailed, error)
}

type RedisClient struct {
	Client *redis.Client
}

func NewCachingService(c *redis.Client) CachingService {
	return RedisClient{
		Client: c,
	}
}

func (r RedisClient) SetRestaurantDetailed(restaurantId uint, data *merchantpb.RestaurantDetailed) error {
	key := fmt.Sprintf("restaurantdetailed:%v", restaurantId)
	path := "$"

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	result := r.Client.JSONSet(ctx, key, path, data)
	if err := result.Err(); err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	if err := r.Client.ExpireNX(ctx, key, 2 * time.Hour).Err(); err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	return nil
}

func (r RedisClient) GetRestaurantDetailed(restaurantId uint) (*merchantpb.RestaurantDetailed, error) {
	key := fmt.Sprintf("restaurantdetailed:%v", restaurantId)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result := r.Client.JSONGet(ctx, key)
	resultByte := []byte(result.Val())
	if result.Val() == "" {
		return nil, nil
	}

	var restaurantData *merchantpb.RestaurantDetailed
	if err := json.Unmarshal(resultByte, &restaurantData); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	
	return restaurantData, nil
}
