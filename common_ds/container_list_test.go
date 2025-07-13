package common_ds

import (
	"container/list"
	"testing"
)

func TestDeque(t *testing.T) {

	// go container list is a double linked list
	// front list.Front() & back list.Back() returns list item
	deque := list.New()
	deque.PushFront(1)
	deque.PushFront(2)
	deque.PushBack(3)

}
