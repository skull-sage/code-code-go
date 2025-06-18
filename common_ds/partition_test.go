package common_ds

import (
	"testing"
)

func partition(ll *ListNode, x int) *ListNode {

	largerXHead := &ListNode{Val: 0, Next: nil}
	lteqXHead := &ListNode{Val: 0, Next: ll}
	largerX := largerXHead
	lteqX := lteqXHead

	for ll != nil {
		if ll.Val >= x {
			largerX.Next = ll
			largerX = largerX.Next

		} else {
			lteqX.Next = ll
			lteqX = lteqX.Next
		}

		ll = ll.Next
	}

	largerX.Next = nil
	//_logLink(lteqXHead)
	//_logLink(largerXHead)
	lteqX.Next = largerXHead.Next

	return lteqXHead.Next
}

func TestPartition(t *testing.T) {
	ll := CreateLL([]int{2, 1})
	ll = partition(ll, 2)

	_logLink(ll)
}
