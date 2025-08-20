package adhoc_critical

// we will need a data structure
// manage priority: heap doesn't answer existance unless implemented with java:TreeMap
//                  ordering by freq and acccessTime 
// manage existance: *ptr reference with key map

import "container/heap"

type Entry struct{
    key int
    val int
    freq int
    accessTime int
    pdx int // priority que index
}

type HStore []*Entry

func (store HStore) Len() int {
    return len(store)
} 
func (store HStore) Less(i, j int) bool {
    
     if store[i].freq == store[j].freq {
        return store[i].accessTime < store[j].accessTime
    } else {
        return store[i].freq < store[j].freq
    }
}
func (h HStore) Swap(i, j int){
    h[i], h[j] = h[j], h[i]
    h[i].pdx = i // to (re)set new index
    h[j].pdx = j // to (re)set new index
}

func (store *HStore) Push(entry any) { 
    *store = append(*store, entry.(*Entry))
}

func (store *HStore) Pop() any {
    l := len(*store)
    h := *store

    if l <= 0 {
        return nil
    }

    e := h[l-1]
    *store = h[:l-1]
    return e
}
 

 
type LFUCache struct {
    store *HStore 
    elMap map[int]*Entry
    capacity int
    clock int
}

func (this LFUCache) hasCapacity() bool {
    return this.store.Len() < this.capacity
}


func Constructor(capacity int) LFUCache {
   // store := new(HStore)
    return LFUCache {
        store : new(HStore), 
        elMap : make(map[int]*Entry, 0),
        capacity: capacity,
        clock: 0 }
}
 

func (this *LFUCache) Get(key int) int {
    
    this.clock++
    entry, ok := this.elMap[key]
    if !ok {
        return -1
    }

    entry.accessTime = this.clock
    entry.freq++
    heap.Fix(this.store, entry.pdx)

    return entry.val

}


func (this *LFUCache) Put(key int, value int)  {
    
    // check exist
    // check capacity for a new key insertion & evict
    // insert in pq, map & list
 
    this.clock++
    entry, ok := this.elMap[key]
    if ok { 
        
        entry.val = value // replace value

        entry.accessTime = this.clock
        entry.freq++
        heap.Fix(this.store, entry.pdx)
        return
    }  
    
    if this.hasCapacity() == false { // make room for new entry
        entry := heap.Pop(this.store).(*Entry)
        delete(this.elMap, entry.key)
    }

    newEntry := &Entry{key:key, val:value, freq:1, accessTime:this.clock, pdx: len(*this.store)}
    heap.Push(this.store, newEntry)
    this.elMap[key] = newEntry

}