package bitstorage

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/tangthinker/bitmap-index-go/bitmap"
	"github.com/tangthinker/bitmap-index-go/constant"
	"github.com/tangthinker/bitmap-index-go/storage_driver"
)

func (s *bitStorage) SetBits(ctx context.Context, key string, index ...int) error {

	key = bitmapKey(key)

	if err := s.redisMutex.Lock(ctx, key); err != nil {
		return fmt.Errorf("RedisBitStorage: lock key %s failed: %w", key, err)
	}
	defer func() {
		if err := s.redisMutex.Unlock(ctx, key); err != nil {
			// log error
		}
	}()

	bm := bitmap.NewBitmap()

	value, err := s.storageDriver.Get(ctx, key)
	if err != nil {
		if errors.Is(err, storage_driver.StorageNil) {
			bm.SetBits(index...)
			if err := s.storageDriver.Set(ctx, key, bm.String()); err != nil {
				return fmt.Errorf("RedisBitStorage: set key %s failed: %w", key, err)
			}
			return nil
		}
		return fmt.Errorf("RedisBitStorage: get key %s failed: %w", key, err)
	}

	bm = bitmap.ToBitmap(value)
	bm.SetBits(index...)
	bmStr := bm.String()
	if err := s.storageDriver.Set(ctx, key, bmStr); err != nil {
		return fmt.Errorf("RedisBitStorage: set key %s failed: %w", key, err)
	}
	return nil
}

func (s *bitStorage) ClearBits(ctx context.Context, key string, index ...int) error {

	key = bitmapKey(key)

	if err := s.redisMutex.Lock(ctx, key); err != nil {
		return fmt.Errorf("RedisBitStorage: lock key %s failed: %w", key, err)
	}
	defer func() {
		if err := s.redisMutex.Unlock(ctx, key); err != nil {
			// log error
		}
	}()

	value, err := s.storageDriver.Get(ctx, key)
	if err != nil {
		if errors.Is(err, storage_driver.StorageNil) {
			return nil
		}
		return fmt.Errorf("RedisBitStorage: get key %s failed: %w", key, err)
	}

	bm := bitmap.ToBitmap(value)
	bm.ClearBits(index...)
	if err := s.storageDriver.Set(ctx, key, bm.String()); err != nil {
		return fmt.Errorf("RedisBitStorage: set key %s failed: %w", key, err)
	}
	return nil
}

func (s *bitStorage) Traverse(ctx context.Context, key string, fn func(index int)) error {

	key = bitmapKey(key)

	value, err := s.storageDriver.Get(ctx, key)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil
		}
		return fmt.Errorf("RedisBitStorage: get key %s failed: %w", key, err)
	}

	bm := bitmap.ToBitmap(value)
	bm.Traverse(fn)
	return nil
}

func (s *bitStorage) Bitmap(ctx context.Context, key string) (*bitmap.Bitmap, error) {

	key = bitmapKey(key)

	value, err := s.storageDriver.Get(ctx, key)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return bitmap.NewBitmap(), nil
		}
		return nil, fmt.Errorf("RedisBitStorage: get key %s failed: %w", key, err)
	}

	return bitmap.ToBitmap(value), nil
}

func (s *bitStorage) SaveBitmap(ctx context.Context, key string, bm *bitmap.Bitmap) error {

	key = bitmapKey(key)

	if err := s.redisMutex.Lock(ctx, key); err != nil {
		return fmt.Errorf("RedisBitStorage: lock key %s failed: %w", key, err)
	}
	defer func() {
		if err := s.redisMutex.Unlock(ctx, key); err != nil {
			// log error
		}
	}()

	bmStr := bm.String()
	if err := s.storageDriver.Set(ctx, key, bmStr); err != nil {
		return fmt.Errorf("RedisBitStorage: set key %s failed: %w", key, err)
	}
	return nil
}

func bitmapKey(key string) string {
	return fmt.Sprintf("%s:%s", constant.RedisBitmapCacheKey, key)
}
