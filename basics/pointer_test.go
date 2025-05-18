package basics

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mutatellA(llA *ListNode) {
	llA.Val = 2
	llA.Next = &ListNode{Val: 3, Next: nil}
}

func TestStructPtr(t *testing.T) {
	llA := &ListNode{Val: 1, Next: nil}
	mutatellA(llA)
	fmt.Println(llA.Val, "next->", llA.Next.Val)

	llB := llA
	llB.Val = 4

	fmt.Println(llA.Val, "next->", llB.Next.Val)

	newLL := &ListNode{Val: 10, Next: nil}
	llB.Next = newLL
	//llB.Next = llA

	fmt.Println(llA.Val, "llB:", llB.Val, "next->", llA.Next.Val)
}
