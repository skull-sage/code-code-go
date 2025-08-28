package concurrency

import (
	"fmt"
	"sync"
	"testing"
)

func goDuck(idx int, name string, wg *sync.WaitGroup) {

	defer wg.Done()

	//time.Sleep(1000 * time.Millisecond)

	fmt.Printf("%d %s: Duck Duck Go\n", idx, name)

}

func TestHelloRoutine(t *testing.T) {

	duckList := []string{"Puke", "Duke", "Nuke"}

	wg := sync.WaitGroup{}
	for idx, duck := range duckList {
		wg.Add(1)
		go goDuck(idx, duck, &wg)
	}
	wg.Wait()

	fmt.Println("Main: Hello, Duck!")

}
