package common_ds

import (
	"fmt"
	"testing"
)

func removeNthFromEnd(head *ListNode, k int) *ListNode {

	fast := head
	slow := head

	for idx := 1; idx <= k-1; idx++ {
		fast = fast.Next // k=4, 2->3->4 for idx <= k -1 = 4 - 1 = 3
	}

	fast = fast.Next
	var prev *ListNode

	for fast != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next
	} 

	if prev != nil {
		prev.Next = slow.Next
	} else { // if prev == nil then k > list length

		head = slow.Next
	}

	return head

}

func TestRemoveNthFromEnd(t *testing.T) {
	arr := []int{1}
	ll := CreateLL(arr)
	ll = removeNthFromEnd(ll, 1)

	_logLink(ll)
}
