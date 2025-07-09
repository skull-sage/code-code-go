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
type IdxKeyExtractor[K Idxkey] func(v any) K
type WeightKeyExtractor[W WeightKey] func(v any) W

type PQueueMap[K Idxkey, W WeightKey, V any] struct {
	dataHeap  *DataHeap[K]
	dataMap   map[K]V
	keyExt    IdxKeyExtractor[K]
	weightExt WeightKeyExtractor[W]
}

func NewMinPQ[K Idxkey, W WeightKey, V *any](keyExt IdxKeyExtractor[K], weightExt WeightKeyExtractor[W]) *PQueueMap[K, V] {

	aMap := make(map[K]V, 4)
	fmt.Printf("created map address: %p\n", aMap)
	return &PQueueMap[K, W, V]{
		dataHeap: &DataHeap[V]{
			arr:      []V{},
			lessComp: func(a, b V) bool { return a < b },
		},
		dataMap: aMap,
		keyExt:  keyExt,
	}
}

func TestPQMap(t *testing.T) {
	type vertx struct {
		key  int
		rank int
	}
	keyExt := func(v vertx) int { return v.key }
	pq := NewMinPQ(keyExt, lessComp)

}
