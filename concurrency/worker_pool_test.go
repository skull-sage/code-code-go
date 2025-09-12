package concurrency

import (
	"fmt"
	"testing"
	"time"
)

func worker(id int, jobIn <-chan int, result chan<- string) {
	for job := range jobIn {
		fmt.Printf("Worker %d is processing job: %d\n", id, job)
		time.Sleep(3 * time.Second)
		result <- fmt.Sprintf("job %d is completed!", job)
	}

}

func TestWorkerPool(t *testing.T) {

	jobChan := make(chan int, 8)
	resChan := make(chan string, 0)

	for id := range 4 {
		go worker(id, jobChan, resChan)
	}

	jobEmitter := func() {
		jobCounter := 0
		for {
			jobCounter++
			jobChan <- jobCounter
			time.Sleep(1 * time.Second)
		}
	}

	go jobEmitter()

	for r := range resChan {
		fmt.Println(r)
	}

	//for a bounded job scenerio, you generally shoot all the job
	// close job channel
	// read results from result channel
	// close result channel

}
