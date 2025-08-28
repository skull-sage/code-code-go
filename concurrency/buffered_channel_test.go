package concurrency

import (
	"fmt"
	"testing"
	"time"
)

type CountSum struct {
	counter int
	sum     int
}

func fillBuffer(chanBuff chan int) {
	num := 0

	for {
		num++
		chanBuff <- num
		fmt.Println("- filling: ", num)
		time.Sleep(1 * time.Second)
	}
}

func TestBufferChannel(t *testing.T) {
	chanBuff := make(chan int, 4)
	count := &CountSum{}

	go fillBuffer(chanBuff)

	count.sum = 0
	for {

		if count.counter%4 == 0 {
			time.Sleep(1 * time.Second)
		}

		num := <-chanBuff
		count.counter++
		count.sum += num
		fmt.Printf("# Computing %d => num: %d, sum:%d\n", count.counter, num, count.sum)

	}

	//fmt.Println("Hola Buffer")
}
