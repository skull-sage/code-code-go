package ordered_srch

import (
	"fmt"
	"testing"
)

func merge(llA *ListNode, llB *ListNode) *ListNode {
	headA := &ListNode{Val: -1, Next: nil}
	headA.Next = llA
	prevA := headA

	for llA != nil && llB != nil {

		//fmt.Println("prevA->", prevA.Val, "LLA->", llA.Val, "LLB->", llB.Val)

		if llA.Val <= llB.Val { // move llA pointers: llA & prevA
			prevA = llA    // llA is prevA.Next
			llA = llA.Next // move llA to its next
		} else {
			copyNextB := llB.Next
			prevA.Next = llB
			llB.Next = llA
			prevA = prevA.Next
			llB = copyNextB
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
		itemSlice[0] = listLL[0]
		start = 1
	}

	jdx := start
	for idx := start; idx < l; idx += 2 {
		item := merge(listLL[idx], listLL[idx+1])
		itemSlice[jdx] = item // merge listLL[idx] and listLL[idx+1]
		jdx++
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

func TestMergeKLists(t *testing.T) {

	arr1 := [][]int{
		{1, 4, 5},
		{1, 3, 4},
		{2, 6},
		//	{2, 3, 7, 8}, // the last comma is needed
	}

	llArr := _createLLArr(arr1)

	// for idx := 0; idx < len(llArr); idx++ {
	// 	_logLink(llArr[idx])
	// }

	result := mergeKLists(llArr)
	fmt.Println("\nFinal Result:")
	_logLink(result)
}
