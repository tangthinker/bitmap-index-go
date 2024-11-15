package redis_mutex

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

const (
	WLockKeyPrefix string = "w-lock:"
	RLockKeyPrefix string = "r-lock:"
)

type RedisRWMutex struct {
	redisClient *redis.Client
}

func NewRedisRWMutex(client *redis.Client) *RedisRWMutex {
	return &RedisRWMutex{
		redisClient: client,
	}
}

func (r *RedisRWMutex) Lock(ctx context.Context, key string) error {

outer:
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			isKeyNil, err := r.isKeyNil(ctx, rLockKey(key), wLockKey(key))
			if err != nil {
				return fmt.Errorf("RedisRWMutex Lock failed: %w", err)
			}
			if !isKeyNil {
				time.Sleep(DefaultCheckInterval)
				continue
			}
			break outer
		}
	}

	if err := r.redisClient.SetNX(ctx, wLockKey(key), "1", DefaultMutexTTL).Err(); err != nil {
		return fmt.Errorf("RedisRWMutex Lock failed: %w", err)
	}

	return nil
}

func (r *RedisRWMutex) Unlock(ctx context.Context, key string) error {
	return r.redisClient.Del(ctx, wLockKey(key)).Err()
}

func (r *RedisRWMutex) RLock(ctx context.Context, key string) error {
outer:
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			isKeyNil, err := r.isKeyNil(ctx, wLockKey(key))
			if err != nil {
				return fmt.Errorf("RedisRWMutex RLock failed: %w", err)
			}
			if !isKeyNil {
				time.Sleep(DefaultCheckInterval)
				continue
			}
			break outer
		}
	}

	if err := r.redisClient.Incr(ctx, rLockKey(key)).Err(); err != nil {
		return fmt.Errorf("RedisRWMutex RLock failed: %w", err)
	}
	return nil
}

func (r *RedisRWMutex) RUnlock(ctx context.Context, key string) error {
	return r.redisClient.Del(ctx, key).Err()
}

func (r *RedisRWMutex) isKeyNil(ctx context.Context, key ...string) (bool, error) {
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

func rLockKey(key string) string {
	return RLockKeyPrefix + key
}

func wLockKey(key string) string {
	return WLockKeyPrefix + key
}
