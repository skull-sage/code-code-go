package concurrency

import (
	"fmt"
	"testing"
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

}
