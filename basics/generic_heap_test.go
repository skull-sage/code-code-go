package basics

import (
	"container/heap"
	"fmt"
	"testing"
)

type DataHeap[T comparable] struct {
	heapArr  []T
	lessComp func(a T, b T) bool
}

func (h DataHeap[T]) Len() int {
	return len(h.heapArr)
}

func (h DataHeap[T]) Less(i, j int) bool {
	return h.lessComp(h.heapArr[i], h.heapArr[j])
}

func (h *DataHeap[T]) Swap(i, j int) {
	h.heapArr[i], h.heapArr[j] = h.heapArr[j], h.heapArr[i]
}

func (h *DataHeap[T]) Push(val any) {
	h.heapArr = append(h.heapArr, val.(T))
}

func (h *DataHeap[T]) Pop() any {
	l := len(h.heapArr)
	val := h.heapArr[l-1]
	h.heapArr = h.heapArr[0 : l-1]
	return val
}

func NewHeap[T comparable](less func(a, b T) bool, capacity int) *DataHeap[T] {
	dh := &DataHeap[T]{
		heapArr:  make([]T, 0, capacity),
		lessComp: less,
	}
	heap.Init(dh)
	return dh
}

func (dh *DataHeap[T]) String() string {
	return fmt.Sprintf("%v", dh.heapArr)
}

func TestDataHeap(t *testing.T) {
	type Vertex struct {
		phi  *Vertex
		key  int
		name string
	}

	lessComp := func(u, v Vertex) bool {
		return u.key < v.key
	}
	dataHeap := NewHeap(lessComp, 8)

	dataHeap.Push(Vertex{key: 12, name: "Rank 12"})
	dataHeap.Push(Vertex{key: 24, name: "Rank 24"})
	dataHeap.Push(Vertex{key: 5, name: "Rank 5(2)"})
	dataHeap.Push(Vertex{key: 10, name: "Rank 10"})
	dataHeap.Push(Vertex{key: 5, name: "Rank 5"})

	for dataHeap.Len() > 0 {
		v := dataHeap.Pop().(Vertex)
		fmt.Println("#Pop:", v.key, "name: ", v.name)
	}

}
