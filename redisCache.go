package goRoCache

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
)

// Defining the ExpiringCache interface
type ExpiringCache interface {
	// Define the methods for this interface
}

const (
	errorTypeRedisError errorType = "RedisError"
)

type RedisCache struct {
	keysSet        map[string]struct{}
	removeChannels map[interface{}]*cacheChannel

	client *redis.Client

	mutex sync.Mutex
}

var _ (ExpiringCache) = (*RedisCache)(nil)

func NewRedisCache(address, password string, db int) *RedisCache {
	return &RedisCache{
		keysSet:        map[string]struct{}{},
		removeChannels: map[interface{}]*cacheChannel{},
		client: redis.NewClient(&redis.Options{
			Addr:     address,
			Password: password,
			DB:       db,
		}),
	}
}
func (r *RedisCache) store(key, val interface{}, ttl time.Duration) error {
	strKey := fmt.Sprintf("%v", key)
	err := r.client.Set(context.TODO(), strKey, val, ttl).Err()

	if err != nil {
		return newError(errorTypeRedisError, fmt.Sprintf("could not store key %v: %v", strKey, err))
	}

	r.keysSet[strKey] = struct{}{}

	return nil
}
