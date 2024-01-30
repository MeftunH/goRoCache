package goRoCache

import (
	"testing"
)

func Test_lfuHeap_Len(t *testing.T) {
	tests := []struct {
		name string
		h    lfuHeap
		want int
	}{
		{
			name: "Empty heap",
			h:    lfuHeap{},
			want: 0,
		},
		{
			name: "Heap with 3 items",
			h: lfuHeap{
				&lfuHeapItem{value: "item1", frequency: 2},
				&lfuHeapItem{value: "item2", frequency: 3},
				&lfuHeapItem{value: "item3", frequency: 1},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.Len(); got != tt.want {
				t.Errorf("Len() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_lfuHeap_Less(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		h    lfuHeap
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("Less() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestLessFrequencyOfItemAtIndexILessThanFrequencyOfItemAtIndexJ(t *testing.T) {
	heap := lfuHeap{
		&lfuHeapItem{value: "item1", frequency: 2},
		&lfuHeapItem{value: "item2", frequency: 3},
	}
	result := heap.Less(0, 1)
	if !result {
		t.Errorf("Less() = %v, want %v", result, true)
	}
}
func Test_lfuHeap_Pop_EmptyHeap(t *testing.T) {
	h := lfuHeap{}
	initialLen := h.Len()
	got := h.Pop()
	newLen := h.Len()
	if newLen != initialLen {
		t.Errorf("Pop() should not change length of empty heap")
	}
	if got != nil {
		t.Errorf("Pop() should return nil for empty heap")
	}
}
func Test_lfuHeap_Push_NonNilItem_IncreasesLength(t *testing.T) {
	h := lfuHeap{}
	initialLen := h.Len()
	item := &lfuHeapItem{value: "item1", frequency: 2}
	h.Push(item)
	newLen := h.Len()
	if newLen != initialLen+1 {
		t.Errorf("Push() did not increase length by 1")
	}
}
func Test_lfuHeap_Pop_DecreasesLength(t *testing.T) {
	h := lfuHeap{
		&lfuHeapItem{value: "item1", frequency: 2},
		&lfuHeapItem{value: "item2", frequency: 3},
	}
	initialLen := h.Len()
	h.Pop()
	newLen := h.Len()
	if newLen != initialLen-1 {
		t.Errorf("Pop() did not decrease length by 1")
	}
}
func Test_lfuHeap_Push_HigherFrequencyItem(t *testing.T) {
	h := lfuHeap{
		&lfuHeapItem{value: "item1", frequency: 2},
		&lfuHeapItem{value: "item2", frequency: 3},
	}
	initialLen := h.Len()
	item := &lfuHeapItem{value: "item3", frequency: 4}
	h.Push(item)
	newLen := h.Len()
	if newLen != initialLen+1 {
		t.Errorf("Push() did not increase length by 1")
	}
}
func Test_lfuHeap_Push_LowerFrequencyItem(t *testing.T) {
	h := lfuHeap{
		&lfuHeapItem{value: "item1", frequency: 2},
		&lfuHeapItem{value: "item2", frequency: 3},
	}
	initialLen := h.Len()
	item := &lfuHeapItem{value: "item3", frequency: 1}
	h.Push(item)
	newLen := h.Len()
	if newLen != initialLen+1 {
		t.Errorf("Push() did not increase length by 1")
	}
}
func Test_lfuCache_get_ExistingItem(t *testing.T) {
	cache := NewLfu(2)
	cache.Store("key1", "value1")
	value, err := cache.get("key1")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if value != "value1" {
		t.Errorf("get() = %v, want %v", value, "value1")
	}
}

func Test_lfuCache_get_NonExistingItem(t *testing.T) {
	cache := NewLfu(2)
	_, err := cache.get("key1")
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func Test_lfuCache_get_FrequencyIncrement(t *testing.T) {
	cache := NewLfu(2)
	cache.Store("key1", "value1")
	cache.get("key1")
	cache.get("key1")
	if cache.heap[0].frequency != 2 {
		t.Errorf("Frequency = %v, want %v", cache.heap[0].frequency, 2)
	}
}

func Test_lfuCache_get_FullCache(t *testing.T) {
	cache := NewLfu(2)
	cache.Store("key1", "value1")
	cache.Store("key2", "value2")
	value, err := cache.get("key1")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if value != "value1" {
		t.Errorf("get() = %v, want %v", value, "value1")
	}
}

func Test_lfuCache_get_EmptyCache(t *testing.T) {
	cache := NewLfu(2)
	_, err := cache.get("key1")
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}
func Test_lfuCache_store_EmptyCache(t *testing.T) {
	cache := NewLfu(2)
	err := cache.store("key1", "value1")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func Test_lfuCache_store_FullCache(t *testing.T) {
	cache := NewLfu(2)
	cache.store("key1", "value1")
	cache.store("key2", "value2")
	err := cache.store("key3", "value3")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func Test_lfuCache_store_ExistingItem(t *testing.T) {
	cache := NewLfu(2)
	cache.store("key1", "value1")
	err := cache.store("key1", "value2")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func Test_lfuCache_store_NilKey(t *testing.T) {
	cache := NewLfu(2)
	err := cache.store(nil, "value1")
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func Test_lfuCache_store_NilValue(t *testing.T) {
	cache := NewLfu(2)
	err := cache.store("key1", nil)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func Test_lfuCache_store_CapacityOne(t *testing.T) {
	cache := NewLfu(1)
	cache.store("key1", "value1")
	err := cache.store("key2", "value2")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func Test_lfuCache_store_CapacityZero(t *testing.T) {
	cache := NewLfu(0)
	err := cache.store("key1", "value1")
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func Test_lfuCache_store_LargeValue(t *testing.T) {
	cache := NewLfu(2)
	largeValue := string(make([]byte, 1024*1024))
	err := cache.store("key1", largeValue)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func Test_lfuCache_store_LargeKey(t *testing.T) {
	cache := NewLfu(2)
	largeKey := string(make([]byte, 1024*1024))
	err := cache.store(largeKey, "value1")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func Test_lfuCache_store_MultipleItems(t *testing.T) {
	cache := NewLfu(10)
	for i := 0; i < 10; i++ {
		err := cache.store(i, i)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
	}
}
func Test_lfuCache_Store_EmptyCache(t *testing.T) {
	cache := NewLfu(2)
	err := cache.Store("key1", "value1")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}
func Test_lfuCache_Store_EmptyCache(t *testing.T) {
	cache := NewLfu(2)
	err := cache.Store("key1", "value1")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Retrieve the stored item
	value, err := cache.Get("key1")
	_ = err
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Check if the retrieved value is correct
	if value != "value1" {
		t.Errorf("Get() = %v, want %v", value, "value1")
	}
}
