package bitstorage

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/tangthinker/bitmap-index-go/bitmap"
	"github.com/tangthinker/bitmap-index-go/redis_mutex"
	"github.com/tangthinker/bitmap-index-go/storage_driver"
	"github.com/tangthinker/bitmap-index-go/storage_driver/redis_driver"
)

type BitmapStorage interface {
	SetBits(ctx context.Context, key string, index ...int) error
	ClearBits(ctx context.Context, key string, index ...int) error
	Traverse(ctx context.Context, key string, fn func(index int)) error

	Bitmap(ctx context.Context, key string) (*bitmap.Bitmap, error)
	SaveBitmap(ctx context.Context, key string, bm *bitmap.Bitmap) error
}

type bitStorage struct {
	storageDriver storage_driver.StorageDriver
	redisMutex    *redis_mutex.RedisMutex
}

func NewRedisBitStorage(client *redis.Client) BitmapStorage {
	return &bitStorage{
		storageDriver: redis_driver.NewRedisDriver(client),
		redisMutex:    redis_mutex.NewRedisMutex(client),
	}
}
