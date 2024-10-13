package redis

import (
	"chain-traverser/internal/config"
	"context"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

const DB_VERSION = "1"

type RedisClient struct {
	redis          *redis.Client
	redisAnalytics *redis.Client
	redisQueue     *redis.Client
}

func NewClient(config *config.RedisConfig) *RedisClient {
	redisMain := redis.NewClient(&redis.Options{
		Addr:     config.Address,
		Password: config.Password,
		DB:       config.MAIN_DB,
	})
	redisAnalytics := redis.NewClient(&redis.Options{
		Addr:     config.Address,
		Password: config.Password,
		DB:       config.ANALYTICS_DB,
	})
	redisQueue := redis.NewClient(&redis.Options{
		Addr:     config.Address,
		Password: config.Password,
		DB:       config.QUEUE_DB,
	})
	res := RedisClient{redis: redisMain, redisAnalytics: redisAnalytics, redisQueue: redisQueue}
	return &res
}
