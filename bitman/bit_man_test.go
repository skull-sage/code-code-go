package bitman

import (
	"fmt"
	"testing"
)

func findingLowestSetBit() {
	x := 4

	fmt.Printf("x:%b, x ^ (x-1):%b\n", x, x^(x-1))

	num := x & (x ^ (x - 1))
	fmt.Printf("x=%b, lsb=%d\n", x, num)
}

func countSetBitExample() {
	// hamming weight: counting number of set-bit (1)
	num := 111 - 1
	bitCount := 0

	fmt.Printf("in binary %b\n", num)
	for num != 0 {
		bitCount++
		num = num & (num - 1)
	}

	fmt.Println("bit count\n", bitCount)
}

func TestToggleBit(t *testing.T) {
	findingLowestSetBit() // or right most set bit
	//countSetBitExample()
	//ClearLowestSetBitExample()
	//findUniqueElmInArr() // where all other element appear twice

}
