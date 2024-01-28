package goRoCache

import (
	"container/heap"
	"sync"
)

type lfuHeapItem struct {
	value     interface{}
	frequency int
	index     int
}
type lfuHeap []*lfuHeapItem

func (h lfuHeap) Less(i, j int) bool {
	return h[i].frequency < h[j].frequency
}

func (h lfuHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].index = j
	h[j].index = i
}

func (h *lfuHeap) Push(x interface{}) {
	n := len(*h)
	item := x.(*lfuHeapItem)
	item.index = n
	*h = append(*h, item)
}

func (h *lfuHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*h = old[0 : n-1]
	return item
}

var _ heap.Interface = (*lfuHeap)(nil)

func (h lfuHeap) Len() int {
	return len(h)
}

type lfuItem struct {
	heapItem *lfuHeapItem
	value    interface{}
}

type lfuCache struct {
	capacity int

	storage Cache

	heap lfuHeap

	mutex sync.Mutex
}

func (lfu *lfuCache) Replace(key, val interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (lfu *lfuCache) Clear() error {
	//TODO implement me
	panic("implement me")
}

func (lfu *lfuCache) Keys() ([]interface{}, error) {
	//TODO implement me
	panic("implement me")
}

var _ Cache = (*lfuCache)(nil)

func NewLfu(capacity int) *lfuCache {
	return &lfuCache{
		capacity: capacity,
		storage:  NewMapCache(),
		heap:     lfuHeap{},
	}
}
func NewLfuWithCustomCache(capacity int, cache Cache) (*lfuCache, error) {
	keys, err := cache.Keys()
	if err != nil {
		return nil, err
	}

	if len(keys) > 0 {
		return nil, newError(errorTypeCacheNotEmpty, "supplied cache must be empty")
	}

	return &lfuCache{
		capacity: capacity,
		storage:  cache,
		heap:     lfuHeap{},
	}, nil
}
func (lfu *lfuCache) Store(key, val interface{}) error {
	lfu.mutex.Lock()
	defer lfu.mutex.Unlock()

	return lfu.store(key, val)
}

func (lfu *lfuCache) store(key, val interface{}) error {
	heapItem := &lfuHeapItem{
		value:     key,
		frequency: 0,
	}

	item := lfuItem{heapItem, val}

	err := lfu.storage.Store(key, item)
	if err != nil {
		return err
	}

	heap.Push(&lfu.heap, heapItem)

	if lfu.heap.Len() > lfu.capacity {
		heapItem := heap.Pop(&lfu.heap).(*lfuHeapItem)
		err := lfu.storage.Remove(heapItem.value)
		if err != nil {
			return err
		}
	}

	return nil
}
func (lfu *lfuCache) Get(key interface{}) (interface{}, error) {
	lfu.mutex.Lock()
	defer lfu.mutex.Unlock()

	return lfu.get(key)
}

func (lfu *lfuCache) get(key interface{}) (interface{}, error) {
	item, err := lfu.storage.Get(key)
	if err != nil {
		return nil, err
	}

	lfuItem := item.(lfuItem)
	lfuItem.heapItem.frequency++

	heap.Init(&lfu.heap)

	return lfuItem.value, nil
}
func (lfu *lfuCache) GetLeastFrequentlyUsedKey() interface{} {
	if lfu.isEmpty() {
		return nil
	}
	return lfu.heap[0].value
}
func (lfu *lfuCache) Remove(key interface{}) error {
	lfu.mutex.Lock()
	defer lfu.mutex.Unlock()

	return lfu.remove(key)
}
func (lfu *lfuCache) remove(key interface{}) error {
	value, err := lfu.storage.Get(key)
	if err != nil {
		return err
	}

	err = lfu.storage.Remove(key)
	if err != nil {
		return err
	}

	// TODO: find a better way to remove the item from the heap (if there is one).
	lfuItem := value.(lfuItem)
	for i, heapItem := range lfu.heap {
		if heapItem == lfuItem.heapItem {
			heap.Remove(&lfu.heap, i)
		}
	}

	return nil
}
func (lfu *lfuCache) IsEmpty() bool {
	lfu.mutex.Lock()
	defer lfu.mutex.Unlock()

	return lfu.isEmpty()
}

func (lfu *lfuCache) isEmpty() bool {
	return lfu.heap.Len() < 1
}
