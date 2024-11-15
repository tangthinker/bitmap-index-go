package storage_driver

import (
	"context"
	"fmt"
)

type StorageDriver interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value string) error
}

const (
	StorageNil = StorageError("key not found")
)

type StorageError string

func (e StorageError) Error() string {
	return fmt.Sprintf("storage driver error: %s", string(e))
}
