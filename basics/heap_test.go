package basics

import (
	"container/heap"
	"fmt"
	"testing"
)

type IntHeap []int

func (h *IntHeap) Len() int { return len(*h) }

func (h *IntHeap) Less(i, j int) bool {

	arr := *h
	return arr[i] < arr[j]
}

func (h *IntHeap) Swap(i, j int) {
	arr := *h
	arr[i], arr[j] = arr[j], arr[i]
}

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int)) // mutate the array
}

func (h *IntHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[0 : len(old)-1]
	return x
}

func TestHeap(t *testing.T) {
	h := &IntHeap{3, 4, 2, 6}

	heap.Init(h)
	fmt.Println(heap.Pop(h), "#", h)
	h.Push(3)
	fmt.Println(h.Pop(), "#", h)
	heap.Push(h, 1)
	fmt.Println(heap.Pop(h), "#", h)

}
