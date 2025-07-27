package common_ds

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func _logLink(ll *ListNode) {
	for ptr := ll; ptr != nil; ptr = ptr.Next {
		fmt.Print(ptr.Val, "->")
	}
	fmt.Print(nil)
	fmt.Println()
}

func CreateLL(arr []int) *ListNode {
	head := &ListNode{Val: arr[0], Next: nil}
	curr := head
	for idx := 1; idx < len(arr); idx++ {
		curr.Next = &ListNode{Val: arr[idx], Next: nil}
		curr = curr.Next
	}
	return head
}

func _createLLArr(arr [][]int) []*ListNode {

	var llArr []*ListNode = make([]*ListNode, len(arr), len(arr))

	for idx := 0; idx < len(arr); idx++ {
		elmArr := arr[idx]
		llArr[idx] = &ListNode{Val: elmArr[0], Next: nil}
		head := llArr[idx] // head of the linked list for this arr dat

		for jdx := 1; jdx < len(elmArr); jdx++ {
			head.Next = &ListNode{Val: elmArr[jdx], Next: nil}
			head = head.Next
		}
	}

	return llArr

}
