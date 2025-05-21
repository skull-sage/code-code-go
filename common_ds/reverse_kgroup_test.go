package common_ds

import (
	"fmt"
	"testing"
)

func reverse_group(curr *ListNode, prev *ListNode, count int, k int) *ListNode {

	var tail *ListNode

	if count == k {
		tail = curr.Next
		curr.Next = prev

	} else if count < k {
		tail = reverse_group(curr.Next, curr, count+1, k)
		curr.Next = prev
	}

	fmt.Println(curr.Val, ".Next->", prev.Val, "tail: ", tail.Val)

	return tail

}

func reverseKGroup(head *ListNode, k int) *ListNode {

	if k == 1 {
		return head
	}
	size := 0
	for start := head; start != nil; {
		size++
		start = start.Next
	}

	segment := size / k

	start := head
	for idx := 0; idx < segment; idx++ {
		start = reverse_group(start, nil, 1, k)
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

	reverse_group(ll.Next, ll, 2, 3)

	//_logLink(next)
}
