package concurrency

import (
	"fmt"
	"testing"
)

func BenchmarkContextSwitch(t *testing.B) {
	var data struct{}
	fmt.Println(data)
}
