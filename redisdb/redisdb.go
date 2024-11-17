package redisdb

import (
	"context"
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
)

var (
	ctx          = context.Background()
	RedisClient  *redis.Client
	redisAddress string
)

func InitRedis() *redis.Client {
	const redisEnv = "REDIS_ADDRESS"
	redisAddress = os.Getenv(redisEnv)
	if redisAddress == "" {
		_, _ = fmt.Fprintln(os.Stderr, redisEnv)
		redisAddress = "localhost" // defualt
	}

	client := redis.NewClient(&redis.Options{
		Addr:     redisAddress,
		Password: "",
		DB:       0,
	})

	return client
}
