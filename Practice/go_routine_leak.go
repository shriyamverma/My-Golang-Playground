package main

import (
	"fmt"
	"time"
)

func main() {
	// infinite loop
	go func() {
		for {
			select {
			default:
				fmt.Println("DOING WORK")
			}
		}
	}()

	// this will waste our resources, as the above loop will run for the lifetime of our server
	time.Sleep(time.Hour * 299)
}
