package config

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

var rdClient *redis.Client
var duration = time.Hour

type RedisClient struct{}

func InitRedis() (*RedisClient, error) {
	rdClient = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.url"),
		Password: "",
		DB:       0,
	})
	if _, err := rdClient.Ping(context.Background()).Result(); err != nil {
		return nil, err
	}
	return &RedisClient{}, nil
}

func (rd *RedisClient) Set(key string, val any, rest ...any) error {
	return rdClient.Set(context.Background(), key, val, duration).Err()
}

func (rd *RedisClient) Get(key string) (any, error) {
	return rdClient.Get(context.Background(), key).Result()
}

func (rd *RedisClient) Del(key ...string) error {
	return rdClient.Del(context.Background(), key...).Err()
}
