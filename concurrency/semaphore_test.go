package concurrency

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type Semaphore struct {
	signalBuff chan struct{}
}

func (sb *Semaphore) acquire() {
	sb.signalBuff <- struct{}{}
}

func (sb *Semaphore) release() {
	<-sb.signalBuff
}

func TestSemaphor(t *testing.T) {
	// buffered channel as a semaphore blocking mechanism

	var wg sync.WaitGroup
	semaphore := Semaphore{signalBuff: make(chan struct{}, 3)}
	cArr := make([]int, 0)
	// worker will recieve it and send it
	// a recive must follow  a send:
	var worker func(counter int) = func(counter int) {
		defer wg.Done()
		// sending to buffer gets blocked only after buff is full
		fmt.Println("# Goroutine entering: ", counter)
		semaphore.acquire()
		time.Sleep(3 * time.Second)
		fmt.Println("# Goroutine processing: ", counter)
		cArr = append(cArr, counter)
		semaphore.release() // receiving
		fmt.Println("# Semaphore released: ", counter)
	}

	fmt.Println("Semaphore with capacity 3")
	wg.Add(10)
	for counter := range 10 {

		// critical section start we want 3 worker to work at a time
		go worker(counter)
		//time.Sleep(5 * time.Millisecond)
		//critical section ends

	}

	// give other routine sometime to work before exiting
	wg.Wait()

	fmt.Println(cArr)

}
