package basics

import (
	"container/list"
	"fmt"
	"testing"
)

func mutateCopy(dList list.List) {
	dList.PushBack(100)
}

func mutateRef(dList *list.List) {
	dList.PushBack(200)
}

func TestDList(t *testing.T) {
	// go's list container is essentially a doubly-linked-list
	// can be used as a stack or queue
	dList := list.New() // New() creates a List and return a *list.List
	dList.PushFront(1)
	dList.PushFront(2)
	dList.PushFront(3)

	frontVal := dList.Front().Value
	backVal := dList.Back().Value

	if backVal != 1 {
		t.Errorf("Expected backVal to be 1, but got %v", backVal)
	}

	if frontVal != 3 {
		t.Errorf("Expected frontVal to be 3, but got %v", frontVal)
	}

	len := dList.Len()
	mutateCopy(*dList)
	fmt.Println(dList)
	if len != dList.Len() {
		t.Errorf("Expected length to be unchanged, old: %d, new: %d", len, dList.Len())
	}

	len = dList.Len()
	mutateRef(dList)
	fmt.Println(dList)
	if len == dList.Len() {
		t.Errorf("Expected length to be changed, old: %d, new: %d", len+1, dList.Len())
	}

}
