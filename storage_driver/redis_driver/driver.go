package redis_driver

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"github.com/tangthinker/bitmap-index-go/storage_driver"
)

type RedisDriver struct {
	redisClient *redis.Client
}

func NewRedisDriver(client *redis.Client) storage_driver.StorageDriver {
	return &RedisDriver{
		redisClient: client,
	}
}

func (r RedisDriver) Get(ctx context.Context, key string) (string, error) {
	val, err := r.redisClient.Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", storage_driver.StorageNil
		}
		return "", storage_driver.StorageError(err.Error())
	}
	return val, nil
}

func (r RedisDriver) Set(ctx context.Context, key string, value string) error {
	if err := r.redisClient.Set(ctx, key, value, 0).Err(); err != nil {
		return storage_driver.StorageError(err.Error())
	}
	return nil
}
