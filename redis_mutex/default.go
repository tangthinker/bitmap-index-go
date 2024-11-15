package redis_mutex

import "time"

const (
	DefaultMutexTTL      = time.Second * 5
	DefaultCheckInterval = time.Millisecond * 20
)
