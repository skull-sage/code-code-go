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
	h := &IntHeap{3, 4, 6}

	heap.Init(h)
	//fmt.Println("popped-heap", heap.Pop(h), "#", h)
	heap.Push(h, 4)
	heap.Push(h, 2)
	heap.Push(h, 7)
	heap.Push(h, 5)

	fmt.Println("popped-h", heap.Pop(h), "#", h)
	heap.Push(h, 1)
	fmt.Println("popped-heap", heap.Pop(h), "#", h)

}
