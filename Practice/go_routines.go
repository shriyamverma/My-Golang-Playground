package main

import (
	"fmt"
	"time"
)

func someFunc(str string) {
	fmt.Println(str)
}

func main() {
	// async call
	go someFunc("1")
	go someFunc("2")
	go someFunc("3")

	fmt.Println("Practice")

	// to wait for the go routines to complete otherwise main thread will not wait and will close automatically
	time.Sleep(time.Second * 2)
}
