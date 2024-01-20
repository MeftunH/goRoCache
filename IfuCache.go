package goRoCache

import "container/heap"

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
	//TODO implement me
	panic("implement me")
}

func (h lfuHeap) Push(x any) {
	//TODO implement me
	panic("implement me")
}

func (h lfuHeap) Pop() any {
	//TODO implement me
	panic("implement me")
}

var _ heap.Interface = (*lfuHeap)(nil)

func (h lfuHeap) Len() int {
	return len(h)
}
