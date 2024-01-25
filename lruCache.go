package goRoCache

import (
	"container/list"
	"sync"
)

type lruItem struct {
	value interface{}
	node  *list.Element
}

type lruCache struct {
	capacity      int
	numberOfItems int
	storage       Cache
	list          *list.List

	mutex sync.Mutex
}

var _ Cache = (*lruCache)(nil)

func NewLru(capacity int) *lruCache {
	return &lruCache{
		capacity: capacity,
		storage:  NewMapCache(),
		list:     list.New(),
	}
}

func NewLruWithCustomCache(capacity int, cache Cache) (*lruCache, error) {
	keys, err := cache.Keys()
	if err != nil {
		return nil, err
	}

	if len(keys) > 0 {
		return nil, newError(errorTypeCacheNotEmpty, "supplied cache must be empty")
	}

	return &lruCache{
		capacity: capacity,
		storage:  cache,
		list:     list.New(),
	}, nil
}

func (lru *lruCache) Store(key, val interface{}) error {
	lru.mutex.Lock()
	defer lru.mutex.Unlock()

	return lru.store(key, val)
}
