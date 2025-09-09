package concurrency

import (
	"fmt"
	"testing"
	"time"
)

func TestSemaphor(t *testing.T) {
	// buffered channel as a semaphore blocking mechanism

	signalBuff := make(chan bool, 3)

	// worker will recieve it and send it
	// a recive must follow  a send:
	var worker func(wid, counter int) = func(wid, counter int) {
		// sending to buffer gets blocked only after buff is full
		signalBuff <- true
		// critical section start
		time.Sleep(3 * time.Second)
		fmt.Println("Routine id:", wid, "Printing Counter: ", counter)

		//critical section ends

		<-signalBuff // receiving

	}

	fmt.Println("Semaphore with capacity 3")
	counter := 0
	for id := range 10 {
		counter++
		go worker(id, counter)
		counter = counter % cap(signalBuff)
	}

	// give other routine sometime to work before exiting
	time.Sleep(10 * time.Second)

}
