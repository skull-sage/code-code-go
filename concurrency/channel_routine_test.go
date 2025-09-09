package concurrency

import (
	"fmt"
	"testing"
	"time"
)

func tryChannel(nums []int, sumChan, productChan chan int) {
	sum := 0
	pd := 1
	for _, x := range nums {
		sum += x
		pd *= x
	}

	sumChan <- sum
	productChan <- pd
	fmt.Println("Sending to channel completed")
}

func TestRegularChannel(t *testing.T) {
	sumChan := make(chan int)
	productChan := make(chan int)

	go tryChannel([]int{2, 3, 5}, sumChan, productChan)

	sum := <-sumChan
	product := <-productChan

	fmt.Println("sum", sum, "product", product)

}

func TestChanBlock(t *testing.T) {

	signalChan := make(chan int)

	var worker func() = func() {
		fmt.Println("worker routine: will be blocked until a data arrives")
		time.Sleep(3 * time.Second)
		val := <-signalChan
		fmt.Println("Recieved", val, "let me think for a second")
		time.Sleep(1 * time.Second)
		fmt.Println("Papa here is my ans:", val*2)
	}

	go worker()
	signalChan <- 20
	fmt.Println("main routine:", "is working sleeping ? I am here only after someone recieves it")
	signalChan <- 10
	fmt.Println("I will be never be printed as worker has already returned")

}

func TestBuffBlock(t *testing.T) {
	signalChan := make(chan int, 1)

	var worker func() = func() {
		fmt.Println("worker routine: will be blocked until a data arrives")
		time.Sleep(3 * time.Second)
		val := <-signalChan
		fmt.Println("Recieved", val, "let me think for a second")
		time.Sleep(1 * time.Second)
		fmt.Println("Papa here is my ans:", val*2)
	}

	go worker()
	signalChan <- 20
	fmt.Println("main routine:", "I will be here even if worker is sleeping as signal is now buffered")
	signalChan <- 10
	fmt.Println("I will be here   as worker received first sent item")
}
