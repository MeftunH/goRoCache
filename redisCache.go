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
func (r *RedisCache) get(key interface{}) (interface{}, error) {
	strKey := fmt.Sprintf("%v", key)

	if _, ok := r.keysSet[strKey]; !ok {
		return nil, newError(errorTypeDoesNotExist,
			fmt.Sprintf("cannot get key %v", strKey))
	}

	val, err := r.client.Get(context.TODO(), strKey).Result()
	if err == redis.Nil {
		return nil, newError(errorTypeDoesNotExist,
			fmt.Sprintf("key %v doesn't exist", strKey))
	}

	if err != nil {
		return nil, newError(errorTypeRedisError,
			fmt.Sprintf("failed to get %v from redis", strKey))
	}

	return val, nil
}
func (r *RedisCache) remove(key interface{}) error {
	strKey := fmt.Sprintf("%v", key)

	if _, ok := r.keysSet[strKey]; !ok {
		return newError(errorTypeDoesNotExist,
			fmt.Sprintf("cannot remove key %v", strKey))
	}

	res := r.client.Del(context.TODO(), strKey).Val()
	if res < 1 {

		fmt.Println("ITAY", res)
		return newError(errorTypeDoesNotExist,
			fmt.Sprintf("could not delete key %v", strKey))
	}

	return nil
}
func (r *RedisCache) replace(key, val interface{}) error {
	return r.store(key, val, 0)
}
func (r *RedisCache) clear() error {
	for key := range r.keysSet {
		err := r.remove(key)
		if err != nil {
			return err
		}
	}

	return nil
}
func (r *RedisCache) keys() ([]interface{}, error) {
	keys := []interface{}{}

	for key := range r.keysSet {
		keys = append(keys, key)
	}

	return keys, nil
}
func (r *RedisCache) storeWithExpiration(key, val interface{}, ttl time.Duration) error {
	if ttl <= 0 {
		return newError(errorTypeNonPositivePeriod, "period must be greater than zero")
	}

	r.store(key, val, ttl)

	r.createExpirationRoutine(key, ttl)

	return nil
}

func (r *RedisCache) createExpirationRoutine(key interface{}, ttl time.Duration) {

}
