package basics

import (
	"fmt"
	"testing"
)

func TestStruct(t *testing.T) {

}

func TestSlice(t *testing.T) {
	s1 := new([]int)                              // returns an arr pointer
	fmt.Println("len", len(*s1), "cap", cap(*s1)) // refering array will require *
	*s1 = append(*s1, 1)
	fmt.Println(*s1)

	s2 := make([]int, 5, 5)
	// doesn't require pointer
	// set len to be 5, all elements are zero-ed
	fmt.Println("len", len(s2), cap(s2))
	s2 = append(s2, 6)
	fmt.Println("prints all 6", s2)
}
