package concurrency

import (
	"fmt"
	"testing"
	"time"
)

func TestSelectClaus(t *testing.T) {

	chanA := make(chan string)
	chanB := make(chan string)

	go func() {
		defer close(chanA)
		time.Sleep(2 * time.Second)
		chanA <- "Channel A Bullet"
	}()

	go func() {
		defer close(chanB)
		time.Sleep(1 * time.Second)
		chanB <- "Channel B Bullet"
	}()

	select {

	case msgA := <-chanA:
		fmt.Println(msgA)

	case msgB := <-chanB:
		fmt.Println(msgB)

		//default:
		//	fmt.Println("No channel is ready")

	}

}
