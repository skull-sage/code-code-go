package common_ds

import "testing"

func swapPairs(ll *ListNode) *ListNode {
	head := ll

	for ll != nil && ll.Next != nil {
		first := ll
		scnd := ll.Next

		tmp := first.Val
		first.Val = scnd.Val
		scnd.Val = tmp
		ll = ll.Next.Next
	}

	return head
}

func TestSwapPair(t *testing.T) {
	ll := CreateLL([]int{1, 2, 3, 4})
	ll = swapPairs(ll)
	_logLink(ll)
}
