package concurrency

import (
	"fmt"
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

	semaphore := Semaphore{signalBuff: make(chan struct{}, 3)}

	// worker will recieve it and send it
	// a recive must follow  a send:
	var worker func(counter int) = func(counter int) {
		// sending to buffer gets blocked only after buff is full
		sleepTime := (counter % 3) + 2
		time.Sleep(time.Duration(sleepTime) * time.Second)
		fmt.Println("Printing Counter: ", counter)

	}

	fmt.Println("Semaphore with capacity 3")
	for counter := range 10 {

		semaphore.acquire()
		// critical section start we want 3 worker to work at a time
		go worker(counter)
		//critical section ends
		if counter%3 == 0 {
			fmt.Println()
		}
		semaphore.release() // receiving

		counter = counter % cap(semaphore.signalBuff)
	}

	// give other routine sometime to work before exiting
	time.Sleep(10 * time.Second)

}
