package bitman

import (
	"fmt"
	"testing"
)

func TestFindLowestSetBit(t *testing.T) {
	x := 4

	fmt.Printf("x:%b, x ^ (x-1):%b", x, x^(x-1))

	num := x & (x ^ (x - 1))
	fmt.Printf("lowest-sb of (x)=%d\n", num)

	fmt.Printf("least-significant-bit of (x)=%d\n", (x & -x))
}

func TestCountSetBit(t *testing.T) {
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

func TestBitAnd(t *testing.T) {

	fmt.Printf("%b %b x&(-x)=%b\n", 5, -5, 5&(-5))
}
