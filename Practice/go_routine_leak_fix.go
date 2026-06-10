package main

import (
	"fmt"
	"time"
)

func doWork(done <-chan bool) {
	for {
		select {
		case <-done:
			return
		default:
			fmt.Println("DOING WORK")
		}
	}
}

func main() {
	done := make(chan bool)

	go doWork(done)

	time.Sleep(time.Second * 3)

	// channel closed by user
	close(done)

	fmt.Println("Main program is still running...")
	time.Sleep(time.Second * 5)
	fmt.Println("bye")
}
