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

func readMultiply(goId int, workTime int, sc *SyncCounter, x int) {

	fmt.Println("# Read Routine", goId, "is trying to  read-lock")
	sc.rwMutex.RLock()
	defer sc.rwMutex.RUnlock()

	fmt.Println("=> Read Routine", goId, "acquired read-lock")
	time.Sleep(time.Duration(workTime) * time.Second)
	result := sc.val * x
	fmt.Println("=> Read Routine", goId, " computed ", result)

}

func increament(goId int, workTime int, sc *SyncCounter) {

	fmt.Println("# Write Routine", goId, "is trying to write-lock")
	sc.rwMutex.Lock()
	defer sc.rwMutex.Unlock()

	fmt.Println("=> Write Lock acquired")
	time.Sleep(time.Duration(workTime) * time.Second)
	sc.val++

	fmt.Println("=> Write Routine Computed: ", sc.val)

}

func TestRWMutex(t *testing.T) {

	counter := SyncCounter{val: 2}

	go readMultiply(1, 1, &counter, 5)
	time.Sleep(100 * time.Microsecond)

	go increament(2, 3, &counter)
	time.Sleep(100 * time.Microsecond)
	
	go readMultiply(3, 2, &counter, 5)
	time.Sleep(100 * time.Microsecond)
	
	go readMultiply(4, 1, &counter, 5)
	
	
	time.Sleep(15 * time.Second)

}
