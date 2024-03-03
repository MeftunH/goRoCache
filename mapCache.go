package goRoCache

import (
	_ "fmt"
	"sync"
	"time"
	_ "time"
)

type cacheChannel struct {
	stopChannel chan bool
}
type mapCache struct {
	cacheMap map[interface{}]interface{}

	removeChannels map[interface{}]*cacheChannel

	updateChannels map[interface{}]*cacheChannel

	mutex sync.Mutex
}

func (m mapCache) Store(key, val interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (m mapCache) Get(key interface{}) (interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (m mapCache) Remove(key interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (m mapCache) Replace(key, val interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (m mapCache) Clear() error {
	//TODO implement me
	panic("implement me")
}

func (m mapCache) Keys() ([]interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (m mapCache) StoreWithExpiration(key interface{}, item lfuItem, ttl time.Duration) interface{} {
	//TODO implement me
	panic("implement me")
}

func (m mapCache) StoreWithUpdate(key, initialValue interface{}, updateFunc func(currValue interface{}) interface{}, period time.Duration) error {
	//TODO implement me
	panic("implement me")
}

func (m mapCache) ReplaceWithUpdate(key, initialValue interface{}, updateFunc func(currValue interface{}) interface{}, period time.Duration) error {
	//TODO implement me
	panic("implement me")
}

type UpdatingExpiringCache interface {
	UpdatingCache
}

var _ UpdatingExpiringCache = (*mapCache)(nil)

func NewMapCache() *mapCache {
	return &mapCache{
		cacheMap:       map[interface{}]interface{}{},
		removeChannels: map[interface{}]*cacheChannel{},
		updateChannels: map[interface{}]*cacheChannel{},
	}
}
