package basics

import (
	"fmt"
	"testing"
)

/*
* the primary difference of heap and PriorityQueue use case is the need of decrease_key operation
* and in practice PQ is used for struct type hence the map implementation
 */

 

 type KeyExtractor[K int | int64 | string, V any] func(v V) K

type PQueueMap[K int | int64 | string, V any] struct {
	dataHeap *DataHeap[K]
	dataMap  map[K]V
}

func NewMinPQ[K int | int64 | string, V any](keyExt KeyExtractor[K, V]) *PQueueMap[K, V] {

	aMap := make(map[K]V, 4)
	fmt.Printf("created map address: %p\n", aMap)
	return &PQueueMap[K, V]{
		dataHeap: &DataHeap[K]{
			arr: []K{},
			lessComp: func(a, b V) bool {
				aKey := keyExt(a)
				bKey := keyExt(b)
				return aKey < bKey
			},
		},
		dataMap: aMap,
	}
}

func TestPQMap(t *testing.T) {
	type vertx struct {
		key  int
		rank int
	}
	lessComp := func(a, b int) bool { return rank < b.rank }
	keyExt := func(v vertx) int { return v.key }
	pq := NewMinPQ(keyExt, lessComp)

}
