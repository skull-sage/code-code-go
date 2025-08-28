package concurrency

import (
	"fmt"
	"testing"
)

func triggerRange(rangeClose int, numChan chan int) {

	for x := range rangeClose {
		numChan <- x
	}
	close(numChan)
}

func TestRangeClose(t *testing.T) {

	numChan := make(chan int)
	go triggerRange(10, numChan)

	for {
		// we continue to recieve from channel
		num, ok := <-numChan
		if ok { // check if channel is receiving data
			fmt.Println("# Received: ", num)
		} else { // ok is false when not receiving
			fmt.Println("Not Okay!")
			return
		}
	}

}
