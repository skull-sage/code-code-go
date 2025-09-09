package basics

import (
	"fmt"
	"sync"
	"testing"
)

func TestZeroValue(t *testing.T) {
	var structVal struct{}
	fmt.Println("struct", structVal)

	var strVal string
	fmt.Println("str", strVal)

	var wg sync.WaitGroup
	fmt.Println("wait-group", wg)
}

func TestNilValue(t *testing.T) {
	// any reference type value that are not initialized
	// will be nil: nil reference

	var ptr *int
	fmt.Println("ptr", ptr)

	var arr []int
	fmt.Println("slice", arr)

	var dataMap map[int]int
	fmt.Println("datamap", dataMap)

	var chanInt chan int
	fmt.Println("chan", chanInt)

}
