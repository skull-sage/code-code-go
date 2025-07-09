package basics

import (
	"fmt"
	"testing"
)

/*
* the primary difference of heap and PriorityQueue use case is the need of decrease_key operation
* and in practice PQ is used for struct type hence the map implementation
 */

type Idxkey interface{ int | int64 | string }
type WeightKey interface{ int | int64 | float64 }

// need to transform into interface
type IdxKeyExtractor[T any, K Idxkey] func(t T) K
type WeightKeyExtractor[T any, W WeightKey] func(t T) W

type PQueueMap[T any, K Idxkey, W WeightKey] struct {
	dataHeap  *DataHeap[T]
	dataMap   map[K]int
	keyExt    IdxKeyExtractor[T, K]
	weightExt WeightKeyExtractor[T, W]
}

func NewMinPQ[T any, K Idxkey, W WeightKey](keyExt IdxKeyExtractor[T, K], weightExt WeightKeyExtractor[T, W]) *PQueueMap[T, K, W] {

	aMap := make(map[K]int, 4)
	fmt.Printf("created map address: %p\n", aMap)
	return &PQueueMap[T, K, W]{
		dataHeap: &DataHeap[T]{
			arr: []T{},
			lessComp: func(a, b T) bool {
				aWeight := weightExt(a)
				bWeight := weightExt(b)
				return aWeight < bWeight
			},
		},
		dataMap: aMap,
		keyExt:  keyExt,
	}
}

func TestPQMap(t *testing.T) {
	type vertx struct {
		id     int
		weight int
	}
	//var keyExt IdxKeyExtractor[vertx, int]
	keyExt := func(v vertx) int { return v.id }
	wExt := func(v vertx) int { return v.weight }
	pq := NewMinPQ(keyExt, wExt)

	fmt.Printf("returned map ref: %p\n", pq.dataMap)

}
