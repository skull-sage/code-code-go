package basics

import (
	"fmt"
	"testing"
)

/*
1. new(T) creates a ptr *T with zeroed value
2. new(T); T can be any type: int, string, struct
3. For map[k]val, channel & slice, make(T) is more suitable
4. &[]int{} vs make([]int)
*/
func TestMakeSlicePtr(t *testing.T) {
	arr := make([]int, 0, 4)
	fmt.Printf("ptr %p\n", arr)
	arr = append(arr, 1)
	arr = append(arr, 2)
	arr = append(arr, 3)
	arr = append(arr, 4)
	fmt.Printf("ptr %p\n", arr)
	arr = append(arr, 5)
	fmt.Printf("ptr %p\n", arr)
}

func TestMakeWithDefinedType(t *testing.T) {
	type IntArr []int
	arr := make(IntArr, 0, 4)
	arr = append(arr, 1)
}

func TestMapMutation(t *testing.T) {
	type Name struct {
		first  string
		second string
	}

	nameMap := make(map[int]Name)
	nameMap[1] = Name{
		first:  "first",
		second: "second",
	}

	aname := nameMap[1]
	aname.first = "first changed?"
	fmt.Println(nameMap[1]) // unchanged, still first

	nameMapPtr := make(map[int]*Name)
	nameMapPtr[1] = &Name{
		first:  "first",
		second: "second",
	}

	aNamePtr := nameMapPtr[1]
	aNamePtr.first = "first changed"
	fmt.Println(nameMapPtr[1]) // first changed

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
