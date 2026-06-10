package main

import "fmt"

func main() {
	// buffered channel
	charChannel := make(chan string, 3)
	chars := []string{"a", "b", "c"}

	// for select loop pattern
	for _, char := range chars {
		select {
		case charChannel <- char:
		}
	}

	close(charChannel)

	for result := range charChannel {
		fmt.Println(result)
	}
}
