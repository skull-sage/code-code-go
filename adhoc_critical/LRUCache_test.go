package adhoc_critical

import "container/list"
 
type Entry struct {
   key int
   val int
}


type LRUCache struct {
   ll *list.List
   elmMap map[int]*list.Element
   capacity int
}

func (this LRUCache) hasCapacity() bool {
   return this.ll.Len() < this.capacity
}

func Constructor(capacity int) LRUCache {
   return LRUCache{list.New(), make(map[int]*list.Element), capacity}
}


func (this *LRUCache) Get(key int) int {

   el, ok:= this.elmMap[key]
   if !ok {
	   return -1
   } 

   //this.ll.Remove(el)
   this.ll.MoveToFront(el)
   entry := el.Value.(Entry)

   return entry.val
}


func (this *LRUCache) Put(key int, value int)  {

  // fmt.Println("# putting: ", key, value, "capacity", this.hasCapacity())
  
   elm, ok := this.elmMap[key]
   if ok { // first remove if exist, we are gonna re insert anyway 
	  this.ll.MoveToFront(elm)
	  entry := elm.Value.(Entry)
	  entry.val = value;
	  return
   } 

   if this.hasCapacity() == false {
	//   fmt.Println("capacity drop: ", this.ll.Back().Value.(Entry))
	   back := this.ll.Back() 
	   entry := back.Value.(Entry)
	   delete(this.elmMap, entry.key)
	   this.ll.Remove(back)
	  
   }
	
   this.ll.PushFront(Entry{key, value}) 
   this.elmMap[key] = this.ll.Front()
   
}

func logList(ll *list.List) {

   for el := ll.Front(); el != nil; el = el.Next() {
	   fmt.Printf("(%d %d) ", el.Value.(Entry))
   }

   fmt.Println()

}
/**
* Your LRUCache object will be instantiated and called as such:
* obj := Constructor(capacity);
* param_1 := obj.Get(key);
* obj.Put(key,value);
*/