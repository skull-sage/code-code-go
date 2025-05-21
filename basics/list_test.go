package basics

import (
	"container/list"
	"fmt"
	"testing"
)

func TestList(t *testing.T) {
	stack := list.New()
	stack.PushFront(1)
	stack.PushFront(2)
	stack.PushFront(3)

	frontVal := stack.Front().Value
	backVal := stack.Back().Value
	fmt.Println("front is: ", frontVal, "back is: ", backVal)

}
