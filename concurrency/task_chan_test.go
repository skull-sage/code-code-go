package concurrency

import (
	"fmt"
	"testing"
	"time"
)

func triggerSqr(x int, sqrChan chan int) {
	time.Sleep(2 * time.Second)
	sqrChan <- x * x
}

func TestSqrChan(t *testing.T) {
	sqrChan := make(chan int, 5)

	result := make([]int, 0)
	for idx := range cap(sqrChan) {
		go triggerSqr(idx, sqrChan)
	}

	select {
	case sq := <-sqrChan:
		result = append(result, sq)
		fmt.Println(result)
		// default:
		// 	fmt.Println("# To Sleep on result: ", result)
		// 	time.Sleep(1 * time.Second)
	}

	fmt.Println("#finalally:", result)
}
