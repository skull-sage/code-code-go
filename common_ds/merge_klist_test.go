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
	var head *ListNode

	head = compareNdCreate(&llA, &llB)
	prev := head

	for llA != nil && llB != nil {
		prev.Next = compareNdCreate(&llA, &llB)
		prev = prev.Next
		_logLink(head)
	}

	for llA != nil {
		prev.Next = &ListNode{Val: llA.Val, Next: nil}
		prev = prev.Next
		llA = llA.Next
	}

	for llB != nil {
		prev.Next = &ListNode{Val: llB.Val, Next: nil}
		prev = prev.Next
		llB = llB.Next
	}

	return head

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

	arrA := []int{1, 4, 5}
	arrB := []int{1, 3, 4, 6, 7}
	var llA *ListNode = _createLinkedList(arrA)
	var llB *ListNode = _createLinkedList(arrB)

	_logLink(llA)
	_logLink(llB)
	llNew := merge(llA, llB)

	_logLink(llNew)
}
