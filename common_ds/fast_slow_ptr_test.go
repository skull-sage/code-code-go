package common_ds

import (
	"fmt"
	"testing"
)

func fastSlowPtr(ll *ListNode) {
	fastPtr := ll
	slowPtr := ll

	for fastPtr != nil && fastPtr.Next != nil {
		fastPtr = fastPtr.Next.Next
		// we can move slowPtr without nil check below
		// this
		if fastPtr != nil {
			slowPtr = slowPtr.Next
		}

	}

	fmt.Println("slow is: ", slowPtr.Val)
}

func TestFastSlowPtr(t *testing.T) {
	sampleArr := []int{1, 2}

	for idx := 3; idx < 8; idx++ {
		sampleArr = append(sampleArr, idx)
		ll := CreateLL(sampleArr)
		_logLink(ll)
		fastSlowPtr(ll)
	}

}
