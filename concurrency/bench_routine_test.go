package concurrency

import (
	"fmt"
	"testing"
)

func BenchmarkContextSwitch(b *testing.B) {
	var data struct{}
	fmt.Println(data)
}
