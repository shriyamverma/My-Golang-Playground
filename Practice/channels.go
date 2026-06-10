package main

import "fmt"

func main() {
	// unbuffered channel has max capacity of 1
	myChannel := make(chan string)

	go func() {
		fmt.Println("Data inserted in channel")
		myChannel <- "data" // blocking, since channel is unbuffered

		// this line will wait until the main code calls for receive data from channel
		fmt.Println("........end of go routine")
	}()

	fmt.Println("1")

	// blocking code (main thread will wait for this receive data or for it to close)
	msg := <-myChannel
	fmt.Println(msg)

	fmt.Println("2")
}
