package common_ds

import (
	"testing"
)

func reverseLL(start *ListNode) (*ListNode, *ListNode) {
	if start == nil || start.Next == nil {
		return start, start
	}

	// 2->3->4
	prev := start     // prev~> 2
	cur := start.Next // curr ~> 3
	for cur != nil {

		next := cur.Next // next ~> 4
		cur.Next = prev  // cur ~-> 3, 3.Next -> 2
		prev = cur       // prev ~> 3
		cur = next       // cur ~> 4
	}
	start.Next = nil
	return prev, start

}

func reverseBetween(head *ListNode, left int, right int) *ListNode {

	if head == nil {
		return nil
	}

	var startPrev *ListNode
	start := head

	for idx := 1; idx < left && start != nil; idx++ {
		startPrev = start
		start = start.Next
	}

	end := start
	for idx := left; idx < right; idx++ {
		end = end.Next
	}

	//fmt.Println("#", start, end)
	endNext := end.Next
	end.Next = nil

	revStart, revEnd := reverseLL(start)

	if startPrev == nil {
		head = revStart
	} else {
		startPrev.Next = revStart
	}

	if endNext != nil {
		revEnd.Next = endNext
	}

	return head

}

func TestReverseLL(t *testing.T) {
	head := CreateLL([]int{1, 2, 3})
	head = reverseBetween(head, 2, 3)
	logLL(head)
}
