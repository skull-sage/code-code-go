package adhoc_critical

import (
	"fmt"
	"sort"
	"testing"
)

// key and value consist of lowercase English letters and digits.
// All the timestamps timestamp of set are strictly increasing.

type StampList []ValStamp

type ValStamp struct {
	val  string
	time int
}

type TimeMap struct {
	valMap map[string]*StampList
}

func constructor() TimeMap {
	return TimeMap{
		valMap: make(map[string]*StampList),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {

	if this.valMap[key] == nil {
		this.valMap[key] = new(StampList)
	}
	list := this.valMap[key]
	*list = append(*list, ValStamp{value, timestamp})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	list, ok := this.valMap[key]

	if !ok {
		return ""
	}

	ptrList := *list

	idx := sort.Search(len(ptrList), func(i int) bool { return ptrList[i].time >= timestamp })

	// we don't have any time <= timestamp for a key mapped by set function
	if idx == 0 && ptrList[idx].time != timestamp {
		return ""
	}

	if idx < len(ptrList) && ptrList[idx].time == timestamp {
		return ptrList[idx].val
	}

	// idx is either len(ptrList) or idx_of_time greater-than timestamp)
	return ptrList[idx-1].val

}

func TestTimedList(t *testing.T) {
	timeMAP := constructor()
	timeMAP.Set("love", "high", 10)
	timeMAP.Set("love", "low", 20)

	fmt.Println(timeMAP.Get("love", 5))
	fmt.Println(timeMAP.Get("love", 10))
	fmt.Println(timeMAP.Get("love", 15))
	fmt.Println(timeMAP.Get("love", 20))
	fmt.Println(timeMAP.Get("love", 25))

}
