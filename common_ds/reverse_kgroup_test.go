package common_ds

import (
	"testing"
)

func reverse_segment(start *ListNode, end *ListNode) *ListNode {

	if start == end || end == nil {
		return end
	}

	//var prev *ListNode
	prev := start      // 1
	curr := start.Next // 2
	tail := end.Next
	for curr != tail && curr != nil {
		nextCurr := curr.Next
		curr.Next = prev
		prev = curr
		curr = nextCurr

	}

	start.Next = nil
	return prev

}

func reverseKGroup(head *ListNode, k int) *ListNode {

	if k == 1 {
		return head
	}

	start := head
	end := head.Next
	count := 1
	for count < k && end != nil {
		end = end.Next //{3, 2}
		count++
	}

	if end == nil && count < k {
		return head
	} else if end.Next == nil {
		head = reverse_segment(start, end)
	} else {
		reverseKGroup(end.Next, k)
	}

	return head

}

func createLL(arr []int) *ListNode {
	head := &ListNode{Val: arr[0], Next: nil}
	curr := head
	for idx := 1; idx < len(arr); idx++ {
		curr.Next = &ListNode{Val: arr[idx], Next: nil}
		curr = curr.Next
	}
	return head
}

func TestLLRevers(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	ll := createLL(arr)

	ll = reverseKGroup(ll, 3)

	//ll = reverseKGroup(ll, 3)

	_logLink(ll)
}
