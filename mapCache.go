package goRoCache

import (
	"fmt"
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

func NewMapCache() *mapCache {
	return &mapCache{
		cacheMap:       map[interface{}]interface{}{},
		removeChannels: map[interface{}]*cacheChannel{},
		updateChannels: map[interface{}]*cacheChannel{},
	}
}
func (m *mapCache) Store(key, val interface{}) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	return m.store(key, val)
}
func (m *mapCache) store(key, val interface{}) error {
	if _, exists := m.cacheMap[key]; exists {
		return newError(errorTypeAlreadyExists,
			fmt.Sprintf("key %v is already in use", key))
	}

	m.cacheMap[key] = val

	return nil
}
func (m *mapCache) Get(key interface{}) (interface{}, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	return m.get(key)
}
func (m *mapCache) get(key interface{}) (interface{}, error) {
	if _, exists := m.cacheMap[key]; !exists {
		return nil, newError(errorTypeDoesNotExist,
			fmt.Sprintf("key %v doesn't exist", key))
	}

	return m.cacheMap[key], nil
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
