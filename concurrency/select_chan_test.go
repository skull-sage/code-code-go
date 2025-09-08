package concurrency

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestSelectClaus(t *testing.T) {

	chanA := make(chan string)
	chanB := make(chan string)

	go func() {
		defer close(chanA)
		time.Sleep(2 * time.Second)
		chanA <- "Channel A Bullet"
	}()

	go func() {
		defer close(chanB)
		time.Sleep(1 * time.Second)
		chanB <- "Channel B Bullet"
	}()

	select {

	case msgA := <-chanA:
		fmt.Println(msgA)

	case msgB := <-chanB:
		fmt.Println(msgB)

		/*
			default: // default terminates immediately not waiting for any channel
				fmt.Println("No channel is ready")
		*/

	}

}

func TestTimeout(t *testing.T) {
	chanA := make(chan string)
	chanB := make(chan string)

	go func() {
		defer close(chanA)
		time.Sleep(3 * time.Second)
		chanA <- "Channel A Bullet"
	}()

	go func() {
		defer close(chanB)
		time.Sleep(1 * time.Second)
		chanB <- "Channel B Bullet"
	}()

	select {

	case msgA := <-chanA:
		fmt.Println(msgA)

	case msgB := <-chanB:
		fmt.Println(msgB)

	case <-time.After(2 * time.Second):
		fmt.Println("Exit on Timeout(2)")

	}
}

func TestStopMechanics(t *testing.T) {
	chanA := make(chan string)
	chanB := make(chan string)
	chanStop := make(chan bool)

	_, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	go func() {
		count := 0
		for {
			time.Sleep(1 * time.Second)
			count++
			chanA <- ("Channel A Bullet: " + strconv.Itoa(count))
			fmt.Println("channel A triggered: ", count)
		}
	}()

	go func() {
		count := 0
		for {
			time.Sleep(2 * time.Second)
			count++
			chanB <- "Channel B Bullet" + strconv.Itoa(count)
			fmt.Println("channel B triggered: ", count)
		}

	}()

	count := 0

	for {
		count++
		fmt.Println("# Select Case: ", count)

		if count == 10 {
			close(chanStop)
		}

		select {

		case msgA := <-chanA:
			fmt.Println("# Received", msgA)

		case msgB := <-chanB:
			fmt.Println("# Received", msgB)

		case <-chanStop:
			//close(chanA)
			//close(chanB)

			return

			// case <-time.After(10 * time.Second):
			// 	fmt.Println("Timeout(2) - :(")

		default:
			// Do some work or just wait
			fmt.Println("# Goroutine is running...")
			time.Sleep(1 * time.Second)

		}

	}

}
