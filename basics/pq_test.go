package basics

import (
	"container/heap"
	"fmt"
	"testing"
)

type DataItem[K comparable, V any] struct {
	ValueArr []V
	LessFunc func(aVal, bVal V) bool
}

func (d *DataItem[K, V]) Len() int     { return len(d.ValueArr) }
func (d *DataItem[K, V]) Push(val any) { d.ValueArr = append(d.ValueArr, val.(V)) }
func (d *DataItem[K, V]) Pop() any {
	l := len(d.ValueArr)
	val := d.ValueArr[l-1]
	d.ValueArr = d.ValueArr[:l-1]
	return val
}

func (d *DataItem[K, V]) Less(aIdx, bIdx int) bool {
	aVal := d.ValueArr[aIdx]
	bVal := d.ValueArr[bIdx]
	return d.LessFunc(aVal, bVal)
}

func (d *DataItem[K, V]) Swap(i, j int) {
	d.ValueArr[i], d.ValueArr[j] = d.ValueArr[j], d.ValueArr[i]
}

func TestDataHeap(t *testing.T) {
	type vertx struct {
		rank int
		val  string
	}
	h := &DataItem[int, vertx]{
		ValueArr: []vertx{},
		LessFunc: func(aVal, bVal vertx) bool {
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
	fmt.Println(h.ValueArr)
	fmt.Println()
}
