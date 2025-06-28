package basics

type Heap[T any] struct {
	arr       []T
	extractor func(T) int
}

func NewHeap[T any](extractor func(T) int) *Heap[T] {
	return &Heap[T]{
		arr:       make([]T, 0),
		extractor: extractor,
	}
}
