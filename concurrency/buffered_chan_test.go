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

func sqBuffer(chanBuff chan int) {
	num := 0

	for num < cap(chanBuff) {
		num++
		chanBuff <- (num * num)
		fmt.Println("- filling: ", num)
		time.Sleep(1 * time.Second)
	}
}

func TestBufferChannel(t *testing.T) {
	chanBuff := make(chan int, 4)
	count := CountSum{}

	result := make([]int, 0)
	go sqBuffer(chanBuff)

	for len(result) < cap(chanBuff) {

		val := <-chanBuff
		count.counter++
		result = append(result, val)
		fmt.Printf("# Computing %d => num: %d, sum:%d\n", count.counter, val, count.sum)

	}

	fmt.Println("# Hola Buffer", result)
}
