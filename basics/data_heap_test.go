package basics

import (
	"container/heap"
	"fmt"
	"testing"
)

type LessComparator[V any] func(aVal, bVal V) bool
type DataHeap[V any] struct {
	arr      []V
	lessComp LessComparator[V]
}

func (d *DataHeap[V]) Len() int     { return len(d.arr) }
func (d *DataHeap[V]) Push(val any) { d.arr = append(d.arr, val.(V)) }
func (d *DataHeap[V]) Pop() any {
	l := len(d.arr)
	val := d.arr[l-1]
	d.arr = d.arr[:l-1]
	return val
}

func (d *DataHeap[V]) Less(aIdx, bIdx int) bool {
	aVal := d.arr[aIdx]
	bVal := d.arr[bIdx]
	return d.lessComp(aVal, bVal)
}

func (d *DataHeap[V]) Swap(i, j int) {
	d.arr[i], d.arr[j] = d.arr[j], d.arr[i]
}

func (d *DataHeap[V]) Top() V {
	return d.arr[0]
}

func NewMinHeap[V Idxkey]() *DataHeap[V] {
	return &DataHeap[V]{
		arr: []V{},
		lessComp: func(aVal, bVal V) bool {
			return aVal < bVal
		},
	}
}

func TestLinearHeap(t *testing.T) {
	
}
func TestStructHeap(t *testing.T) {
	type vertx struct {
		rank int
		val  string
	}
	h := &DataHeap[vertx]{
		arr: []vertx{},
		lessComp: func(aVal, bVal vertx) bool {
			return aVal.rank < bVal.rank
		},
	}

	heap.Init(h)
	heap.Push(h, vertx{rank: 4, val: "a"})
	heap.Push(h, vertx{rank: 3, val: "b"})
	heap.Push(h, vertx{rank: 1, val: "c"})
	heap.Push(h, vertx{rank: 2, val: "d"})
	heap.Push(h, vertx{rank: 5, val: "e"})

	fmt.Println(heap.Pop(h))
	fmt.Println(h.arr)
	fmt.Println()
}
