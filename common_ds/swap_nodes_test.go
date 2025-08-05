package common_ds

import "testing"

func swapNodes(head *ListNode, k int) *ListNode {

	first := head
	snd := head
	fast := head
	slow := head

	for idx := 1; idx <= k-1; idx++ {
		fast = fast.Next // k=4, 2->3->4 for idx <= k -1 = 4 - 1 = 3
	}

	first = fast
	fast = fast.Next
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}

	snd = slow

	val := first.Val
	first.Val = snd.Val
	snd.Val = val

	return head

}

func TestSwapNodes(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	ll := CreateLL(arr)
	swapNodes(ll, 2)
	logLL(ll)
}
