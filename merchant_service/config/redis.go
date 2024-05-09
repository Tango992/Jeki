package config

import (
	"crypto/tls"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

func InitRedisClient() *redis.Client {
	opt, err := redis.ParseURL(os.Getenv("REDIS_URL"))
	if err != nil {
		log.Fatal(err)
	}

	client := redis.NewClient(&redis.Options{
		Addr:               opt.Addr,
		Password:           opt.Password,
		DB:                 opt.DB,
		MaxRetries:         3,
        TLSConfig: &tls.Config{
            MinVersion: tls.VersionTLS12,
        },
	})
	return client
}
