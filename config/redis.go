package config

import (
	"context"

	"github.com/go-redis/redis/v8"
)

func initRedis(database Redis) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     database.Addr,
		Password: database.Password,
		DB:       database.DB,
	})

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}

	return rdb, nil
}
