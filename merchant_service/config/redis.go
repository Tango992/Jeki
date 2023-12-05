package config

import (
	"os"

	"github.com/redis/go-redis/v9"
)

func InitRedisClient() *redis.Client {
    rdb := redis.NewClient(&redis.Options{
        Addr:     os.Getenv("REDIS_HOST"),
        Password: os.Getenv("REDIS_PASS"),
        DB:       0,
    })
    return rdb
}