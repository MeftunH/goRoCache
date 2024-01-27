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
	// The maximal amount of cached items.
	capacity int

	// A cache that holds tha data.
	storage Cache

	// A min heap that behaves like a priority queue, where the lowest
	// frequency is the higher priority to remove from the heap.
	heap lfuHeap

	mutex sync.Mutex
}

var _ Cache = (*lfuCache)(nil)

// NewLfu creates a new lfuCache instance using mapCache.
func NewLfu(capacity int) *lfuCache {
	return &lfuCache{
		capacity: capacity,
		storage:  NewMapCache(),
		heap:     lfuHeap{},
	}
}
