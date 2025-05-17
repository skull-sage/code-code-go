package common_ds

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func compareNdCreate(llA **ListNode, llB **ListNode) *ListNode {
	var aNode *ListNode = new(ListNode)

	if (*llA).Val <= (*llB).Val {
		aNode.Val = (*llA).Val
		*llA = (*llA).Next
	} else {
		aNode.Val = (*llB).Val
		*llB = (*llB).Next
	}

	return aNode
}

func merge(llA *ListNode, llB *ListNode) *ListNode {
	headA := &ListNode{Val: -1, Next: nil}
	headA.Next = llA
	prevA := headA

	for llA != nil && llB != nil {

		_logLink(headA)
		fmt.Println("prevA->", prevA.Val, "LLA->", llA.Val, "LLB->", llB.Val)
		fmt.Println()

		if llA.Val <= llB.Val {
			prevA = llA
			llA = llA.Next
		} else {
			copyNextB := llB.Next

			prevA.Next = llB
			prevA = prevA.Next
			llB = copyNextB
		}
		if llB.Val <= llA.Val {
			// establish current llb link in the LLA list
			nextB := llB.Next // first: copy where LLB is pointing t
			prevA.Next = llB  // insert : prevA -> LLB
			llB.Next = llA    // insert: set LLB.Next -> LLA

			// move forward
			prevA = prevA.Next //  move prevA reference to its next
			llB = nextB        // forward LLB to the copied address

		} else {
			prevA = llA    // forward back pointer
			llA = llA.Next // forward LLA

		}

	}

	for llB != nil {
		prevA.Next = llB
		prevA = prevA.Next
		llB = llB.Next
	}

	return headA.Next

}

func mergeItems(listLL []*ListNode) *ListNode {

	l := len(listLL)
	var start int

	var itemSlice []*ListNode

	if l%2 == 0 {
		itemSlice = make([]*ListNode, l/2)
		start = 0
	} else {
		itemSlice = make([]*ListNode, (l/2)+1)
		itemSlice = append(itemSlice, listLL[0])
		start = 1
	}

	for idx := start; idx < l; idx += 2 {
		item := merge(listLL[idx], listLL[idx+1])
		itemSlice = append(itemSlice, item)
	}

	if len(itemSlice) > 1 {
		return mergeItems(itemSlice)
	}

	return itemSlice[0]
}

func mergeKLists(listLL []*ListNode) *ListNode {

	l := len(listLL)
	var resultLL *ListNode

	if l == 0 {
		return nil
	}

	if l == 1 {
		return listLL[0]
	}

	resultLL = mergeItems(listLL)

	return resultLL
}

func _logLink(ll *ListNode) {
	for ptr := ll; ptr != nil; ptr = ptr.Next {
		fmt.Print(ptr.Val, "->")
	}
	fmt.Println()
}

func _createLinkedList(arr []int) *ListNode {
	head := &ListNode{Val: arr[0], Next: nil}
	var prev *ListNode = head
	for idx := 1; idx < len(arr); idx++ {
		prev.Next = &ListNode{Val: arr[idx], Next: nil}
		prev = prev.Next
	}
	return head
}
func TestMergeKLists(t *testing.T) {

	arrA := []int{2, 4, 7, 8}
	arrB := []int{1, 2, 4, 5, 9}
	var llA *ListNode = _createLinkedList(arrA)
	var llB *ListNode = _createLinkedList(arrB)

	//_logLink(llA)
	//_logLink(llB)
	llNew := merge(llA, llB)

	_logLink(llNew)
}
