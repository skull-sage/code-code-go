package bitman

import (
	"fmt"
	"testing"
)

func TestBitOps(t *testing.T) {
	// to know k-th bit use AND
	x := 36
	fmt.Printf("x=%d in binary => %b\n", x, x)
	k := 3 // if bit index are counting from 1
	resultWithAnd := x & (1 << (k - 1))
	fmt.Printf("%d-th bit is %b\n", k, resultWithAnd)

	// To Set k-th bit use OR
	resultWithOR := x | (1 << (k - 1))
	fmt.Printf("After bit-set with OR, is %b\n", resultWithOR)

	// To toggle k-th bit use XOR
	resultWithXOR := x ^ (1 << (k - 1))
	fmt.Printf("After toggling: %b\n", resultWithXOR)

	// to toggle all bit use uniary XOR (^)
	resultWithUnaryXOR := ^x
	fmt.Printf("After unary XOR: %b\n", resultWithUnaryXOR)

}

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

func TestBitMan(t *testing.T) {
	//findingLowestSetBit() // or right most set bit
	//countSetBitExample()
	//ClearLowestSetBitExample()
	//findUniqueElmInArr() // where all other element appear twice

}
