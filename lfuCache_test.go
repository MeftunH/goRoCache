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
