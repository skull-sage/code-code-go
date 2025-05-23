package common_ds

import (
	"fmt"
	"testing"
)

func reverse_segment(start *ListNode, end *ListNode) *ListNode {

	if start == end || start.Next == end || end == nil {
		return end
	}

	//var prev *ListNode
	var prev *ListNode = start // 1
	curr := start.Next         // 2
	for curr.Next != end && curr != nil {
		nextCurr := curr.Next
		curr.Next = prev
		prev = curr
		curr = nextCurr
		fmt.Println("curr", curr.Val, "prev", prev.Val)
	}

	start.Next = end
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
		end = end.Next //
		count++
	}

	if end == nil && count < k {
		return head
	} else if end != nil {
		reverseKGroup(end, k)

	} else {

		reverse_segment(start, end)

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

	start := ll
	end := ll.Next.Next.Next
	fmt.Println("start", start.Val, "end", end.Val)
	ll = reverse_segment(start, end)

	fmt.Println(ll.Val, ll.Next.Val, ll.Next.Next.Val)

	//ll = reverseKGroup(ll, 3)

	//_logLink(ll)
}
