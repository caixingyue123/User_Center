package db

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func InitRedis() (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "123456",
		DB:       0,
	})
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}
	fmt.Println("Redis connected successfully")
	return rdb, nil
}
