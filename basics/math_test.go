package basics

import (
	"fmt"
	"math"
	"testing"
)

func TestFraction(t *testing.T) {
	Val := 1
	A := 0.02
	val := math.Mod(float64(Val)*A, 1.0) // x.y % 1 will result in the fractional part y => x*1 + 0.y = x.y
	fmt.Println("mod 1", val)
}
