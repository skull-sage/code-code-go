package common_ds

import (
	"testing"
)

func reverseList(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	prev := head
	curr := head.Next

	for curr != nil {
		currNext := curr.Next // copy curr next
		// reverse curr ptr
		curr.Next = prev
		// take next step
		prev = curr
		curr = currNext
	}

	head.Next = nil

	return prev
}

func TestReverse(t *testing.T) {
	ll := CreateLL([]int{1, 2, 3, 4, 5})
	rl := reverseList(ll)
	_logLink(rl)
}
