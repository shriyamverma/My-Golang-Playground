package main

import "fmt"

func main() {
	chan1 := make(chan string)
	chan2 := make(chan string)

	go func() {
		fmt.Println("Data inserted in channel1")
		chan1 <- "dog"
	}()

	go func() {
		fmt.Println("Data inserted in channel2")
		chan2 <- "cat"
	}()

	fmt.Println("1")

	// select statement will block until one of the statement runs
	select {
	case msgFromChan1 := <-chan1:
		fmt.Println(msgFromChan1)
	case msgFromChan2 := <-chan2:
		fmt.Println(msgFromChan2)
	}

	fmt.Println("2")
}
