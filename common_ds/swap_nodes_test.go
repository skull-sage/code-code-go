package common_ds

import "testing"

func swapNodes(head *ListNode, k int) *ListNode {

	fst := head
	snd := head
	fast := head
	slow := head

	for idx := 0; idx < k-1; idx++ {
		fast = fast.Next
	}

	fst = fast

	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}

	val := fst.Val
	fst.Val = snd.Val
	snd.Val = val

	return head

}

func TestSwapNodes(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	ll := CreateLL(arr)
	swapNodes(ll, 4)
	_logLink(ll)
}
