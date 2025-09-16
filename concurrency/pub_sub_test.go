package concurrency

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type SyncCounter struct {
	rwMutex sync.RWMutex
	val     int
}

func (sc *SyncCounter) get(sleepTime int) int {
	sc.rwMutex.RLock()
	defer sc.rwMutex.RUnlock()
	time.Sleep(time.Duration(sleepTime) * time.Second)
	return sc.val

}

func readMultiply(goId int, sc SyncCounter, x int) {

	sleepTime := 3 - goId
	result := sc.get(sleepTime) * x
	fmt.Println("# Read Routine", goId, "=>", result)

}

func increament(counter SyncCounter) {
	counter.rwMutex.Lock()
	defer counter.rwMutex.Unlock()

	time.Sleep(1 * time.Second)
	counter.val++
}

func TestRWMutex(t *testing.T) {

	counter := SyncCounter{val: 2}

	var wg sync.WaitGroup

	for idx := range 3 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// recommended way to recive a result from routine is using channel-select
			result := readMultiply(counter, idx, 3-idx)
			time.Sleep(100 * time.Millisecond)
			fmt.Println("routine", idx, result)
		}()
	}

	wg.Wait()

}
