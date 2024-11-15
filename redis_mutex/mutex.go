package redis_mutex

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"

	"time"
)

type RedisMutex struct {
	redisClient *redis.Client
}

func NewRedisMutex(client *redis.Client) *RedisMutex {
	return &RedisMutex{
		redisClient: client,
	}
}

func (r *RedisMutex) TryLock(ctx context.Context, key string) (bool, error) {
	key = mKey(key)

	isKeyNil, err := r.isKeyNil(ctx, key)
	if err != nil {
		return false, fmt.Errorf("RedisMutex TryLock failed: %w", err)
	}
	if !isKeyNil {
		return false, nil
	}
	if err := r.redisClient.SetNX(ctx, key, "1", DefaultMutexTTL).Err(); err != nil {
		return false, fmt.Errorf("RedisMutex TryLock failed: %w", err)
	}
	return true, nil
}

func (r *RedisMutex) Lock(ctx context.Context, key string) error {
	key = mKey(key)

outer:
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			isKeyNil, err := r.isKeyNil(ctx, key)
			if err != nil {
				return fmt.Errorf("RedisMutex Lock failed: %w", err)
			}
			if !isKeyNil {
				time.Sleep(DefaultCheckInterval)
				continue
			}
			break outer
		}
	}

	if err := r.redisClient.SetNX(ctx, key, "1", DefaultMutexTTL).Err(); err != nil {
		return fmt.Errorf("RedisMutex Lock failed: %w", err)
	}
	return nil
}

func (r *RedisMutex) Unlock(ctx context.Context, key string) error {
	key = mKey(key)

	err := r.redisClient.Del(ctx, key).Err()
	if err != nil {
		return fmt.Errorf("RedisMutex Unlock failed: %w", err)
	}
	return nil
}

func (r *RedisMutex) isKeyNil(ctx context.Context, key ...string) (bool, error) {
	for _, k := range key {
		ret, err := r.redisClient.Exists(ctx, k).Result()
		if err != nil {
			return false, err
		}
		if ret == 1 {
			return false, nil
		}
	}
	return true, nil
}

func mKey(key string) string {
	return fmt.Sprintf("%s-mutex", key)
}
