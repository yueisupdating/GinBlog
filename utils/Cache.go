package utils

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,
		PoolSize: 10,
	})
	_, err := RedisClient.Ping(context.Background()).Result()

	if err != nil {
		panic(err)
	}
	fmt.Println("\n Redis connected")
}
