package goRoCache

import "time"

type Cache interface {
	Store(key, val interface{}) error
	Get(key interface{}) (interface{}, error)
	Remove(key interface{}) error
	Replace(key, val interface{}) error
	Clear() error
	Keys() ([]interface{}, error)
}

type CacheStoreWithExpiration interface {
	StoreWithExpiration(key, val interface{}, ttl time.Duration) error
}

type ExpiringCache struct {
	CacheStoreWithExpiration(key, val interface{}, ttl time.Duration) error

	ReplaceWithExpiration(key, val interface{}, ttl time.Duration) error

	Expire(key interface{}, ttl time.Duration) error
}
