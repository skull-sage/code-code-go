package concurrency

import (
	"fmt"
	"strconv"
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

func CreateSemaphore(capacity int) Semaphore {
	return Semaphore{signalBuff: make(chan struct{}, capacity)}
}

func TestSemaphor(t *testing.T) {
	// buffered channel as a semaphore blocking mechanism

	type task struct {
		id       int
		workTime int
	}
	// we will use waitgroup to make calling goroutine wait (in this context main to wait)
	var wg sync.WaitGroup
	semaphore := CreateSemaphore(3)
	result := make([]string, 10)

	work := func(t task) {
		defer wg.Done()
		defer semaphore.release()

		fmt.Println("# Processing task", t.id)
		time.Sleep(time.Duration(t.workTime) * time.Second)
		result = append(result, strconv.Itoa(t.id)+" Completed")
	}

	for id := range len(result) {
		task := task{id, 3}
		wg.Add(1)
		// process task
		semaphore.acquire()
		go work(task)
		fmt.Println(id, "triggered")
	}

	// make calling wg to wait
	wg.Wait()

}
